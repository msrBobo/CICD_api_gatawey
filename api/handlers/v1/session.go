package v1

import (
	"context"
	e "dennic_api_gateway/api/handlers/regtool"
	"dennic_api_gateway/api/models"
	"dennic_api_gateway/api/models/model_session_service"
	pb "dennic_api_gateway/genproto/session_service"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// GetUserSessions ...
// @Summary GetUserSessions
// @Description GetUserSessions - Api for get session
// @Tags Session
// @Accept json
// @Produce json
// @Param user_id query string true "user_id"
// @Success 200 {object} model_session_service.ListSessions
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/session [get]
func (h *HandlerV1) GetUserSessions(c *gin.Context) {
	id := c.Query("user_id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	sessions, err := h.serviceManager.SessionService().SessionService().GetUserSessions(ctx, &pb.StrUserReq{
		UserId:   id,
		IsActive: false,
	})

	if e.HandleError(c, err, h.log, http.StatusBadRequest, "GetUserSessions") {
		return
	}

	var sessionsRes model_session_service.ListSessions
	for _, session := range sessions.UserSessions {
		sessionsRes.Sessions = append(sessionsRes.Sessions, &model_session_service.SessionRes{
			Id:           session.Id,
			Order:        session.Order,
			IpAddress:    session.IpAddress,
			UserId:       session.UserId,
			FcmToken:     session.FcmToken,
			PlatformName: session.PlatformName,
			PlatformType: session.PlatformType,
			LoginAt:      session.LoginAt,
			CreatedAt:    session.CreatedAt,
			UpdatedAt:    e.UpdateTimeFilter(session.UpdatedAt),
		})
	}
	c.JSON(http.StatusOK, model_session_service.ListSessions{
		Sessions: sessionsRes.Sessions,
		Count:    sessions.Count,
	})
}

// DeleteSessionById ...
// @Summary DeleteSessionById
// @Description DeleteSessionById - Api for delete session
// @Tags Session
// @Accept json
// @Produce json
// @Param id query string true "id"
// @Success 200 {object} models.StatusRes
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/session [delete]
func (h *HandlerV1) DeleteSessionById(c *gin.Context) {
	id := c.Query("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	_, err := h.serviceManager.SessionService().SessionService().DeleteSessionById(ctx, &pb.StrReq{
		Id: id,
	})

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "DeleteSessionById") {
		return
	}

	c.JSON(http.StatusOK, models.StatusRes{Status: true})
}
