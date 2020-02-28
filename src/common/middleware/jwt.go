package middleware

import (
	"time"
	"tripyuk_api_go/src/module/model"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var identityKey = "id"

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func DefaultMW(db *gorm.DB) (*jwt.GinJWTMiddleware, error) {
	// the jwt middleware
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour * 24 * 30,
		MaxRefresh:  time.Hour * 24 * 30,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*model.User); ok {
				return jwt.MapClaims{
					identityKey: v.Email,
					"user":      v.ID,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &model.User{
				Name: claims[identityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			email := loginVals.Username
			password := loginVals.Password

			var user model.User

			if err := db.Where("email = ?", email).Find(&user).Error; err == nil {
				match := CheckPasswordHash(password, user.Password)
				if match {
					return &model.User{
						Name:  user.Name,
						Email: user.Email,
					}, nil
				}
			}
			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			return true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

	return authMiddleware, err
}

func HelloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	c.JSON(200, gin.H{
		"email": claims[identityKey],
		"text":  "Hello World.",
	})
}
