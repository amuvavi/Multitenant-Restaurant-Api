package middlewares

import (
	"time"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth_gin"
	"github.com/gin-gonic/gin"
)

func RateLimiter() gin.HandlerFunc {
	// Creating a new rate limiter with a limit of 1 request per second
	limiter := tollbooth.NewLimiter(1, nil)

	// Setting additional configuration
	limiter.SetHeaderEntryExpirationTTL(time.Hour)
	limiter.SetBurst(5)
	limiter.SetMethods([]string{"GET", "POST", "PUT", "DELETE"})
	limiter.SetIPLookups([]string{"RemoteAddr", "X-Forwarded-For", "X-Real-IP"})
	return tollbooth_gin.LimitHandler(limiter)
}
