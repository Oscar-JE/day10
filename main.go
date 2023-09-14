package main

import (
	"day10/cpu"
	"day10/pars"
	"fmt"
)

func main() {
	println(part1())
}

func part1() int {
	inputIter := pars.InitInputIter("input.txt")
	cpu := cpu.InitClock(inputIter)
	//Vilka värden var intressanta
	importantCycles := []int{20, 60, 100, 140, 180, 220}
	importantCyclesIndex := 0
	currentCycle := 1
	sum := 0
	for cpu.HasNext() {
		registerValue := cpu.GetNextCycle()
		if importantCyclesIndex < len(importantCycles) && currentCycle == importantCycles[importantCyclesIndex] {
			importantCyclesIndex++
			sum += registerValue * currentCycle
			fmt.Printf("cycle %d , registervalue %d , sum %d  \n", currentCycle, registerValue, sum)
		}
		//println(registerValue)
		currentCycle++
	}
	return sum
}
