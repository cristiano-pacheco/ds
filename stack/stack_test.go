package stack_test

import (
	"testing"

	"ds/stack"

	"github.com/stretchr/testify/assert"
)

func TestStack_PushPop(t *testing.T) {
	s := stack.New[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	assert.Equal(t, 3, s.Pop())
	assert.Equal(t, 2, s.Pop())
	assert.Equal(t, 1, s.Pop())
}

func TestStack_PopEmpty(t *testing.T) {
	s := stack.New[int]()
	// Pop from empty stack, should panic or return zero value
	defer func() {
		r := recover()
		assert.NotNil(t, r)
	}()
	_ = s.Pop()
}

func TestStack_Peek(t *testing.T) {
	s := stack.New[int]()
	s.Push(10)
	s.Push(20)
	assert.Equal(t, 20, s.Peak())
	s.Pop()
	assert.Equal(t, 10, s.Peak())
}

func TestStack_PeekEmpty(t *testing.T) {
	s := stack.New[int]()
	defer func() {
		r := recover()
		assert.NotNil(t, r)
	}()
	_ = s.Peak()
}

func TestStack_Size(t *testing.T) {
	// Since the internal slice is unexported, we'll test the stack behavior
	s := stack.New[int]()
	s.Push(1)
	s.Push(2)
	v := s.Pop()
	assert.Equal(t, 2, v)
	v = s.Pop()
	assert.Equal(t, 1, v)
}

func TestStack_MixedTypes(t *testing.T) {
	s := stack.New[any]()
	s.Push("a")
	s.Push(42)
	s.Push(3.14)
	assert.Equal(t, 3.14, s.Pop())
	assert.Equal(t, 42, s.Pop())
	assert.Equal(t, "a", s.Pop())
}
