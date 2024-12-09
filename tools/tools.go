package tools

import (
	"bufio"
	"os"
	"strings"
)

type SortList interface {
	SortNumber() ([]float64, error)
	SortBytes() ([]string, error)
}

type randomList struct {
	list []string
}

func Input() (*randomList, string) {
	var (
		InputMass []string
	)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	values := strings.Split(input, " ")
	for _, value := range values {
		InputMass = append(InputMass, value)
	}
	InterfaceMass := &randomList{list: InputMass}
	InputMassElement := InputMass[0]
	return InterfaceMass, InputMassElement
}
