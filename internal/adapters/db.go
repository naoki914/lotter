package adapters

import (
	"github.com/naoki914/lotter/internal/domain"
)

type DB interface {
	Create(domain.Draw) error
	Get(int) (*domain.Draw, error)
	GetAll(string) ([]domain.Draw, error)
}
