package service

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (s *templateService) SampleProcess(ginCtx *gin.Context) (map[string]any, error) {
	s.log.Info(ginCtx.Request.Context(), "got request", zap.String("client_ip", ginCtx.GetString("client_ip")))
	return map[string]any{
		"message": "success",
	}, nil
}
