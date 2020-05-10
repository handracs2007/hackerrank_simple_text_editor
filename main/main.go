// https://www.hackerrank.com/challenges/simple-text-editor/problem

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type StringStack struct {
	data string
}

func NewStringStack() *StringStack {
	var stringStack = &StringStack{
		data: "",
	}

	return stringStack
}

func (stack *StringStack) Push(elem string) {
	stack.data += elem
}

func (stack *StringStack) Pop() (string, error) {
	if len(stack.data) > 0 {
		var lastIndex = len(stack.data) - 1

		var value = stack.data[lastIndex]
		stack.data = stack.data[0:lastIndex]

		return string(value), nil
	}

	return "", errors.New("stack is empty")
}

func (stack *StringStack) Peek(index int) (string, error) {
	if index < len(stack.data) {
		return string(stack.data[index]), nil
	}

	return "", errors.New("stack is empty")
}

func (stack *StringStack) Reinitialise(newData string) {
	stack.data = newData
}

func (stack *StringStack) String() string {
	return stack.data
}

type StringArrayStack struct {
	data []string
}

func NewStringArrayStack() *StringArrayStack {
	var stringArrayStack = &StringArrayStack{
		data: make([]string, 0),
	}

	return stringArrayStack
}

func (stack *StringArrayStack) Push(elem string) {
	stack.data = append(stack.data, elem)
}

func (stack *StringArrayStack) Pop() (string, error) {
	if len(stack.data) > 0 {
		var lastIndex = len(stack.data) - 1

		var value = stack.data[lastIndex]
		stack.data = stack.data[0:lastIndex]

		return value, nil
	}

	return "", errors.New("stack is empty")
}

func (stack *StringArrayStack) Peek(index int) (string, error) {
	if index < len(stack.data) {
		return stack.data[index], nil
	}

	return "", errors.New("stack is empty")
}

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var numOfCasesStr, _ = reader.ReadString('\n')
	var numOfCases, _ = strconv.Atoi(strings.TrimSpace(numOfCasesStr))

	var stringStack = NewStringStack()
	var undoStack = NewStringArrayStack()

	for i := 0; i < numOfCases; i++ {
		var commandLine, _ = reader.ReadString('\n')
		var commands = strings.Split(strings.TrimSpace(commandLine), " ")
		var command, _ = strconv.Atoi(commands[0])

		switch command {
		case 1:
			// Append command
			var param = commands[1]
			undoStack.Push(stringStack.String())
			stringStack.Push(param)

		case 2:
			// Delete command
			undoStack.Push(stringStack.String())

			var count, _ = strconv.Atoi(commands[1])
			for j := 0; j < count; j++ {
				_, _ = stringStack.Pop()
			}

		case 3:
			// Print command
			var loc, _ = strconv.Atoi(commands[1])
			var value, _ = stringStack.Peek(loc - 1)
			fmt.Println(value)

		case 4:
			// Undo command
			var prevData, _ = undoStack.Pop()
			stringStack.Reinitialise(prevData)
		}
	}
}
