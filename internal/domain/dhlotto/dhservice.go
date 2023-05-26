package dhlotto

import (
	"fmt"
	"math/rand"

	"github.com/naoki914/lotter/internal/domain"
)

type DHService struct{}

func (*DHService) LottoSolution(draws []domain.Draw) ([]int, []int, error) {

	dhDraws := make([]*DHDraw, len(draws))

	for i, draw := range draws {

		// fmt.Printf("%+v\n", draw)
		dhDraw, ok := draw.(*DHDraw)
		if !ok {
			return nil, nil, fmt.Errorf("invalid draw type")
		}
		dhDraws[i] = dhDraw
	}

	freq := getSingleFrequency(dhDraws)
	choices := getWeightedChoices(freq)

	// fmt.Printf("%+v\n", freq)

	result := drawSolution(choices)
	return result, make([]int, 0), nil

}

func getSingleFrequency(draws []*DHDraw) map[int]int {
	f := make(map[int]int)

	for _, draw := range draws {
		f[draw.Draw1]++
		f[draw.Draw2]++
		f[draw.Draw3]++
		f[draw.Draw4]++
		f[draw.Draw5]++
		f[draw.Draw6]++
	}
	return f
}

func getWeightedChoices(freq map[int]int) []int {
	result := make([]int, 0)
	for k, v := range freq {
		for i := 0; i < v; i++ {
			result = append(result, k)
		}
	}
	return result
}

func drawOne(choices []int, avoid []int) int {

	for {
		done := true
		ix := rand.Intn(len(choices))

		for _, r := range avoid {
			if choices[ix] == r {
				done = false
				break
			}
		}
		if done {
			return choices[ix]
		}
	}

}

func drawSolution(choices []int) []int {
	result := make([]int, 6)
	for i := 0; i < 6; i++ {
		result[i] = drawOne(choices, result)
	}
	return result
}
