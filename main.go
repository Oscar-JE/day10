package main

import (
	"day10/cpu"
	"day10/pars"
	"fmt"
)

func main() {
	println(part1())
	part2()
}

func part1() int {
	inputIter := pars.InitInputIter("input.txt")
	cpu := cpu.InitClock(inputIter)
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
		println(registerValue)
		currentCycle++
	}
	return sum
}

func part2() {
	inputIter := pars.InitInputIter("input.txt")
	cpu := cpu.InitClock(inputIter)
	currentPixel := 1
	for cpu.HasNext() {
		spritePositon := cpu.GetNextCycle()
		if pixelInSprite(spritePositon, currentPixel) {
			fmt.Printf("#")
		} else {
			fmt.Printf(".")
		}
		if currentPixel%40 == 0 {
			fmt.Printf("\n")
			currentPixel = 0
		}
		currentPixel++
	}
}

func pixelInSprite(spritePositon int, currentPixel int) bool {
	return spritePositon <= currentPixel && currentPixel < spritePositon+3
}
