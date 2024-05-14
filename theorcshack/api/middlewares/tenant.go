package middlewares

import (
	"theorcshack/helpers"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func TenantMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tenantID := c.GetHeader("X-Tenant-ID")
		if tenantID == "" {
			helpers.RespondWithError(c, 400, "X-Tenant-ID header is required")
			c.Abort()
			return
		}

		tenantUUID, err := uuid.Parse(tenantID)
		if err != nil {
			helpers.RespondWithError(c, 400, "Invalid X-Tenant-ID header")
			c.Abort()
			return
		}

		c.Set("tenantID", tenantUUID)
		c.Next()
	}
}
