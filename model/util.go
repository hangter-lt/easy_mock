package model

import (
	"container/list"
	"fmt"
)

type RealTimeChan struct {
	C chan string
}

type CircularQueue struct {
	queue  *list.List
	maxLen int
}

func NewCircularQueue(maxLen int) *CircularQueue {
	return &CircularQueue{
		queue:  list.New(),
		maxLen: maxLen,
	}
}

func (c *CircularQueue) Push(data string) {
	if c.queue.Len() >= c.maxLen {
		c.queue.Remove(c.queue.Front())
	}
	c.queue.PushBack(data)
}

func (c *CircularQueue) GetSlice() []string {
	result := make([]string, 0, c.queue.Len())
	for e := c.queue.Front(); e != nil; e = e.Next() {
		if str, ok := e.Value.(string); ok {
			result = append(result, str)
		}
	}
	fmt.Printf("result: %v\n", result)
	return result
}
