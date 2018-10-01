package tracking

import (
	"math"
	"strconv"
)

func Tracking(input [][]string) float64 {
	var total float64
	for row := range input {
		for day := range input[row] {
			if len(input[row][day]) > 1 {
				listData := input[row][day]
				price, _ := strconv.ParseFloat(listData[1:], 64)
				total += price
			}
		}

	}
	return math.Round((total/7)*100) / 100
}
