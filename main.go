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

type symbolSlice []rune

type floatSlice []float64

func main() {

	inputString, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	inputString = strings.Trim(inputString, "\n")

	if inputString == "" {

		fmt.Println("no input detected")
		os.Exit(2)

	}

	inputSlice := strings.Split(inputString, " ")

	sortingSlice := typeDefinition(inputSlice)

	sortingSlice.quickSort()

	output := sortingSlice.formatting()
	fmt.Println("result: ", strings.Join(output, " "))

}

func typeDefinition(slice []string) sort {

	resultFloat := make(floatSlice, 0)
	resultString := make(symbolSlice, 0)

	if _, ok := strconv.ParseFloat(slice[0], 64); ok == nil {

		for _, element := range slice {

			number, errorTransformations := strconv.ParseFloat(element, 64)

			if errorTransformations != nil {
				fmt.Println("the symbols have different types")
				os.Exit(2)
			}

			resultFloat = append(resultFloat, number)

		}

		return &resultFloat

	} else {

		for _, str := range slice {

			resultString = append(resultString, []rune(str)...)

		}

		return &resultString

	}

}

func (f *floatSlice) quickSort() {

	if len(*f) <= 0 {
		return
	}

	smaller := make(floatSlice, 0)
	largest := make(floatSlice, 0)

	var (
		supportElement = (*f)[0]
	)

	for _, element := range (*f)[1:] {

		if element <= supportElement {

			smaller = append(smaller, element)

		} else {

			largest = append(largest, element)

		}

	}

	smaller.quickSort()
	largest.quickSort()

	*f = append(append(smaller, supportElement), largest...)

}

func (s *symbolSlice) quickSort() {

	if len(*s) <= 0 {
		return
	}

	smaller := make(symbolSlice, 0)
	largest := make(symbolSlice, 0)

	var (
		supportElement = (*s)[0]
	)

	for _, element := range (*s)[1:] {

		if element <= supportElement {

			smaller = append(smaller, element)

		} else {

			largest = append(largest, element)

		}

	}

	smaller.quickSort()
	largest.quickSort()

	*s = append(append(smaller, supportElement), largest...)

}

func (f *floatSlice) formatting() []string {

	outputSlice := make([]string, len(*f))

	for i, number := range *f {

		formatNumber := strconv.FormatFloat(number, 'g', 10, 64)
		outputSlice[i] = formatNumber

	}

	return outputSlice

}

func (s *symbolSlice) formatting() []string {

	outputSlice := make([]string, len(*s))

	for i, element := range *s {

		outputSlice[i] = string(element)

	}

	return outputSlice

}
