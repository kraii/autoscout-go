package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueuePushAndPop(t *testing.T) {
	queue := EmptyQueue[string]()
	queue.Push("hello")
	result, _ := queue.Pop()

	assert.Equal(t, "hello", result)
}

func TestNewqueueIsEmpty(t *testing.T) {
	queue := EmptyQueue[string]()

	assert.Equal(t, true, queue.IsEmpty())
}

func TestQueueIsEmptyAfterPop1Item(t *testing.T) {
	queue := EmptyQueue[string]()
	queue.Push("hello")
	result, _ := queue.Pop()

	assert.Equal(t, "hello", result)
	assert.Equal(t, true, queue.IsEmpty())
}

func TestPopOnce(t *testing.T) {
	queue := EmptyQueue[string]()
	queue.Push("hello")
	queue.Push("goodbye")
	result, _ := queue.Pop()

	assert.Equal(t, "hello", result)
	assert.Equal(t, false, queue.IsEmpty())
}
