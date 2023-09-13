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
	scanner     *bufio.Scanner
	haveScanned bool
}

func InitInputIter(filename string) InputIter {
	fileReader, _ := os.Open(filename)
	file_scanner := bufio.NewScanner(fileReader)
	return InputIter{scanner: file_scanner, haveScanned: true}
}

func (i *InputIter) HasNext() bool { // det här blev bara dumt

	return i.haveScanned
}

func (i *InputIter) GetNext() Input {
	i.haveScanned = i.scanner.Scan()
	raw := i.scanner.Text()
	if strings.Contains(raw, " ") {
		splited := strings.Split(raw, " ")
		value, _ := strconv.Atoi(splited[1])
		instructionTime := 2
		return Input{ExecutionTime: instructionTime, Value: value}
	}
	return Input{ExecutionTime: 1, Value: 0}
}
