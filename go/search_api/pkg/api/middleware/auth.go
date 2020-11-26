// Package middleware package provide middleware for Gin framework like auth, correlation id, cors and recover.
package middleware

import (
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

const (
	tokenTimeout    = 15 // minutes
	tokenMaxRefresh = 15 // minutes
)

// User keeps information about authorized user. It is created at login event and after that
// can be accessed during requests for identifying user.
type User struct {
	Username string
}

// AuthMiddleware responsible for processing JWT operations like login, refresh token, token is valid, etc.
// For more information check gin-jwt documentation: https://github.com/appleboy/gin-jwt
func AuthMiddleware(jwtSecret string) (*jwt.GinJWTMiddleware, error) {
	return jwt.New(&jwt.GinJWTMiddleware{
		Key:             []byte(jwtSecret),
		Timeout:         tokenTimeout * time.Minute,
		MaxRefresh:      tokenMaxRefresh * time.Minute,
		PayloadFunc:     payloadFunc,
		IdentityHandler: identityHandler,
		Authenticator:   authenticator,
	})
}

func authenticator(c *gin.Context) (interface{}, error) {
	request := struct {
		Username string `form:"username" json:"username" binding:"required"`
		Password string `form:"password" json:"password" binding:"required"`
	}{}

	if err := c.ShouldBind(&request); err != nil {
		return "", jwt.ErrMissingLoginValues
	}

	user := request.Username
	password := request.Password

	if (user == "admin" && password == "admin") || (user == "test" && password == "test") {
		return &User{
			Username: user,
		}, nil
	}

	return nil, jwt.ErrFailedAuthentication
}

// payloadFunc converts User to jwt claims.
func payloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(*User); ok {
		return jwt.MapClaims{
			jwt.IdentityKey: v.Username,
		}
	}

	return jwt.MapClaims{}
}

// identityHandler converts jwt claims to User.
func identityHandler(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)

	return &User{
		Username: claims[jwt.IdentityKey].(string),
	}
}
