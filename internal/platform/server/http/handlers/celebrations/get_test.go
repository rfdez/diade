package celebrations_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rfdez/diade/internal/fetching"
	"github.com/rfdez/diade/internal/platform/server/http/handlers/celebrations"
	"github.com/rfdez/diade/kit/query/querymocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestHandler_Get(t *testing.T) {
	queryBus := new(querymocks.Bus)
	queryBus.On(
		"Ask",
		mock.Anything,
		mock.AnythingOfType("fetching.CelebrationByDateQuery"),
	).Return([]fetching.CelebrationResponse{}, nil)

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/celebrations", celebrations.GetHandler(queryBus))

	t.Run("given a valid request it returns 200", func(t *testing.T) {
		req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "/celebrations", http.NoBody)
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusOK, res.StatusCode)
	})
}
