package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type httpResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func responseData(ginCtx *gin.Context, status int, response httpResponse) {
	ginCtx.JSON(status, response)
}

func responseError(ginCtx *gin.Context, err error) {
	ginCtx.JSON(
		http.StatusInternalServerError,
		httpResponse{
			Message: "Internal Server Error",
			Data:    err.Error(),
		},
	)
}
