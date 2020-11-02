package leetcode

import (
	"testing"
)

func TestConstructorMyCircularQueue(t *testing.T) {
	var queue MyCircularQueue

	t.Run("ConstructorMyCircularQueue", func(t *testing.T) {
		queue = ConstructorMyCircularQueue(3)
		queue.EnQueue(1)
		queue.EnQueue(2)
		queue.EnQueue(3)
		t.Log("rear:", queue.Rear())
		result4 := queue.EnQueue(4)
		t.Log("result4:", result4)
		// assert.True(t, result4, true)
		// assert.Equal(t, queue.Rear(), 3)
		// assert.True(t, queue.IsFull(), true)
		// assert.True(t, queue.DeQueue(), true)
		// assert.True(t, queue.EnQueue(4), true)
		// assert.Equal(t, queue.Rear(), 4)
	})
}
