package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/anatoliin/metrics-collector/internal/model"
	"github.com/anatoliin/metrics-collector/internal/repository"
	"github.com/anatoliin/metrics-collector/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func TestMetricHandler_ServeHTTP(t *testing.T) {
	tests := []struct {
		name           string
		path           string
		wantStatusCode int
	}{
		{
			name:           "updates counter",
			path:           "/update/counter/requests/42",
			wantStatusCode: http.StatusOK,
		},
		{
			name:           "updates counter to zero",
			path:           "/update/counter/requests/0",
			wantStatusCode: http.StatusOK,
		},
		{
			name:           "updates gauge",
			path:           "/update/gauge/temperature/21.5",
			wantStatusCode: http.StatusOK,
		},
		{
			name:           "updates gauge to zero",
			path:           "/update/gauge/temperature/0",
			wantStatusCode: http.StatusOK,
		},
		{
			name:           "rejects unknown metric type",
			path:           "/update/histogram/requests/42",
			wantStatusCode: http.StatusBadRequest,
		},
		{
			name:           "rejects invalid counter value",
			path:           "/update/counter/requests/not-an-int",
			wantStatusCode: http.StatusBadRequest,
		},
		{
			name:           "rejects invalid gauge value",
			path:           "/update/gauge/temperature/not-a-float",
			wantStatusCode: http.StatusBadRequest,
		},
		{
			name:           "returns not found for an unmatched path",
			path:           "/update/counter/requests",
			wantStatusCode: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			storage := make(map[string]*model.Metrics)
			repo := repository.NewMemoryRepository(storage)
			metricService := service.NewMetricService(repo)

			r := chi.NewRouter()
			r.Method(http.MethodPost, "/update/{type}/{name}/{value}", NewMetricHandler(metricService))

			req := httptest.NewRequest(http.MethodPost, tt.path, nil)
			recorder := httptest.NewRecorder()
			r.ServeHTTP(recorder, req)

			assert.Equal(t, tt.wantStatusCode, recorder.Code)
		})
	}
}
