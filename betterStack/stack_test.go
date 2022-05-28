package betterStack

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockedBlock struct {
	mock.Mock
}

func (mBt *MockedBlock) Create(payload string) BlockType {
	mBt.Called(payload)
	return &MockedBlock{}
}

func (mBt *MockedBlock) FreeUp(blockType BlockType) string {
	return ""

}

func TestStack_Push(t *testing.T) {
	mockedBlock := MockedBlock{}
	mockedBlock.On("Create", "Test").Return(&MockedBlock{})

	s := Stack{blockFactory: &mockedBlock}

	s.Push("Test")
	s.Push("Test")
	mockedBlock.AssertNumberOfCalls(t, "Create", 2)

	assert.Equal(t, 2, s.size, "length should be equal")
}
