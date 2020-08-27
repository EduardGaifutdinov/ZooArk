package middleware

import (
	"errors"
	"github.com/ZooArk/src/repository"
	"github.com/ZooArk/src/schemes/request"
	"github.com/ZooArk/src/types"
	"github.com/ZooArk/src/utils"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

// IdentityKeyID is used to tell
// by what field we will identify user
const IdentityKeyID = "id"

// UserID struct
type UserID struct {
	ID string
}

var userRepo = repository.NewUserRepo()

// Passport is middleware for user authentication
func Passport() *jwt.GinJWTMiddleware {
	authMiddleware, _ := jwt.New(&jwt.GinJWTMiddleware{
		Realm: "Zoo",
		Key: []byte(os.Getenv("JWTSECRET")),
		Timeout: time.Hour * 4,
		MaxRefresh: time.Hour * 25,
		IdentityKey: IdentityKeyID,
		SendCookie: true,
		CookieName: "jwt",
		TokenLookup: "cookie:jwt",
		LoginResponse: func(c *gin.Context, i int, s string, t time.Time) {
			value, _ := Passport().ParseTokenString(s)
			id := jwt.ExtractClaimsFromToken(value)["id"]
			result, err := userRepo.GetByID(id.(string))

			if err != nil {
				utils.CreateError(http.StatusUnauthorized, err.Error(), c)
				return
			}

			status := utils.DerefString(result.Status)

			if status == types.StatusTypesEnum.Deleted {
				utils.CreateError(http.StatusForbidden, "user was deleted", c)
				return
			}

			if status == types.StatusTypesEnum.Invited {
				code, err := userRepo.UpdateStatus(result.ID, types.StatusTypesEnum.Active)
				if err != nil {
					utils.CreateError(code, err.Error(), c)
					return
				}
			}

			c.JSON(http.StatusOK, result)
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*UserID); ok {
				return jwt.MapClaims{
					IdentityKeyID: v.ID,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &UserID{
				ID: claims[IdentityKeyID].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var body request.LoginUserRequest
			if err := c.ShouldBind(&body); err != nil {
				return "", errors.New("missing email or password")
			}

			result, err := userRepo.GeByKey("email", body.Email)
			if err == nil {
				status := utils.DerefString(result.Status)
				if status == types.StatusTypesEnum.Deleted {
					return nil, errors.New("user was deleted")
				}
				equal := utils.CheckPasswordHash(body.Password, result.Password)
				if equal {
					return &UserID{
						ID: result.ID.String(),
					}, nil
				}
			}
			return nil, errors.New("incorrect email or password")
		},
		Unauthorized: func (c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code": code,
				"message": message,
			})
		},
	})
	return authMiddleware
}