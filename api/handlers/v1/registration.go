package v1

import (
	"context"
	e "dennic_api_gateway/api/handlers/regtool"
	"dennic_api_gateway/api/models/model_user_service"
	pb "dennic_api_gateway/genproto/user_service"
	"encoding/json"
	"errors"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"google.golang.org/protobuf/encoding/protojson"
	"net/http"
	"time"
)

// Register ...
// @Summary Register
// @Description Register - Api for registering users
// @Tags Register
// @Accept json
// @Produce json
// @Param Register body model_user_service.RegisterRequest true "RegisterRequest"
// @Success 200 {object} model_user_service.MessageRes
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/auth/register/ [post]
func (h *HandlerV1) Register(c *gin.Context) {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	var (
		body        model_user_service.Redis
		jsonMarshal protojson.MarshalOptions
	)
	jsonMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)

	if e.HandleError(c, err, h.log, http.StatusBadRequest, "Register") {
		return
	}

	if len(body.PhoneNumber) != 13 && !govalidator.IsNumeric(body.PhoneNumber) {
		err = errors.New("invalid phone number")
		_ = e.HandleError(c, err, h.log, http.StatusBadRequest, "Register")
		return
	}

	if !e.ValidatePassword(body.Password) {
		err = errors.New("invalid password")
		_ = e.HandleError(c, err, h.log, http.StatusBadRequest, "Register")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	existsPhone, err := h.serviceManager.UserService().UserService().CheckField(ctx, &pb.CheckFieldUserReq{
		Field: "phone_number",
		Value: body.PhoneNumber,
	})

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "Register") {
		return
	}
	if existsPhone.Status {
		err = errors.New("failed to check phone number uniques")
		_ = e.HandleError(c, err, h.log, http.StatusBadRequest, "Register")
		return
	}

	// TODO A method that sends a code to a number
	body.Code = 7777

	byteDate, err := json.Marshal(&body)
	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "Register") {
		return
	}

	err = rdb.Set(ctx, body.PhoneNumber, byteDate, time.Minute*h.ContextTimeout).Err()
	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "Register") {
		return
	}

	c.JSON(http.StatusOK, model_user_service.MessageRes{
		Message: "Code has been sent to you phone number, please check.",
	})
}

// Verify ...
// @Summary Verify
// @Description Authorization - Api for registering users
// @Tags Register
// @Accept json
// @Produce json
// @Param Verify body model_user_service.Verify true "RegisterModelReq"
// @Failure 200 {object} model_user_service.Verify
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/auth/verify [post]
func (h *HandlerV1) Verify(c *gin.Context) {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	var (
		body        model_user_service.Verify
		user        model_user_service.Redis
		jsonMarshal protojson.MarshalOptions
	)

	jsonMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "Verify") {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	if len(body.PhoneNumber) != 13 && !govalidator.IsNumeric(body.PhoneNumber) {
		err = errors.New("invalid phone number")
		_ = e.HandleError(c, err, h.log, http.StatusBadRequest, "Register")
		return
	}

	redisRes, err := rdb.Get(ctx, body.PhoneNumber).Result()

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "Verify") {
		return
	}

	err = json.Unmarshal([]byte(redisRes), &user)

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "Verify") {
		return
	}

	if body.Code != user.Code {
		err = errors.New("invalid code")
		_ = e.HandleError(c, err, h.log, http.StatusBadRequest, "Register")
		return
	}

	err = rdb.Del(ctx, body.PhoneNumber).Err()

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "Verify") {
		return
	}

	//sessions, err := h.serviceManager.SessionService().SessionService().GetUserSessions(ctx, &ps.StrUserReq{
	//	UserId: user.Id,
	//})
	//
	//if e.HandleError(c, err, h.log, http.StatusInternalServerError, "Verify") {
	//	return
	//}

	//if sessions.Count >= 3 {
	//	err = errors.New("the number of devices has exceeded the limit")
	//	_ = e.HandleError(c, err, h.log, http.StatusBadRequest, "Verify")
	//}

	//sessionId := uuid.New().String()
	//
	//_, err = h.serviceManager.SessionService().SessionService().CreateSession(ctx, &ps.SessionRequests{
	//	Id:           sessionId,
	//	IpAddress:    c.RemoteIP(),
	//	UserId:       user.Id,
	//	FcmToken:     body.FcmToken,
	//	PlatformName: body.PlatformName,
	//	PlatformType: body.PlatformType,
	//})
	//
	//if e.HandleError(c, err, h.log, http.StatusInternalServerError, "Verify") {
	//	return
	//}

	user.Id = uuid.New().String()

	user.Password, err = e.HashPassword(user.Password)
	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "Verify") {
		return
	}

	access, refresh, err := h.jwthandler.GenerateAuthJWT(user.PhoneNumber, user.Id, "user")

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "Verify") {
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

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "Verify") {
		return
	}

	c.JSON(http.StatusOK, &model_user_service.Response{
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		BrithDate:   user.BrithDate,
		PhoneNumber: user.PhoneNumber,
		Gender:      user.Gender,
		AccessToken: access,
	})

}

