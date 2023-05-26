package dhlotto

import (
	"fmt"
)

type DHDraw struct {
	ID             int    `json:"drwNo" bson:"_id"`
	Sold           int    `json:"totSellamnt" bson:"sold"`
	WinPrize       int    `json:"firstWinamnt" bson:"winPrize"`
	WinTotalAmount int    `json:"firstAccumamnt" bson:"winTotalAmount"`
	Winners        int    `json:"firstPrzwnerCo" bson:"winners"`
	DateStr        string `json:"drwNoDate" bson:"dateStr"`
	Draw1          int    `json:"drwtNo1" bson:"draw1"`
	Draw2          int    `json:"drwtNo2" bson:"draw2"`
	Draw3          int    `json:"drwtNo3" bson:"draw3"`
	Draw4          int    `json:"drwtNo4" bson:"draw4"`
	Draw5          int    `json:"drwtNo5" bson:"draw5"`
	Draw6          int    `json:"drwtNo6" bson:"draw6"`
	DrawBonus      int    `json:"bnusNo" bson:"drawbonus"`
	Success        string `json:"returnValue" bson:"success"`
}

func (draw *DHDraw) GetID() string {
	return fmt.Sprintf("dhlottery-%d", draw.ID)
}

func (draw *DHDraw) GetSold() int {
	return draw.Sold
}

func (draw *DHDraw) GetWinPrize() int {
	return draw.WinPrize
}

func (draw *DHDraw) GetPrimaryNumbers() []int {
	numbers := make([]int, 6)
	numbers[0] = draw.Draw1
	numbers[1] = draw.Draw2
	numbers[2] = draw.Draw3
	numbers[3] = draw.Draw4
	numbers[4] = draw.Draw5
	numbers[5] = draw.Draw6
	return numbers
}

func (draw *DHDraw) GetSecondaryNumbers() []int {
	return nil
}
func (draw *DHDraw) IsSuccess() bool {
	return draw.Success == "success"
}
