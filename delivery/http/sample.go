package delivery

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (d *delivery) SampleGinHandler(ginCtx *gin.Context) {
	data, err := d.service.SampleProcess(ginCtx)
	if err != nil {
		// Do something here
		d.logger.Error(ginCtx.Request.Context(), "SampleGinHandler error", err)
		return
	}

	d.logger.Info(context.Background(), "SampleGinHandler Info", zap.Any("message", "Success"))
	ginCtx.JSON(http.StatusOK, data)
}
