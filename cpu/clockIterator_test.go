package cpu

import (
	"day10/pars"
	"testing"
)

/*
func TestShortest(t *testing.T) {
	input := pars.InitInputIter("../input_short.txt")
	cpu := InitClock(input)
	for cpu.HasNext() {
		println(cpu.GetNextCycle())
	}
}*/

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
