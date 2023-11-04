package delivery

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (d *delivery) SampleGinHandler(ginCtx *gin.Context) {
	data, err := d.service.SampleProcess(ginCtx)
	if err != nil {
		// Do something here
		log.Printf("ERROR: %w", err)
	}

	ginCtx.JSON(http.StatusOK, data)
}
