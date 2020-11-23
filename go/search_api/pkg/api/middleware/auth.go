package middleware

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"time"
)

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// User keeps information about authorized user. It is created at login event and after that
// can be accessed during requests for identifying user.
type User struct {
	UserName string
}

// AuthMiddleware responsible for processing JWT operations like login, refresh token, token is valid, etc.
// For more information check gin-jwt documentation: https://github.com/appleboy/gin-jwt
func AuthMiddleware(jwtSecret string) (*jwt.GinJWTMiddleware, error) {
	return jwt.New(&jwt.GinJWTMiddleware{
		Key:        []byte(jwtSecret),
		Timeout:    15 * time.Minute,
		MaxRefresh: 15 * time.Minute,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{
					jwt.IdentityKey: v.UserName,
				}
			}

			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)

			return &User{
				UserName: claims[jwt.IdentityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals login

			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			userID := loginVals.Username
			password := loginVals.Password

			if (userID == "admin" && password == "admin") || (userID == "test" && password == "test") {
				return &User{
					UserName: userID,
				}, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
	})
}
