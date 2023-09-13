package cpu

import (
	"day10/pars"
	"testing"
)

func TestShortest(t *testing.T) {
	input := pars.InitInputIter("../input_short.txt")
	cpu := InitClock(input)
	for cpu.HasNext() {
		println(cpu.GetNext())
	}
}
