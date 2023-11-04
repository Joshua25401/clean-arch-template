package delivery

import (
	"clean-arch-template/service"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type (
	HTTPDependencies struct {
		Service service.Service
	}

	delivery struct {
		// Depend to service layer
		service service.Service
		// Depend to authentication layer
	}
)

func NewDelivery(dep HTTPDependencies) http.Handler {
	delivery := &delivery{
		service: dep.Service,
	}

	router := gin.New()

	// Use Middleware here
	router.Use(cors.New(CORSConfig()))

	// Register handler here
	router.GET("/", delivery.SampleGinHandler)

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
