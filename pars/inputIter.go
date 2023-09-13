package pars

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type CommandItr interface {
	GetNext() Input
	HasNext() bool
}

type Input struct {
	ExecutionTime int
	Value         int
}

func (i Input) Eq(other Input) bool {
	return i.ExecutionTime == other.ExecutionTime && i.Value == other.Value
}

type InputIter struct {
	scanner     *bufio.Scanner
	haveScanned bool
}

func InitInputIter(filename string) InputIter {
	fileReader, _ := os.Open(filename)
	file_scanner := bufio.NewScanner(fileReader)
	return InputIter{scanner: file_scanner, haveScanned: false}
}

func (i *InputIter) HasNext() bool { // det här blev bara dumt
	if i.haveScanned {
		return true
	} else {
		i.haveScanned = i.scanner.Scan()
		return i.haveScanned
	}

	//return i.scanner.Scan()
}

func (i *InputIter) GetNext() Input {
	raw := i.scanner.Text()
	i.haveScanned = false
	if strings.Contains(raw, " ") {
		splited := strings.Split(raw, " ")
		value, _ := strconv.Atoi(splited[1])
		instructionTime := 2
		return Input{ExecutionTime: instructionTime, Value: value}
	}
	return Input{ExecutionTime: 1, Value: 0}
}
