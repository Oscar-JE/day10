package cpu

import (
	"day10/pars"
)

type ClockIterator struct {
	inputIter     pars.CommandItr
	delayedAdd    DelayedAdd
	registerValue int
}

type DelayedAdd struct {
	valueToAdd int
	delay      int
}

func InitClock(input pars.InputIter) ClockIterator {
	registerValue := 1
	delayedAdd := DelayedAdd{valueToAdd: 0, delay: 0}
	return ClockIterator{inputIter: &input, delayedAdd: delayedAdd, registerValue: registerValue}
}

func InitClockTest(input []pars.Input) ClockIterator{
	var inputIter pars.TestIter = pars.InitTestInput(input)
	delayedAdd := DelayedAdd{valueToAdd: 0, delay: 0}
	return ClockIterator{inputIter: &inputIter, delayedAdd: delayedAdd, registerValue: 1}
}

func (c ClockIterator) HasNext() bool {
	haveInput := c.inputIter.HasNext()
	computesInput := c.delayedAdd.delay > 0
	return haveInput || computesInput
}

func (c *ClockIterator) GetNext() int {
	retValue := c.registerValue
	if c.delayedAdd.delay == 0 {
		c.registerValue += c.delayedAdd.valueToAdd //skicka runt operatorer med delay istället ??
		input := c.inputIter.GetNext()
		c.delayedAdd = DelayedAdd{valueToAdd: input.Value, delay: input.ExecutionTime}
	}
	c.delayedAdd.delay -= 1
	return retValue
}
