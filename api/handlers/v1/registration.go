package v1

import (
	"context"
	e "dennic_api_gateway/api/handlers/regtool"
	"dennic_api_gateway/api/models"
	"dennic_api_gateway/api/models/model_user_service"
	ps "dennic_api_gateway/genproto/session_service"
	pb "dennic_api_gateway/genproto/user_service"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/spf13/cast"
	"google.golang.org/protobuf/encoding/protojson"
)

// Register ...
// @Summary Register
// @Description Register - Api for registering users
// @Tags customer
// @Accept json
// @Produce json
// @Param Register body model_user_service.RegisterRequest true "RegisterRequest"
// @Success 200 {object} model_user_service.MessageRes
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/customer/register [post]
func (h *HandlerV1) Register(c *gin.Context) {
	//rdb := redis.NewClient(&redis.Options{
	//	Addr: "redisdb:6379",
	//})

	var (
		body        model_user_service.Redis
		jsonMarshal protojson.MarshalOptions
	)
	jsonMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)

	if e.HandleError(c, err, h.log, http.StatusBadRequest, INVALID_REQUET_BODY) {
		return
	}

	err = body.Validate()

	if e.HandleError(c, err, h.log, http.StatusBadRequest, err.Error()) {
		return
	}

	body.PhoneNumber = strings.TrimSpace(body.PhoneNumber)
	body.LastName = strings.TrimSpace(body.LastName)
	body.LastName = strings.ToLower(body.LastName)
	body.LastName = strings.Title(body.LastName)
	body.FirstName = strings.TrimSpace(body.FirstName)
	body.FirstName = strings.ToLower(body.FirstName)
	body.FirstName = strings.Title(body.FirstName)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	existsPhone, err := h.serviceManager.UserService().UserService().CheckField(ctx, &pb.CheckFieldUserReq{
		Field: "phone_number",
		Value: body.PhoneNumber,
	})

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, NOT_REGISTERED) {
		return
	}
	if existsPhone.Status {
		err = errors.New("you have already registered, try to login")
		_ = e.HandleError(c, err, h.log, http.StatusBadRequest, "Register")
		return
	}

	// TODO A method that sends a code to a number
	body.Code = 7777

	body.Id = uuid.New().String()

	byteDate, err := json.Marshal(&body)
	if e.HandleError(c, err, h.log, http.StatusBadRequest, SERVICE_ERROR) {
		return
	}

	err = h.redis.Client.Set(ctx, body.PhoneNumber, byteDate, h.cfg.Redis.Time).Err()
	if e.HandleError(c, err, h.log, http.StatusInternalServerError, SERVICE_ERROR) {
		return
	}

	c.JSON(http.StatusOK, model_user_service.MessageRes{
		Message: "Code has been sent to you phone number, please check.",
	})
}

// Verify ...
// @Summary Verify
// @Description customer - Api for registering users
// @Tags customer
// @Accept json
// @Produce json
// @Param Verify body model_user_service.Verify true "RegisterModelReq"
// @Failure 200 {object} model_user_service.Response
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/customer/verify [post]
func (h *HandlerV1) Verify(c *gin.Context) {
	//rdb := redis.NewClient(&redis.Options{
	//	Addr: "redisdb:6379",
	//})

	var (
		body        model_user_service.Verify
		user        model_user_service.Redis
		jsonMarshal protojson.MarshalOptions
	)

	jsonMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)

	if e.HandleError(c, err, h.log, http.StatusBadRequest, INVALID_REQUET_BODY) {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	redisRes, err := h.redis.Client.Get(ctx, body.PhoneNumber).Result()

	if e.HandleError(c, err, h.log, http.StatusBadRequest, SERVICE_ERROR) {
		return
	}

	err = json.Unmarshal([]byte(redisRes), &user)

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, SERVICE_ERROR) {
		return
	}

	if body.Code != user.Code {
		err = errors.New(INVALID_CODE)
		_ = e.HandleError(c, err, h.log, http.StatusBadRequest, INVALID_CODE)
		return
	}

	sessions, err := h.serviceManager.SessionService().SessionService().GetUserSessions(ctx, &ps.StrUserReq{
		UserId: user.Id,
	})

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, SERVICE_ERROR) {
		return
	}

	if sessions != nil {
		if sessions.Count >= 3 {
			err = errors.New("the number of devices has exceeded the limit")
			_ = e.HandleError(c, err, h.log, http.StatusBadRequest, "the number of devices has exceeded the limit")
			return
		}
	}

	sessionId := uuid.New().String()

	session, err := h.serviceManager.SessionService().SessionService().CreateSession(ctx, &ps.SessionRequests{
		Id:           sessionId,
		IpAddress:    c.RemoteIP(),
		UserId:       user.Id,
		FcmToken:     body.FcmToken,
		PlatformName: body.PlatformName,
		PlatformType: body.PlatformType,
	})
	if e.HandleError(c, err, h.log, http.StatusInternalServerError, SERVICE_ERROR) {
		return
	}

	user.Password, err = e.HashPassword(user.Password)

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, SERVICE_ERROR) {
		return
	}

	access, refresh, err := h.jwthandler.GenerateAuthJWT(user.PhoneNumber, user.Id, session.Id, "user")

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, SERVICE_ERROR) {
		return
	}

	_, err = h.serviceManager.UserService().UserService().Create(ctx, &pb.User{
		Id:           user.Id,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		BirthDate:    user.BrithDate,
		PhoneNumber:  user.PhoneNumber,
		Password:     user.Password,
		Gender:       user.Gender,
		RefreshToken: refresh,
	})

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, SERVICE_ERROR) {
		return
	}

	err = h.redis.Client.Del(ctx, body.PhoneNumber).Err()

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, SERVICE_ERROR) {
		return
	}

	c.JSON(http.StatusOK, &model_user_service.Response{
		Id:           user.Id,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		BrithDate:    user.BrithDate,
		PhoneNumber:  user.PhoneNumber,
		Gender:       user.Gender,
		AccessToken:  access,
		RefreshToken: refresh,
	})

}

