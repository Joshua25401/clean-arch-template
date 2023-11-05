package delivery

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

func RealIP() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		realUserIP := ctx.ClientIP()

		// set ip to context
		ctx.Set("client_ip", realUserIP)

		// continue to other handler / middleware
		ctx.Next()
	}
}

func IPRateLimiter(bucket Bucket, rate int64, per time.Duration) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userIpFromContext := ctx.GetString("client_ip")

		if b, exist := bucket[userIpFromContext]; exist {
			if b.TakeAvailable(1) == 0 {
				ctx.JSON(http.StatusTooManyRequests, gin.H{
					"message": "Too many request",
				})
				ctx.Abort()
				return
			}
		} else {
			newBucket := ratelimit.NewBucket(per, rate)
			bucket[userIpFromContext] = newBucket
		}

		ctx.Next()
	}
}
