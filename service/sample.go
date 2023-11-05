package service

import (
	"github.com/gin-gonic/gin"
)

func (s *templateService) SampleProcess(ginCtx *gin.Context) (map[string]any, error) {
	s.log.Info(ginCtx.Request.Context(), "got request")
	return map[string]any{
		"message": "success",
	}, nil
}
