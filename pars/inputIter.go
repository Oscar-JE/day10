package pars

import (
	"bufio"
	"fmt"
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

func (i Input) String() string {
	return fmt.Sprintf("Edecution time: %d ; Value: %d", i.ExecutionTime, i.Value)
}

type InputIter struct {
	inputs []Input
	index  int
}

func InitInputIter(filename string) InputIter {
	fileReader, _ := os.Open(filename)
	file_scanner := bufio.NewScanner(fileReader)
	inputs := []Input{}
	for file_scanner.Scan() {
		line := file_scanner.Text()
		inputs = append(inputs, pars(line))
	}
	return InputIter{inputs: inputs, index: 0}
}

func (i InputIter) HasNext() bool {
	return i.index < len(i.inputs)
}

func (i *InputIter) GetNext() Input {
	retVal := i.inputs[i.index]
	i.index++
	return retVal
}

func pars(line string) Input {
	if strings.Contains(line, " ") {
		splited := strings.Split(line, " ")
		value, _ := strconv.Atoi(splited[1])
		instructionTime := 2
		return Input{ExecutionTime: instructionTime, Value: value}
	}
	return Input{ExecutionTime: 1, Value: 0}
}
