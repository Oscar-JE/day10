package pars

type TestIter struct {
	index    int
	inputArr []Input
}

func (t *TestIter) HasNext() bool {
	return t.index < len(t.inputArr)
}

func (t *TestIter) GetNext() Input {
	var retValue Input = t.inputArr[t.index]
	t.index++
	return retValue
}

func InitTestInput(inputs []Input) TestIter {
	return TestIter{index: 0, inputArr: inputs}
}
