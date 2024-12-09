package tools

import (
	"errors"
	"fmt"
	"strconv"
)

func (rl *randomList) SortBytes() []string {
	n := len(rl.list)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if len(rl.list[j]) > len(rl.list[j+1]) {
				rl.list[j], rl.list[j+1] = rl.list[j+1], rl.list[j]
			}
		}
	}
	return rl.list
}
func (rL *randomList) SortNumber() ([]float64, error) {
	fmt.Println(rL.list)
	floats := make([]float64, 0)
	for _, str := range rL.list {
		floatVal, err := strconv.ParseFloat(str, 64)
		if err == nil {
			floats = append(floats, floatVal)
		} else {
			err = errors.New("error converting to a floating point number")
			return floats, err
		}
	}
	fmt.Println(floats)
	n := len(floats)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if floats[j] > floats[j+1] {
				floats[j], floats[j+1] = floats[j+1], floats[j]
			}
		}
	}
	return floats, nil
}
func ChoosingAsort(InterfaceMass *randomList, InputMassElement string) {
	var (
		sortedNums []float64
	)
	_, err := strconv.Atoi(InputMassElement)
	if err == nil {
		sortedNums, err = InterfaceMass.SortNumber()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Sorted list:", sortedNums)
	} else {
		sortedNums := InterfaceMass.SortBytes()
		fmt.Println("Sorted list:", sortedNums)
	}
}
