package handler

import (
	"net/http"

	"github.com/anatoliin/metrics-collector/internal/model"
	"github.com/anatoliin/metrics-collector/internal/service"
	"github.com/go-chi/chi/v5"
)

type MetricHandler struct {
	metricService *service.MetricService
}

func NewMetricHandler(service *service.MetricService) *MetricHandler {
	return &MetricHandler{service}
}

func (h *MetricHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	typeParam := chi.URLParam(r, "type")
	nameParam := chi.URLParam(r, "name")
	valueParam := chi.URLParam(r, "value")
	newMetric, err := model.NewMetric(typeParam, nameParam, valueParam)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	err = h.metricService.UpdateMetric(newMetric)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