// ForgerPassword ...
// @Summary ForgerPassword
// @Description ForgerPassword - Api for registering users
// @Tags Register
// @Accept json
// @Produce json
// @Param phone_number query string true "phone_number"
// @Success 200 {object} model_user_service.MessageRes
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/auth/forger_password [get]
func (h *HandlerV1) ForgerPassword(c *gin.Context) {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	phoneNumber := c.Query("phone_number")

	if len(phoneNumber) != 13 && !govalidator.IsNumeric(phoneNumber) {
		err := errors.New("invalid phone number")
		_ = e.HandleError(c, err, h.log, http.StatusBadRequest, "Register")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	// TODO A method that sends a code to a number
	code := 7777

	err := rdb.Set(ctx, phoneNumber, code, time.Minute*h.ContextTimeout).Err()
	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "Register") {
		return
	}

	c.JSON(http.StatusOK, model_user_service.MessageRes{
		Message: "Code has been sent to you phone number, please check.",
	})
}

// ForgerPasswordVerify ...
// @Summary ForgerPasswordVerify
// @Description ForgerPasswordVerify - Api for registering users
// @Tags Register
// @Accept json
// @Produce json
// @Param Verify body model_user_service.ForgetPasswordVerify true "RegisterModelReq"
// @Failure 200 {object} model_user_service.MessageRes
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/auth/forger_password_verify [post]
func (h *HandlerV1) ForgerPasswordVerify(c *gin.Context) {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	var (
		body        model_user_service.ForgetPasswordVerify
		code        int
		jsonMarshal protojson.MarshalOptions
	)

	jsonMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	if len(body.PhoneNumber) != 13 && !govalidator.IsNumeric(body.PhoneNumber) {
		err = errors.New("invalid phone number")
		_ = e.HandleError(c, err, h.log, http.StatusBadRequest, "Register")
		return
	}

	if !e.ValidatePassword(body.NewPassword) {
		err = errors.New("invalid password")
		_ = e.HandleError(c, err, h.log, http.StatusBadRequest, "Register")
		return
	}

	redisRes, err := rdb.Get(ctx, body.PhoneNumber).Result()

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "Verify") {
		return
	}

	err = rdb.Del(ctx, body.PhoneNumber).Err()

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "Verify") {
		return
	}

	err = json.Unmarshal([]byte(redisRes), &code)

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "Verify") {
		return
	}

	if body.Code != code {
		err = errors.New("invalid code")
		_ = e.HandleError(c, err, h.log, http.StatusBadRequest, "Register")
		return
	}

	body.NewPassword, err = e.HashPassword(body.NewPassword)
	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "Verify") {
		return
	}

	_, err = h.serviceManager.UserService().UserService().ChangePassword(ctx, &pb.ChangeUserPasswordReq{
		PhoneNumber: body.PhoneNumber,
		Password:    body.NewPassword,
	})

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "Verify") {
		return
	}

	c.JSON(http.StatusOK, &model_user_service.MessageRes{Message: "Password changed successfully"})

}

// Login ...
// @Summary Login
// @Description Login - Api for registering users
// @Tags Register
// @Accept json
// @Produce json
// @Param Login body model_user_service.LoginReq true "Login Req"
// @Success 200 {object} model_user_service.Response
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/auth/login [post]
func (h *HandlerV1) Login(c *gin.Context) {
	var (
		body        model_user_service.LoginReq
		jsonMarshal protojson.MarshalOptions
	)

	jsonMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "Verify") {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	if len(body.PhoneNumber) != 13 && !govalidator.IsNumeric(body.PhoneNumber) {
		err := errors.New("invalid phone number")
		_ = e.HandleError(c, err, h.log, http.StatusBadRequest, "Register")
		return
	}

	if !e.ValidatePassword(body.Password) {
		err := errors.New("invalid password")
		_ = e.HandleError(c, err, h.log, http.StatusBadRequest, "Register")
		return
	}

	user, err := h.serviceManager.UserService().UserService().Get(ctx, &pb.GetUserReq{
		Field:    "phone_number",
		Value:    body.PhoneNumber,
		IsActive: false,
	})

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "Verify") {
		return
	}

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "Verify") {
		return
	}

	if !e.CheckHashPassword(user.Password, body.Password) {
		err = errors.New("invalid password")
		_ = e.HandleError(c, err, h.log, http.StatusBadRequest, "Register")
		return
	}

	//sessions, err := h.serviceManager.SessionService().SessionService().GetUserSessions(ctx, &ps.StrUserReq{
	//	UserId: user.Id,
	//})
	//
	//if e.HandleError(c, err, h.log, http.StatusInternalServerError, "Verify") {
	//	return
	//}
	//
	//if sessions.Count >= 3 {
	//	err = errors.New("the number of devices has exceeded the limit")
	//	_ = e.HandleError(c, err, h.log, http.StatusBadRequest, "Verify")
	//}
	//
	//sessionId := uuid.New().String()
	//
	//_, err = h.serviceManager.SessionService().SessionService().CreateSession(ctx, &ps.SessionRequests{
	//	Id:           sessionId,
	//	IpAddress:    c.RemoteIP(),
	//	UserId:       user.Id,
	//	FcmToken:     body.FcmToken,
	//	PlatformName: body.PlatformName,
	//	PlatformType: body.PlatformType,
	//})

	//if e.HandleError(c, err, h.log, http.StatusInternalServerError, "Verify") {
	//	return
	//}

	access, _, err := h.jwthandler.GenerateAuthJWT(user.PhoneNumber, user.Id, "user")

	// TODO Update Refresh Token by User Id

	c.JSON(http.StatusOK, model_user_service.Response{
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		BrithDate:   user.BirthDate,
		PhoneNumber: user.PhoneNumber,
		Gender:      user.Gender,
		AccessToken: access,
	})
}
