package router

import (
	"net/http"

	"github.com/anatoliin/metrics-collector/internal/handler"
	"github.com/anatoliin/metrics-collector/internal/model"
	"github.com/anatoliin/metrics-collector/internal/repository"
	"github.com/anatoliin/metrics-collector/internal/service"
	"github.com/go-chi/chi/v5"
)

func NewRouter() http.Handler {
	repo := repository.NewMemoryRepository(make(map[string]*model.Metrics))
	metricService := service.NewMetricService(repo)
	r := chi.NewRouter()
	r.Method(http.MethodPost, "/update/{type}/{name}/{value}", handler.NewMetricHandler(metricService))
	return r
}
