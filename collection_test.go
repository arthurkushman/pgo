package pgo_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"pgo"
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	// Some items and their priorities.
	items := map[string]int{
		"banana": 3, "apple": 2, "pear": 4, "peach": 1, "plum": 6,
	}

	// Create a Priority queue, put the items in it, and
	// establish the Priority queue (heap) invariants.
	pq := make(pgo.PriorityQueue, len(items))
	i := 0
	for value, priority := range items {
		pq[i] = &pgo.Item{
			Value:    value,
			Priority: priority,
			Index:    i,
		}
		i++
	}
	pq.Init()

	// Insert a new item and then modify its Priority.
	item := &pgo.Item{
		Value:    "orange",
		Priority: 1,
	}
	pq.Push(item)
	pq.Update(item, item.Value, 5)

	assert.Greater(t, pq.Len(), 0)
	// Take the items out; they arrive in decreasing Priority order.
	for pq.Len() > 0 {
		item := pq.Pop().(*pgo.Item)
		fmt.Printf("%.2d:%s ", item.Priority, item.Value)
	}
	// Output:
	// 06:plum 05:orange 04:pear 03:banana 02:apple 01:peach
}
