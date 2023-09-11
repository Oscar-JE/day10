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
	command string
	value   int
}

func (i Input) Eq(other Input) bool {
	return i.command == other.command && i.value == other.value
}

type InputIter struct {
	scanner *bufio.Scanner
}

func InitInputIter(filename string) InputIter {
	fileReader, _ := os.Open(filename)
	file_scanner := bufio.NewScanner(fileReader)
	return InputIter{file_scanner}
}

func (i InputIter) HasNext() bool {
	return i.scanner.Scan()
}

func (i InputIter) GetNext() Input {
	raw := i.scanner.Text()
	if strings.Contains(raw, " ") {
		splited := strings.Split(raw, " ")
		command := splited[0]
		value, _ := strconv.Atoi(splited[1])
		return Input{command: command, value: value}
	}
	return Input{command: raw, value: 0}
}
