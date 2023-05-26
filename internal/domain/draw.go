package domain

type Draw interface {
	GetID() string
	GetSold() int
	GetWinPrize() int
	GetPrimaryNumbers() []int
	GetSecondaryNumbers() []int
	IsSuccess() bool
}
