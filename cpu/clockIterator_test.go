package cpu

import (
	"day10/pars"
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	inputs := []pars.Input{{ExecutionTime: 1, Value: 0},
		{ExecutionTime: 2, Value: 3},
		{ExecutionTime: 2, Value: -5}}
	cpu := InitClockTest(inputs)
	lastValue := 0
	for cpu.HasNext() {
		lastValue = cpu.GetNextCycle()
	}
	if lastValue != -1 {
		t.Errorf("Wrong value at end of short program , is %d , should be %d", lastValue, -1)
	}
}

func TestMediumInput(t *testing.T) {
	inputIter := pars.InitInputIter("../input_medium.txt")
	cpu := InitClock(inputIter)
	cykleCount := 1
	for cykleCount <= 220 {
		fmt.Printf("cycle: %d \t value: %d \n", cykleCount, cpu.GetNextCycle())
		cykleCount++
	}

}
