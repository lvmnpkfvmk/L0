package main

//Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5. Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	tempGroups := make(map[int][]float64)

	for _, temp := range temperatures {
		var tmp int
		if temp < 0 {
			tmp = int(math.Ceil(temp/10) * 10)
		} else {
			tmp = int(math.Floor(temp/10) * 10)
		}
		tempGroups[tmp] = append(tempGroups[tmp], temp)
	}

	keys := make([]int, 0, len(tempGroups))
	for key := range tempGroups {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	for _, key := range keys {
		fmt.Printf("%d: %v, ", key, tempGroups[key]) // -20: [-25.4 -27 -21], 10: [13 19 15.5], 20: [24.5], 30: [32.5]
	}
	fmt.Println()
}
