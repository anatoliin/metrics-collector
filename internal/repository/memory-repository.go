package repository

import (
	"sync"

	models "github.com/anatoliin/metrics-collector/internal/model"
)

type MemoryRepository struct {
	mu      sync.RWMutex
	storage map[string]*models.Metrics
}

func NewMemoryRepository(storage map[string]*models.Metrics) *MemoryRepository {
	return &MemoryRepository{storage: storage}
}

func (h *MemoryRepository) FindByID(id string) (*models.Metrics, error) {
	h.mu.RLock()
	defer h.mu.RUnlock()
	if val, ok := h.storage[id]; ok {
		return val, nil
	}

	return nil, ErrNotFound
}

func (h *MemoryRepository) Update(value *models.Metrics) error {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.storage[value.ID] = value
	return nil
}