// ForgetPassword ...
// @Summary ForgetPassword
// @Description ForgetPassword - Api for registering users
// @Tags customer
// @Accept json
// @Produce json
// @Param ForgetPassword body model_user_service.PhoneNumberReq true "RegisterModelReq"
// @Success 200 {object} model_user_service.MessageRes
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/customer/forget-password [post]
func (h *HandlerV1) ForgetPassword(c *gin.Context) {
	var (
		body        model_user_service.PhoneNumberReq
		jsonMarshal protojson.MarshalOptions
	)

	jsonMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)

	if e.HandleError(c, err, h.log, http.StatusBadRequest, INVALID_REQUET_BODY) {
		return
	}

	if len(body.PhoneNumber) != 13 && !govalidator.IsNumeric(body.PhoneNumber) {
		err := errors.New(INVALID_PHONE_NUMBER)
		_ = e.HandleError(c, err, h.log, http.StatusBadRequest, INVALID_PHONE_NUMBER)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	existsPhone, err := h.serviceManager.UserService().UserService().CheckField(ctx, &pb.CheckFieldUserReq{
		Field: "phone_number",
		Value: body.PhoneNumber,
	})

	if e.HandleError(c, err, h.log, http.StatusBadRequest, NOT_REGISTERED) {
		return
	}
	if !existsPhone.Status {
		err = errors.New(NOT_REGISTERED)
		_ = e.HandleError(c, err, h.log, http.StatusBadRequest, NOT_REGISTERED)
		return
	}

	codeRed, err := h.redis.Client.Get(ctx, body.PhoneNumber).Result()
	if e.HandleError(c, err, h.log, http.StatusInternalServerError, SERVICE_ERROR) {
		return
	}

	if codeRed != "" {
		err = errors.New(CODE_EXPIRATION_NOT_OVER)
		if e.HandleError(c, err, h.log, http.StatusBadRequest, CODE_EXPIRATION_NOT_OVER) {
			return
		}
	}

	// TODO A method that sends a code to a number
	code := 7777

	err = h.redis.Client.Set(ctx, body.PhoneNumber, code, h.cfg.Redis.Time).Err()
	if e.HandleError(c, err, h.log, http.StatusInternalServerError, SERVICE_ERROR) {
		return
	}

	c.JSON(http.StatusOK, model_user_service.MessageRes{
		Message: "Code has been sent to you phone number, please check.",
	})
}

