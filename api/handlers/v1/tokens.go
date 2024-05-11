package v1

import (
	e "dennic_api_gateway/api/handlers/regtool"
	"dennic_api_gateway/api/models/model_user_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetTokens
// @Summary GetTokens
// @Description GetTokens
// @Tags Token
// @Accept json
// @Produce json
// @Success 200 {object} model_user_service.Tokens "Successful response"
// @Failure 400 {object} model_common.StandardErrorModel
// @Failure 500 {object} model_common.StandardErrorModel
// @Router /v1/token/get-token [GET]
func (h *HandlerV1) GetTokens(c *gin.Context) {

	access, refresh, err := h.jwthandler.GenerateAuthJWT("1234567890", "123e4567-e89b-12d3-a456-426614174001", "1", "user")
	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "GetTokens") {
		return
	}
	c.JSON(http.StatusOK, model_user_service.Tokens{
		AccessToken:  access,
		RefreshToken: refresh,
	})
}
