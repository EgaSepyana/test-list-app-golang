package middleware

import (
	"errors"
	"log"
	"os"
	"time"

	"todolist-api/src/config"
	"todolist-api/src/model"
	"todolist-api/src/service"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

var identityKey = "id"

func InitJwt() (authMiddleware *jwt.GinJWTMiddleware, err error) {
	authMiddleware, err = jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "Geo Sense",
		Key:         []byte(os.Getenv(config.ENV_KEY_JWT_SECRET)),
		Timeout:     24 * time.Hour,
		MaxRefresh:  48 * time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*model.User); ok {
				return jwt.MapClaims{
					identityKey: v.IdDocument,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &model.User{
				Username: claims[identityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (res interface{}, err error) {
			var param model.User_Login
			if err := c.ShouldBind(&param); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			resp, userRes, ok := model.Response{Metadata: model.MetadataResponse{}}, model.User{}, false
			userService := service.NewUserService()
			if param.Username != "" && param.Password != "" {
				resp, ok = userService.GetUserWithValidatePassword("username",
					param.Username, param.Password, &userRes)
			}
			if ok {
				c.Set("user", userRes)
				res = &userRes
				return
			} else {
				return nil, errors.New(resp.Metadata.Message)
			}
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			return true

			//? This will limit only superadmin can access the API
			// if v, ok := data.(*model.User); ok && v.Username == "superadmin" {
			// 	return true
			// }
			// return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
		LoginResponse: func(c *gin.Context, code int, message string, time time.Time) {
			if userRaw, found := c.Get("user"); found {
				user, _ := userRaw.(model.User)
				c.JSON(code, model.Resp_User_Login{
					User: &user,
					Auth: model.Resp_JwtToken{
						Token:  message,
						Expire: time,
					},
				})
			}
		},
	})
	if err != nil {
		log.Println("JWT Error:" + err.Error())
	}

	return
}

func SetUserInContextByJwt() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		claims := jwt.ExtractClaims(ctx)
		log.Println(claims)
		ctx.Set("uid", claims["id"])
	}
}