// UpdatePassword
// @Summary UpdatePassword
// @Description Api for UpdatePassword
// @Tags customer
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param NewPassword  query string true "NewPassword"
// @Success 200 {object} model_user_service.GetUserResp
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/customer/update-password [PUT]
func (h *HandlerV1) UpdatePassword(c *gin.Context) {
	newPassword := c.Query("NewPassword")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	userInfo, err := e.GetUserInfo(c)

	if e.HandleError(c, err, h.log, http.StatusUnauthorized, "missing token in the header") {
		return
	}

	user, err := h.serviceManager.UserService().UserService().Get(ctx, &pb.GetUserReq{
		Field:    "id",
		Value:    userInfo.UserId,
		IsActive: false,
	})

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, SERVICE_ERROR) {
		return
	}

	hashPass, err := e.HashPassword(newPassword)

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, SERVICE_ERROR) {
		return
	}

	response, err := h.serviceManager.UserService().UserService().ChangePassword(ctx, &pb.ChangeUserPasswordReq{
		PhoneNumber: user.PhoneNumber,
		Password:    hashPass,
	})

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, SERVICE_ERROR) {
		return
	}

	c.JSON(http.StatusOK, &models.StatusRes{Status: response.Status})
}

// VerifyOtpCode ...
// @Summary VerifyOtpCode
// @Description VerifyOtpCode - Api for Verify Otp Code users
// @Tags customer
// @Accept json
// @Produce json
// @Param VerifyOtpCode query model_user_service.VerifyOtpCodeReq true "VerifyOtpCode"
// @Failure 200 {object} model_user_service.MessageRes
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/customer/verify-otp-code [post]
func (h *HandlerV1) VerifyOtpCode(c *gin.Context) {
	phoneNumber := c.Query("phone_number")
	code := c.Query("code")

	reqCode := cast.ToInt64(code)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	if len(phoneNumber) != 13 && !govalidator.IsNumeric(phoneNumber) {
		err := errors.New(INVALID_PHONE_NUMBER)
		_ = e.HandleError(c, err, h.log, http.StatusBadRequest, INVALID_PHONE_NUMBER)
		return
	}

	redisRes, err := h.redis.Client.Get(ctx, phoneNumber).Result()

	if e.HandleError(c, err, h.log, http.StatusBadRequest, "code is expired") {
		return
	}

	var redisCode int64

	err = json.Unmarshal([]byte(redisRes), &redisCode)

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, SERVICE_ERROR) {
		return
	}

	if reqCode != redisCode {
		err = errors.New(INVALID_CODE)
		_ = e.HandleError(c, err, h.log, http.StatusBadRequest, INVALID_CODE)
		return
	}

	err = h.redis.Client.Del(ctx, phoneNumber).Err()

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, SERVICE_ERROR) {
		return
	}

	user, err := h.serviceManager.UserService().UserService().Get(ctx, &pb.GetUserReq{
		Field:    "phone_number",
		Value:    phoneNumber,
		IsActive: false,
	})

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, SERVICE_ERROR) {
		return
	}

	access, err := h.jwthandler.GenerateJWT(user.PhoneNumber, user.Id, "user")
	if e.HandleError(c, err, h.log, http.StatusInternalServerError, SERVICE_ERROR) {
		return
	}

	c.JSON(http.StatusOK, &models.AccessToken{Token: access})

}

// Login ...
// @Summary Login
// @Description Login - Api for registering users
// @Tags customer
// @Accept json
// @Produce json
// @Param Login body model_user_service.LoginReq true "Login Req"
// @Success 200 {object} model_user_service.Response
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/customer/login [post]
func (h *HandlerV1) Login(c *gin.Context) {
	var (
		body        model_user_service.LoginReq
		jsonMarshal protojson.MarshalOptions
	)

	jsonMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)

	if e.HandleError(c, err, h.log, http.StatusBadRequest, "Login") {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	if len(body.PhoneNumber) != 13 && !govalidator.IsNumeric(body.PhoneNumber) {
		err := errors.New(INVALID_PHONE_NUMBER)
		_ = e.HandleError(c, err, h.log, http.StatusBadRequest, INVALID_PHONE_NUMBER)
		return
	}

	if !e.ValidatePassword(body.Password) {
		err := errors.New("invalid password")
		_ = e.HandleError(c, err, h.log, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.serviceManager.UserService().UserService().Get(ctx, &pb.GetUserReq{
		Field:    "phone_number",
		Value:    body.PhoneNumber,
		IsActive: false,
	})

	if e.HandleError(c, err, h.log, http.StatusBadRequest, NOT_REGISTERED) {
		return
	}

	if !e.CheckHashPassword(user.Password, body.Password) {
		err = errors.New("incorrect password")
		_ = e.HandleError(c, err, h.log, http.StatusBadRequest, err.Error())
		return
	}

	sessions, err := h.serviceManager.SessionService().SessionService().GetUserSessions(ctx, &ps.StrUserReq{
		UserId:   user.Id,
		IsActive: false,
	})

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, SERVICE_ERROR) {
		return
	}

	if sessions != nil {
		if sessions.Count >= 3 {
			err = errors.New("the number of devices has exceeded the limit")
			_ = e.HandleError(c, err, h.log, http.StatusBadRequest, err.Error())
			return
		}
	}

	sessionId := uuid.New().String()

	_, err = h.serviceManager.SessionService().SessionService().CreateSession(ctx, &ps.SessionRequests{
		Id:           sessionId,
		IpAddress:    c.RemoteIP(),
		UserId:       user.Id,
		FcmToken:     body.FcmToken,
		PlatformName: body.PlatformName,
		PlatformType: body.PlatformType,
	})

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, SERVICE_ERROR) {
		return
	}

	access, refresh, err := h.jwthandler.GenerateAuthJWT(user.PhoneNumber, user.Id, sessionId, "user")
	if e.HandleError(c, err, h.log, http.StatusInternalServerError, SERVICE_ERROR) {
		return
	}

	_, err = h.serviceManager.UserService().UserService().UpdateRefreshToken(ctx, &pb.UpdateRefreshTokenUserReq{
		Id:           user.Id,
		RefreshToken: refresh,
	})

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, SERVICE_ERROR) {
		return
	}

	c.JSON(http.StatusOK, model_user_service.Response{
		Id:           user.Id,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		BrithDate:    user.BirthDate,
		PhoneNumber:  user.PhoneNumber,
		Gender:       user.Gender,
		AccessToken:  access,
		RefreshToken: refresh,
	})
}

