package middleware

import (
	"user-auth/utils/auth"
	"user-auth/utils/errorh"

	"github.com/gin-gonic/gin"
)

// TokenAuthMiddleware validate request by token
func TokenAuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {

		// extract token and check if valid and not expired
		authDitail, err := auth.ExtractTokenAuth(c.Request, secret)
		if err != nil {

			// return not authorized error
			errh := errorh.NotAuthorizedError("You don't have access to this resource")
			c.JSON(errh.Code, errh)

			// stop gin from running other handlers
			c.Abort()
			return
		}

		// set user id to context to use later
		c.Set("user_id", authDitail.UserId)
		c.Next()
	}
}
