package stack

import "testing"

func TestStackPushAndPop(t *testing.T) {
	stack := EmptyStack[string]()
	stack.Push("hello")
	result, _ := stack.Pop()

	if result != "hello" {
		t.Errorf("Expected hello got %s", result)
	}
}

func TestNewStackIsEmpty(t *testing.T) {
	stack := EmptyStack[string]()

	if !stack.IsEmpty() {
		t.Errorf("New stack should be empty")
	}
}
