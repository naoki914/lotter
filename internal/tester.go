package tester

import (
	"fmt"
)

type Draw interface {
	GetID() string
}

type DHDraw struct {
	ID int
}

func (draw *DHDraw) GetID() string {
	return fmt.Sprintf("dhlottery-%d", draw.ID)
}

func runit() {
	var di Draw

	dhd := DHDraw{}

	di = &dhd

	fmt.Printf("%d\n", di)
}
