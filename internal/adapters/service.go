package adapters

import "github.com/naoki914/lotter/internal/domain"

type Service interface {
	LottoSolution([]domain.Draw) ([]int, []int, error)
}
