package status

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetHandler returns an HTTP handler to perform a status check.
func GetHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
	}
}
