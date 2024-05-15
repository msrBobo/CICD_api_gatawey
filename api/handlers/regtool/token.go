package v1

import (
	jwt "dennic_api_gateway/internal/pkg/tokens"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type UserTokenRes struct {
	UserId      string
	PhoneNumber string
	SessionId   string
	Role        string
}

func GetUserInfo(c *gin.Context) (*UserTokenRes, error) {
	token := c.GetHeader("Authorization")

	claims, err := jwt.ExtractClaim(token)
	fmt.Println(err, "----------")
	if err != nil {

		return nil, err
	}

	return &UserTokenRes{
		UserId:      cast.ToString(claims["id"]),
		SessionId:   cast.ToString(claims["session_id"]),
		PhoneNumber: cast.ToString(claims["phone"]),
		Role:        cast.ToString(claims["role"]),
	}, nil

}
