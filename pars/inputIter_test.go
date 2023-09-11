package pars

import (
	"testing"
)

func TestInputPars(t *testing.T) {
	in := InitInputIter("../input_short.txt")
	b1 := in.HasNext()
	if !b1 {
		t.Errorf("Can't find first input")
	}
	i1 := in.GetNext()
	exp1 := Input{command: "noop", value: 0}
	if i1 != exp1 {
		t.Errorf("Wrong first input")
	}
	b2 := in.HasNext()
	if !b2 {
		t.Errorf("Can't find second input")
	}
	i2 := in.GetNext()
	exp2 := Input{command: "addx", value: 3}
	if i2 != exp2 {
		t.Errorf("Wrong Second input")
	}
	b3 := in.HasNext()
	if !b3 {
		t.Errorf("Can't find third input")
	}
	i3 := in.GetNext()
	exp3 := Input{command: "addx", value: -5}
	if i3 != exp3 {
		t.Errorf("Wrong Third input")
	}
	b4 := in.HasNext()
	if b4 {
		t.Errorf("Found impossible input")
	}
}
