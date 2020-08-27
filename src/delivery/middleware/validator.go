package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"https:/src/types"
	"https:/src/utils"
	"net/http"
)

// ValidatorMiddleware used to validate users
// By their roles
type ValidatorMiddleware interface {
	ValidatorRoles(roles ...string) gin.HandlerFunc
}

// Validator struct
type Validator struct {}

// NewValidator return pointer to validator struct
// which includes all validate methods
func NewValidator() *Validator {
	return &Validator{}
}

// ValidateRoles takes roles enums and validate each role
// For the upcoming request, aborts the request
// If role wasn't found in validRoles array
func (v *Validator) ValidateRoles(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var validRoles []string
		claims, _ := Passport().GetClaimsFromJWT(c)

		userID := claims["id"]
		user, _ := userRepo.GeByKey("id", fmt.Sprint("%v", userID))
		status := utils.DerefString(user.Status)

		if status == types.StatusTypesEnum.Deleted {
			utils.CreateError(http.StatusForbidden, "user was deleted", c)
			c.Abort()
			return
		}

		for _, role := range roles {
			if user.Role == role {
				validRoles = append(validRoles, role)
			}
		}

		if len(validRoles) == 0 {
			utils.CreateError(http.StatusForbidden, "no permissons", c)
			c.Abort()
			return
		}
		c.Set("user", user)
		c.Next()
	}
}