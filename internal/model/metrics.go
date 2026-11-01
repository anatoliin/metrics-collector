package model

import (
	"fmt"
	"strconv"
)

const (
	Counter = "counter"
	Gauge   = "gauge"
)

// NOTE: Не усложняем пример, вводя иерархическую вложенность структур.
// Органичиваясь плоской моделью.
// Delta и Value объявлены через указатели,
// что бы отличать значение "0", от не заданного значения
// и соответственно не кодировать в структуру.
type Metrics struct {
	ID    string   `json:"id"`
	MType string   `json:"type"`
	Delta *int64   `json:"delta,omitempty"`
	Value *float64 `json:"value,omitempty"`
	Hash  string   `json:"hash,omitempty"`
}

func NewMetric(metricType, name, valueString string) (*Metrics, error) {
	if metricType == Counter {
		value, err := strconv.ParseInt(valueString, 10, 64)
		if err != nil {
			return nil, err // handle invalid input
		}
		return &Metrics{ID: name, MType: Counter, Delta: &value}, nil
	}
	if metricType == Gauge {
		value, err := strconv.ParseFloat(valueString, 64)
		if err != nil {
			return nil, err // handle invalid input
		}
		return &Metrics{ID: name, MType: Gauge, Value: &value}, nil
	}
	return nil, fmt.Errorf("unknown metric type: %s", metricType)
}
