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

func calculate(operation string, slice []float64) float64 {
	var result float64
	switch operation {
	case SUM:
		for _, v := range slice {
			result += v
		}
	case AVG:
		var sum float64
		for _, v := range slice {
			sum += v
		}
		result = sum / float64(len(slice))
	case MED:
		sortedSlice := sort.Float64Slice(slice)
		if len(sortedSlice)%2 == 0 {
			result += (sortedSlice[len(sortedSlice)/2] + sortedSlice[len(sortedSlice)/2+1]) / 2
		} else {
			result = slice[len(slice)/2]
		}
	}

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