// LogOut ...
// @Summary LogOut
// @Description LogOut - Api for registering users
// @Tags customer
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200 {object} model_user_service.MessageRes
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/customer/logout [post]
func (h *HandlerV1) LogOut(c *gin.Context) {
	userInfo, err := e.GetUserInfo(c)

	if e.HandleError(c, err, h.log, http.StatusUnauthorized, "missing token in the header") {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	_, err = h.serviceManager.SessionService().SessionService().DeleteSessionById(ctx, &ps.StrReq{
		Id: userInfo.SessionId,
	})

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, SERVICE_ERROR) {
		return
	}

	c.JSON(http.StatusOK, &model_user_service.MessageRes{Message: "Log out done!"})
}

// SenOtpCode ...
// @Summary SenOtpCode
// @Description SenOtpCode - Api for sen otp code users
// @Tags customer
// @Accept json
// @Produce json
// @Param SenOtpCode body model_user_service.PhoneNumberReq true "RegisterModelReq"
// @Success 200 {object} model_user_service.MessageRes
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/customer/send-otp [post]
func (h *HandlerV1) SenOtpCode(c *gin.Context) {
	var (
		body        model_user_service.PhoneNumberReq
		jsonMarshal protojson.MarshalOptions
	)

	jsonMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)

	if e.HandleError(c, err, h.log, http.StatusBadRequest, INVALID_REQUET_BODY) {
		return
	}

	if len(body.PhoneNumber) != 13 && !govalidator.IsNumeric(body.PhoneNumber) {
		err := errors.New(INVALID_PHONE_NUMBER)
		_ = e.HandleError(c, err, h.log, http.StatusBadRequest, INVALID_PHONE_NUMBER)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	existsPhone, err := h.serviceManager.UserService().UserService().CheckField(ctx, &pb.CheckFieldUserReq{
		Field: "phone_number",
		Value: body.PhoneNumber,
	})

	codeRed, err := h.redis.Client.Get(ctx, body.PhoneNumber).Result()
	if codeRed != "" {
		err = errors.New(CODE_EXPIRATION_NOT_OVER)
		if e.HandleError(c, err, h.log, http.StatusBadRequest, CODE_EXPIRATION_NOT_OVER) {
			return
		}
	}

	if !existsPhone.Status {
		err = errors.New(NOT_REGISTERED)
		_ = e.HandleError(c, err, h.log, http.StatusBadRequest, NOT_REGISTERED)
		return
	}

	// TODO A method that sends a code to a number
	code := 7777

	err = h.redis.Client.Set(ctx, body.PhoneNumber, code, h.cfg.Redis.Time).Err()
	if e.HandleError(c, err, h.log, http.StatusInternalServerError, SERVICE_ERROR) {
		return
	}

	c.JSON(http.StatusOK, model_user_service.MessageRes{
		Message: "Code has been sent to you phone number, please check.",
	})
}
