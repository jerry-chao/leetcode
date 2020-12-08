package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinStack_GetMin(t *testing.T) {

	var stack MinStack

	t.Run("Constructor", func(t *testing.T) {
		stack = Constructor()
	})

	t.Run("push", func(t *testing.T) {
		stack.Push(-2)
	})
	t.Run("push", func(t *testing.T) {
		stack.Push(0)
	})
	t.Run("push", func(t *testing.T) {
		stack.Push(-3)
	})
	t.Run("getMin", func(t *testing.T) {
		assert.Equal(t, stack.GetMin(), -3)
	})
	t.Run("pop", func(t *testing.T) {
		stack.Pop()
	})
	t.Run("top", func(t *testing.T) {
		assert.Equal(t, stack.Top(), 0)
	})
	t.Run("getMin", func(t *testing.T) {
		assert.Equal(t, stack.GetMin(), -2)
	})
}
