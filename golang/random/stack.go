package main

import (
	"fmt"
)

type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

func (s *Stack) Push(str string) string {
	orgLength := len(*s)
	*s = append(*s, str)

	if len(*s) == orgLength+1 {
		return "YES"
	}
	return "FAILED"
}

func (s *Stack) PrintStack() {
	for _, item := range *s {
		fmt.Println(item)
	}
}

func main() {

	var stack Stack
	fmt.Println("Start stack")
	fmt.Println(stack.Push("Test"))
	fmt.Println(stack.Push("Test2"))
	fmt.Println(stack.Push("Test3"))
	fmt.Println(stack.Push("Test4"))
	fmt.Println(stack.Push("Test5"))
	fmt.Println(stack.Pop())
	stack.PrintStack()

}
