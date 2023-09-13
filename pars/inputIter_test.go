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
	exp1 := Input{ExecutionTime: 1, Value: 0}
	if i1 != exp1 {
		t.Errorf("Wrong first input")
	}
	b2 := in.HasNext()
	if !b2 {
		t.Errorf("Can't find second input")
	}
	i2 := in.GetNext()
	exp2 := Input{ExecutionTime: 2, Value: 3}
	if i2 != exp2 {
		t.Errorf("Wrong Second input")
	}
	b3 := in.HasNext()
	if !b3 {
		t.Errorf("Can't find third input")
	}
	i3 := in.GetNext()
	exp3 := Input{ExecutionTime: 2, Value: -5}
	if i3 != exp3 {
		t.Errorf("Wrong Third input")
	}
	b4 := in.HasNext()
	if b4 {
		t.Errorf("Found impossible input")
	}
}

func TestMultipleHasNext(t *testing.T) {
	// hasnext should not change the state
	in := InitInputIter("../input_short.txt")
	i := 0
	hasNext := true
	for i < 20 {
		hasNext = hasNext && in.HasNext()
		i++
	}
	if !hasNext {
		t.Errorf("HasNext is changing the state of the scanner")
	}
}

func TestShort(t *testing.T) {
	in := InitInputIter("../input_short.txt")
	for in.HasNext() {
		println(in.GetNext().String())
	}
}
