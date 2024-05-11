package v1

import (
	"context"
	e "dennic_api_gateway/api/handlers/regtool"
	"dennic_api_gateway/api/models/model_user_service"
	pb "dennic_api_gateway/genproto/user_service"
	"dennic_api_gateway/internal/pkg/logger"
	jwt "dennic_api_gateway/internal/pkg/tokens"
	"net/http"
	"time"

	"github.com/spf13/cast"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// GetUserByID
// @Summary GetUserByID
// @Description Api for GetUserByID
// @Security ApiKeyAuth
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} model_user_service.GetUserResp
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/user/get [GET]
func (h *HandlerV1) GetUserByID(c *gin.Context) {

	token := c.GetHeader("Authorization")
	claims, err := jwt.ExtractClaim(token)
	if e.HandleError(c, err, h.log, http.StatusUnauthorized, "GetUserByID") {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	response, err := h.serviceManager.UserService().UserService().Get(
		ctx, &pb.GetUserReq{
			Field:    "id",
			Value:    cast.ToString(claims["id"]),
			IsActive: false,
		})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get user", logger.Error(err))
		return
	}

	resp := model_user_service.GetUserResp{
		Id:          response.Id,
		UserOrder:   response.UserOrder,
		FirstName:   response.FirstName,
		LastName:    response.LastName,
		BrithDate:   response.BirthDate,
		PhoneNumber: response.PhoneNumber,
		Password:    response.Password,
		Gender:      response.Gender,
		CreatedAt:   response.CreatedAt,
		UpdatedAt:   response.UpdatedAt,
	}

	c.JSON(http.StatusOK, resp)

}

// ListUsers
// @Summary ListUsers
// @Description Api for ListUsers
// @Tags User
// @Accept json
// @Produce json
// @Param Page  query string false "Page"
// @Param Limit query string false "Limit"
// @Param Field query string false "Field"
// @Param Value query string false "Value"
// @Param OrderBy query string false "OrderBy"
// @Success 200 {object} model_user_service.GetUserResp
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/user [GET]
func (h *HandlerV1) ListUsers(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	page := c.Query("Page")
	limit := c.Query("Limit")
	field := c.Query("Field")
	value := c.Query("Value")
	orderBy := c.Query("OrderBy")

	pageInt, limitInt, err := e.ParseQueryParams(page, limit)
	if e.HandleError(c, err, h.log, http.StatusBadRequest, "failed to list users") {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	response, err := h.serviceManager.UserService().UserService().ListUsers(
		ctx, &pb.ListUsersReq{
			Page:     pageInt,
			Limit:    limitInt,
			IsActive: false,
			Value:    value,
			Field:    field,
			OrderBy:  orderBy,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to list users", logger.Error(err))
		return
	}

	var users model_user_service.ListUserResp

	for _, in := range response.Users {
		user := model_user_service.GetUserResp{
			Id:          in.Id,
			UserOrder:   in.UserOrder,
			FirstName:   in.FirstName,
			LastName:    in.LastName,
			BrithDate:   in.BirthDate,
			PhoneNumber: in.PhoneNumber,
			Password:    in.Password,
			Gender:      in.Gender,
			CreatedAt:   in.CreatedAt,
			UpdatedAt:   in.UpdatedAt,
		}

		users.Users = append(users.Users, user)
		users.Count = response.Count
	}

	c.JSON(http.StatusOK, users)
}

// UpdateUser
// @Summary UpdateUser
// @Description Api for UpdateUser
// @Tags User
// @Accept json
// @Produce json
// @Param UpdUserReq body model_user_service.UpdUserReq true "UpdUserReq"
// @Success 200 {object} model_user_service.GetUserResp
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/user [PUT]
func (h *HandlerV1) UpdateUser(c *gin.Context) {
	var (
		body        model_user_service.UpdUserResp
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", logger.Error(err))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	req := &pb.User{
		Id:        body.Id,
		FirstName: body.FirstName,
		LastName:  body.LastName,
		BirthDate: body.BrithDate,
		Gender:    body.Gender,
	}

	response, err := h.serviceManager.UserService().UserService().Update(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to update user", logger.Error(err))
		return
	}

	resp := model_user_service.UpdUserResp{
		Id:        response.Id,
		FirstName: response.FirstName,
		LastName:  response.LastName,
		BrithDate: response.BirthDate,
		Gender:    response.Gender,
		UpdatedAt: response.UpdatedAt,
	}

	c.JSON(http.StatusOK, resp)
}

// UpdateRefreshToken
// @Summary Update Refresh Token
// @Description Update the refresh token of the user
// @Tags User
// @Accept json
// @Produce json
// @Param RefreshToken body model_user_service.RefreshToken true "RefreshToken"
// @Success 200 {object} model_user_service.UpdateRefreshTokenUserResp "Successful response"
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/user/update-refresh-token [PUT]
func (h *HandlerV1) UpdateRefreshToken(c *gin.Context) {

	var (
		RefreshToken model_user_service.RefreshToken
		jspbMarshal  protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&RefreshToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", logger.Error(err))
		return
	}

	claims, err := jwt.ExtractClaim(RefreshToken.RefreshToken)
	if e.HandleError(c, err, h.log, http.StatusUnauthorized, "UpdateRefreshToken") {
		return
	}

	access, refresh, err := h.jwthandler.GenerateAuthJWT(cast.ToString(claims["phone"]), cast.ToString(claims["id"]), cast.ToString(claims["session_id"]), "user")
	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "Login") {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	_, err = h.serviceManager.UserService().UserService().UpdateRefreshToken(ctx, &pb.UpdateRefreshTokenUserReq{
		Id:           cast.ToString(claims["id"]),
		RefreshToken: refresh,
	})

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "UpdateRefreshToken") {
		return
	}
	resp := model_user_service.UpdateRefreshTokenUserResp{
		AccessToken:  access,
		RefreshToken: refresh,
	}

	c.JSON(http.StatusOK, resp)
}

// DeleteUser
// @Summary DeleteUser
// @Description Api for DeleteUser
// @Tags User
// @Accept json
// @Produce json
// @Param Field query string true "Field"
// @Param Value query string true "Value"
// @Success 200 {object} model_user_service.CheckUserFieldResp
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/user [DELETE]
func (h *HandlerV1) DeleteUser(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	field := c.Query("Field")
	value := c.Query("Value")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	response, err := h.serviceManager.UserService().UserService().Delete(
		ctx, &pb.DeleteUserReq{
			Field:    field,
			Value:    value,
			IsActive: false,
		})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete user", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}
