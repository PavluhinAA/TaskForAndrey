package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

type sort interface {
	quickSort()
	output()
}

type byteSlice struct {
	slice []string
}

type floatSlice struct {
	slice []float64
}

func main() {
	inputSlice, err := input()
	if err != nil {
		fmt.Println("error reading input")
		os.Exit(2)
	}
	if inputSlice == nil {
		fmt.Println("No input detected")
		os.Exit(2)
	}
	sliceByte, sliceFloat := typeDefinition(inputSlice)
	if sliceFloat != nil {
		sliceFloat.quickSort()
		sliceFloat.output()
	}
	if sliceByte != nil {
		sliceByte.quickSort()
		sliceByte.output()
	}
}

func input() ([][]byte, error) {
	sliceByte := make([]byte, 0)
	finalSlice := make([][]byte, 0)
	inputSlice := bufio.NewReader(os.Stdin)
	for {
		input, err := inputSlice.ReadByte()
		if err != nil {
			if err == io.EOF {
				if len(sliceByte) > 0 {
					finalSlice = append(finalSlice, sliceByte)
				}
				break
			}
			return nil, err
		}
		if rune(input) == ' ' {
			copyByte := make([]byte, len(sliceByte))
			copy(copyByte, sliceByte)
			finalSlice = append(finalSlice, copyByte)
			sliceByte = sliceByte[:0]
			continue
		} else if rune(input) == '\n' {
			if len(finalSlice) == 0 {
				return nil, nil
			}
			finalSlice = append(finalSlice, sliceByte)
			break
		}
		sliceByte = append(sliceByte, input)
	}
	return finalSlice, nil
}

func typeDefinition(sliceInterface [][]byte) (sort, sort) {
	sliceFloat := make([]float64, 0)
	sliceByte := make([]string, 0)
	for i := 0; i < len(sliceInterface); i++ {
		if elementFloat, errFloat := strconv.ParseFloat(string(sliceInterface[i]), 64); errFloat == nil {
			sliceFloat = append(sliceFloat, elementFloat)
		} else {
			sliceByte = append(sliceByte, string(sliceInterface[i]))
		}
	}
	return &byteSlice{sliceByte}, &floatSlice{sliceFloat}
}

func (f *floatSlice) quickSort() {
	if len(f.slice) <= 1 {
		return
	}
	supportElement := f.slice[0]
	smallerSlice := 0
	largestSlice := len(f.slice) - 1

	for smallerSlice <= largestSlice {
		for f.slice[smallerSlice] < supportElement {
			smallerSlice++
		}
		for supportElement < f.slice[largestSlice] {
			largestSlice--
		}
		if smallerSlice <= largestSlice {
			f.slice[smallerSlice], f.slice[largestSlice] = f.slice[largestSlice], f.slice[smallerSlice]
			smallerSlice++
			largestSlice--
		}
	}
	if largestSlice > 0 {
		rightSlice := floatSlice{f.slice[:largestSlice+1]}
		rightSlice.quickSort()
	}
	if smallerSlice < len(f.slice) {
		leftSlice := floatSlice{f.slice[smallerSlice:]}
		leftSlice.quickSort()
	}
}

func (b *byteSlice) quickSort() {
	if len(b.slice) <= 1 {
		return
	}
	supportElement := b.slice[0]
	smallerSlice := 0
	largestSlice := len(b.slice) - 1

	for smallerSlice <= largestSlice {
		for utf8.RuneCountInString(b.slice[smallerSlice]) < utf8.RuneCountInString(supportElement) {
			smallerSlice++
		}
		for utf8.RuneCountInString(supportElement) < utf8.RuneCountInString(b.slice[largestSlice]) {
			largestSlice--
		}
		if smallerSlice <= largestSlice {
			b.slice[smallerSlice], b.slice[largestSlice] = b.slice[largestSlice], b.slice[smallerSlice]
			smallerSlice++
			largestSlice--
		}
	}
	if largestSlice > 0 {
		rightSlice := byteSlice{b.slice[:largestSlice+1]}
		rightSlice.quickSort()
	}
	if smallerSlice < len(b.slice) {
		leftSlice := byteSlice{b.slice[smallerSlice:]}
		leftSlice.quickSort()
	}
}

func (f *floatSlice) output() {
	outputSlice := make([]string, len(f.slice))
	for j, number := range f.slice {
		outputSlice[j] = strconv.FormatFloat(number, 'g', 10, 64)
	}
	fmt.Println(strings.Join(outputSlice, " "))
}

func (b *byteSlice) output() {
	fmt.Println(strings.Join(b.slice, " "))
}
