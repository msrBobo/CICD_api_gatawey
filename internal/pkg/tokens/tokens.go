package token

import (
	"dennic_api_gateway/internal/pkg/config"
	"github.com/dgrijalva/jwt-go"
	"go.uber.org/zap"
	"time"
)

// JWTHandler ...
type JWTHandler struct {
	Id        string
	Phone     string
	SessionId string
	Exp       string
	Iat       string
	Aud       []string
	Role      string
	SignKey   string
	Log       *zap.Logger
	Token     string
	Timout    time.Duration
}

type CustomClaims struct {
	*jwt.Token
	SessionId string   `json:"session_id"`
	UserName  string   `json:"user_name"`
	Id        string   `json:"id"`
	Exp       float64  `json:"exp"`
	Iat       float64  `json:"iat"`
	Aud       []string `json:"aud"`
	Role      string   `json:"role"`
}

// GenerateAuthJWT ...
func (jwtHandler *JWTHandler) GenerateAuthJWT(phone, id, sessionId, role string) (access, refresh string, err error) {
	var (
		accessToken  *jwt.Token
		refreshToken *jwt.Token
		claims       jwt.MapClaims
		rtClaims     jwt.MapClaims
	)

	jwtHandler.Timout = 360

	accessToken = jwt.New(jwt.SigningMethodHS256)
	refreshToken = jwt.New(jwt.SigningMethodHS256)
	claims = accessToken.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["phone"] = phone
	claims["session_id"] = sessionId
	claims["exp"] = time.Now().Add(time.Minute * jwtHandler.Timout).Unix()
	claims["iat"] = time.Now().Unix()
	claims["role"] = role
	access, err = accessToken.SignedString([]byte(config.SignKey))
	if err != nil {
		jwtHandler.Log.Log(1, err.Error())
		return
	}

	rtClaims = refreshToken.Claims.(jwt.MapClaims)
	rtClaims["id"] = id
	rtClaims["phone"] = phone
	rtClaims["session_id"] = sessionId
	rtClaims["exp"] = time.Now().Add(time.Minute * jwtHandler.Timout).Unix()
	rtClaims["iat"] = time.Now().Unix()
	rtClaims["role"] = role
	refresh, err = refreshToken.SignedString([]byte(config.SignKey))
	if err != nil {
		jwtHandler.Log.Log(1, err.Error())
		return
	}
	return access, refresh, nil
}

// GenerateJWT ...
func (jwtHandler *JWTHandler) GenerateJWT(phone, id, role string) (access string, err error) {
	var (
		accessToken *jwt.Token
		claims      jwt.MapClaims
	)

	jwtHandler.Timout = 60

	accessToken = jwt.New(jwt.SigningMethodHS256)
	claims = accessToken.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["phone"] = phone
	claims["exp"] = time.Now().Add(time.Minute * jwtHandler.Timout).Unix()
	claims["iat"] = time.Now().Unix()
	claims["role"] = role
	access, err = accessToken.SignedString([]byte(config.SignKey))
	if err != nil {
		jwtHandler.Log.Log(1, err.Error())
		return
	}
	return access, nil
}

// ExtractClaims ...
func (jwtHandler *JWTHandler) ExtractClaims() (jwt.MapClaims, error) {
	var (
		token *jwt.Token
		err   error
	)

	token, err = jwt.Parse(jwtHandler.Token, func(t *jwt.Token) (interface{}, error) {
		return []byte(jwtHandler.SignKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		return nil, err
	}
	return claims, nil
}

// ExtractClaim extracts claims from given token
func ExtractClaim(tokenStr string) (jwt.MapClaims, error) {
	var (
		token *jwt.Token
		err   error
	)
	token, err = jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return []byte(config.SignKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !(ok && token.Valid) {
		return nil, err
	}
	return claims, nil
}
