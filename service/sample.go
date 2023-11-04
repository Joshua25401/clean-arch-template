package service

import "github.com/gin-gonic/gin"

func (s *templateService) SampleProcess(ginCtx *gin.Context) (map[string]any, error) {
	return map[string]any{
		"message": "success",
	}, nil
}
