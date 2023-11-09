package delivery

import (
	pkg "clean-arch-template/pkg/logger"
	"clean-arch-template/service"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

type (
	Bucket           map[string]*ratelimit.Bucket
	HTTPDependencies struct {
		Service  service.Service
		Logger   pkg.Logger
		IpBucket Bucket
	}

	delivery struct {
		// Depend to service layer
		service service.Service

		// Depend on logger
		logger pkg.Logger

		// Rate limit bucket
		bucket Bucket
		// Depend to authentication layer
	}
)

func NewDelivery(dep HTTPDependencies) http.Handler {
	delivery := &delivery{
		service: dep.Service,
		logger:  dep.Logger,
		bucket:  dep.IpBucket,
	}

	router := gin.New()

	// Use Middleware here
	router.Use(RealIP())
	router.Use(IPRateLimiter(delivery.bucket, 80, time.Minute))
	router.Use(cors.New(CORSConfig()))

	// Register handler here
	router.GET("/", delivery.SampleGinHandler)
	router.GET("/sample-error", delivery.SampleError)

	return router
}

func CORSConfig() cors.Config {
	return cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTION"},
		AllowHeaders:     []string{},
		AllowCredentials: true,
		MaxAge:           86400,
	}
}
