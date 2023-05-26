package adapters

import "github.com/naoki914/lotter/internal/domain"

type Api interface {
	FetchDrawWithID(int) (domain.Draw, error)
}
