package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type sort interface {
	quickSort()
	formatting() []string
}

type stringSlice struct {
	slice []rune
}

type floatSlice struct {
	slice []float64
}

func main() {

	inputString, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	inputString = strings.Trim(inputString, "\n")
	inputSlice := strings.Split(inputString, " ")

	sortingSlice := typeDefinition(inputSlice)

	sortingSlice.quickSort()

	output := sortingSlice.formatting()
	fmt.Println(strings.Join(output, " "))

}

func typeDefinition(slice []string) sort {

	resultFloat := make([]float64, 0)
	resultString := make([]rune, 0)

	if _, err := strconv.ParseFloat(slice[0], 64); err == nil {

		for _, element := range slice {

			number, _ := strconv.ParseFloat(element, 64)
			resultFloat = append(resultFloat, number)

		}

		return &floatSlice{resultFloat}

	} else {

		for _, str := range slice {

			resultString = append(resultString, []rune(str)...)

		}

		return &stringSlice{resultString}

	}

}

func (f *floatSlice) quickSort() {

	if len(f.slice) <= 0 {
		return
	}

	var (
		smaller, largest []float64
		supportElement   = f.slice[0]
	)

	for _, element := range f.slice[1:] {

		if element <= supportElement {

			smaller = append(smaller, element)

		} else {

			largest = append(largest, element)

		}

	}

	newSmaller := floatSlice{smaller}
	newLargest := floatSlice{largest}

	newSmaller.quickSort()
	newLargest.quickSort()

	f.slice = append(append(newSmaller.slice, supportElement), newLargest.slice...)

}

func (s *stringSlice) quickSort() {

	if len(s.slice) <= 0 {
		return
	}

	var (
		smaller, largest []rune
		supportElement   = s.slice[0]
	)

	for _, element := range s.slice[1:] {

		if element <= supportElement {

			smaller = append(smaller, element)

		} else {

			largest = append(largest, element)

		}

	}

	newSmaller := stringSlice{smaller}
	newLargest := stringSlice{largest}

	newSmaller.quickSort()
	newLargest.quickSort()

	s.slice = append(append(newSmaller.slice, supportElement), newLargest.slice...)

}

func (f *floatSlice) formatting() []string {

	outputSlice := make([]string, 0)

	for i, number := range f.slice {

		formatNumber := strconv.FormatFloat(number, 'g', 10, 64)
		outputSlice[i] = formatNumber

	}

	return outputSlice

}

func (s *stringSlice) formatting() []string {

	outputSlice := make([]string, 0)

	for i, element := range s.slice {

		outputSlice[i] = string(element)

	}

	return outputSlice

}
