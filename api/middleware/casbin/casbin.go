package casbin

import (
	"dennic_api_gateway/internal/pkg/config"
	jwt "dennic_api_gateway/internal/pkg/tokens"
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type CasbinHandler struct {
	config     config.Config
	enforce    casbin.Enforcer
	jwthandler jwt.JWTHandler
}

func NewAuthorizer() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token1 := ctx.GetHeader("Authorization")
		if token1 == "" {

			sub := "unauthorized"
			obj := ctx.Request.URL.Path
			etc := ctx.Request.Method
			e, _ := casbin.NewEnforcer(`auth.conf`, `auth.csv`)
			t, _ := e.Enforce(sub, obj, etc)
			if t {
				ctx.Next()
				return
			}
			fmt.Println(sub, obj, etc, t)
		}

		claims, err := jwt.ExtractClaim(token1)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
			return
		}

		sub := claims["role"]
		obj := ctx.Request.URL.Path
		etc := ctx.Request.Method

		e, err := casbin.NewEnforcer(`auth.conf`, `auth.csv`)

		if err != nil {
			log.Fatal(err)
			return
		}
		t, err := e.Enforce(sub, obj, etc)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Println(sub, obj, etc)
		if t {
			ctx.Next()
			return
		}
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "permission denied",
		})
	}
}
