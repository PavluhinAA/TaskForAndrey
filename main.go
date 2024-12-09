package main

import "TaskForAndrey/tools"

func main() {
	InterfaceMass, InputMassElement := tools.Input()
	tools.ChoosingAsort(InterfaceMass, InputMassElement)
}
