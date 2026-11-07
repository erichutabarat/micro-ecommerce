package middleware

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"api-gateway/internal/auth"
)

const (
	ctxUserID = "auth_user_id"
	ctxRole   = "auth_role"
)

// OptionalAuth identifies the caller if a valid Bearer token is present.
// It never rejects a request: routes that need a role are enforced in the handler.
// An invalid or expired token is treated the same as no token, so a client
// holding an expired token can still reach public routes such as login.
func OptionalAuth(v *auth.Verifier) gin.HandlerFunc {
	return func(c *gin.Context) {
		parts := strings.SplitN(c.GetHeader("Authorization"), " ", 2)
		if len(parts) == 2 && strings.EqualFold(parts[0], "Bearer") {
			if claims, err := v.Parse(parts[1]); err == nil {
				if id, err := strconv.ParseUint(claims.Subject, 10, 64); err == nil && id > 0 {
					c.Set(ctxUserID, uint(id))
					c.Set(ctxRole, claims.Role)
				}
			}
		}
		c.Next()
	}
}

// Identity returns the caller's ID and role, or (0, "") if unauthenticated.
func Identity(c *gin.Context) (uint, string) {
	id, _ := c.Get(ctxUserID)
	role, _ := c.Get(ctxRole)

	userID, _ := id.(uint)
	roleStr, _ := role.(string)
	return userID, roleStr
}
