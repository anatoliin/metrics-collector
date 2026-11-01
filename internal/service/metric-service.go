package service

import (
	models "github.com/anatoliin/metrics-collector/internal/model"
)

type MetricRepository interface {
	FindByID(id string) (*models.Metrics, error)
	Update(value *models.Metrics) error
}

type MetricService struct {
	repo MetricRepository
}

func NewMetricService(repo MetricRepository) *MetricService {
	return &MetricService{repo}
}

func (s *MetricService) GetMetricByID(id string) (*models.Metrics, error) {
	return s.repo.FindByID(id)
}

func (s *MetricService) UpdateMetric(value *models.Metrics) error {
	return s.repo.Update(value)
}
