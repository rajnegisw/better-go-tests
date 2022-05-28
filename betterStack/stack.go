package betterStack

import "fmt"

var MyStack Stack = Stack{blockFactory: &Block{}}

type BlockType interface {
	Create(string) BlockType
	FreeUp(BlockType) string
}

type Stack struct {
	head         BlockType
	size         int
	blockFactory BlockType
}

type Block struct {
	payload  string
	belowPtr *Block
}

func (block *Block) Create(payload string) BlockType {
	b := &Block{payload: payload, belowPtr: block}
	return b
}

func (block *Block) FreeUp(blockType BlockType) string {

	payload := block.payload
	MyStack.head = block.belowPtr

	return payload
}

func (s *Stack) Push(payload string) {
	newBlock := s.blockFactory.Create(payload)
	(*s).size += 1
	(*s).head = newBlock
}

func (s *Stack) Pop() string {

	payload := s.blockFactory.FreeUp(s.head)
	MyStack.size -= 1
	return payload
}

func (s Stack) PrintSize() {
	fmt.Println("Stack size is: ", s.size)
}
