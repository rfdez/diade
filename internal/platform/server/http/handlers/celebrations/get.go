package celebrations

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rfdez/diade/internal/fetching"
	"github.com/rfdez/diade/kit/errors"
	"github.com/rfdez/diade/kit/query"
)

type celebrationResponse struct {
	ID     string `json:"id"`
	Date   string `json:"date"`
	Name   string `json:"name"`
	Status string `json:"status"`
	Type   string `json:"type"`
}

// GetHandler returns an HTTP handler to perform health checks.
func GetHandler(queryBus query.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resp, err := queryBus.Ask(ctx, fetching.NewCelebrationByDateQuery(ctx.Query("date")))
		if err != nil {
			if errors.IsWrongInput(err) {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"message": err.Error(),
				})
				return
			}

			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		celebrations, ok := resp.([]fetching.CelebrationResponse)
		if !ok {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		var response = make([]celebrationResponse, 0, len(celebrations))
		for _, c := range celebrations {
			response = append(response, celebrationResponse{
				ID:     c.ID(),
				Date:   c.Date(),
				Name:   c.Name(),
				Status: c.Status(),
				Type:   c.Type(),
			})
		}

		ctx.JSON(http.StatusOK, response)
	}
}
