package service

import "github.com/gin-gonic/gin"

type Service interface {
	SampleProcess(ginCtx *gin.Context) (map[string]any, error)
}
