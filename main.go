package main

import (
	"blah/betterStack"
	"blah/stack"
	"fmt"
)

func main() {
	stack.MyStack.Push("Hey there!")
	stack.MyStack.Push("How are you?")
	stack.MyStack.Push("Hey there!")
	stack.MyStack.Push("How are you?")
	stack.MyStack.PrintSize()

	fmt.Println("Popping: ", stack.MyStack.Pop())
	stack.MyStack.PrintSize()
	fmt.Println("Popping: ", stack.MyStack.Pop())
	stack.MyStack.PrintSize()

	betterStack.MyStack.Push("Hey there!")
	betterStack.MyStack.Push("How are you?")
	betterStack.MyStack.Push("How are you?")
	betterStack.MyStack.PrintSize()

	fmt.Println("Popping: ", betterStack.MyStack.Pop())
	betterStack.MyStack.PrintSize()
	fmt.Println("Popping: ", betterStack.MyStack.Pop())
	betterStack.MyStack.PrintSize()
}
