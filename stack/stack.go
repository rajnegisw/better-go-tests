package stack

import "fmt"

var MyStack Stack = Stack{}

type Block struct {
	payload  string
	belowPtr *Block
}

type Stack struct {
	head *Block
	size int
}

func CreateBlock(payload string, blockPtr *Block) *Block {
	b := &Block{payload: payload, belowPtr: blockPtr}
	return b
}

func FreeUpBlock(block *Block) string {
	payload := block.payload
	MyStack.head = block.belowPtr

	return payload
}

func (s *Stack) Push(payload string) {
	newBlock := CreateBlock(payload, s.head)
	(*s).size += 1
	(*s).head = newBlock
}

func (s *Stack) Pop() string {

	payload := FreeUpBlock(s.head)
	MyStack.size -= 1
	return payload
}

func (s Stack) PrintSize() {
	fmt.Println("Stack size is: ", s.size)
}
