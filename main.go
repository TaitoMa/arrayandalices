package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

const (
	SUM = "SUM"
	AVG = "AVG"
	MED = "MED"
)

func main() {
	operation := getOperation()
	slice := getSlice()
	fmt.Println(slice)
	result := calculate(operation, slice)

	fmt.Println(result)
}

var calculateMap = map[string]func([]float64) float64{
	"SUM": func(slice []float64) float64 {
		var result float64
		for _, v := range slice {
			result += v
		}
		return result
	},
	"AVG": func(slice []float64) float64 {
		var result float64
		var sum float64
		for _, v := range slice {
			sum += v
		}
		result = sum / float64(len(slice))
		return result
	},
	"MED": func(slice []float64) float64 {
		var result float64
		sortedSlice := sort.Float64Slice(slice)
		sortedSlice.Sort()
		if len(sortedSlice)%2 == 0 {
			result += (sortedSlice[len(sortedSlice)/2-1] + sortedSlice[len(sortedSlice)/2]) / 2
		} else {
			result = sortedSlice[len(slice)/2]
		}
		return result
	},
}

func calculate(operation string, slice []float64) float64 {
	result := calculateMap[operation](slice)

	return result
}

func getOperation() string {
	var operation string
	for {
		fmt.Println("Введите операцию (SUM/AVG/MED/)")
		_, err := fmt.Scan(&operation)
		if err != nil {
			continue
		}
		if operation != SUM && operation != AVG && operation != MED {
			fmt.Println("Такой операции нет")
			continue
		}
		break
	}
	return operation
}

func getSlice() []float64 {
	slice := []float64{}
	fmt.Println("Введите числа через запятую")
	var str string
	fmt.Scan(&str)
	trimmed := strings.TrimSpace(str)
	splitted := strings.Split(trimmed, ",")
	for _, s := range splitted {
		num, err := strconv.ParseFloat(s, 64)
		if err != nil {
			continue
		}
		slice = append(slice, num)
	}

	return slice
}

/*
	NOTES
*/
/*
Удалить 3й элемент
s = append(s[:2], s[3:]...)

*/
