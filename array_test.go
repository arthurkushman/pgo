package pgo_test

import (
	"pgo"
	"testing"
)

var testInts = []struct {
	val    interface{}
	slice  interface{}
	result bool
}{
	{3, []int{1, 2, 3}, true},
	{7, []int{1, 2, 3}, false},
	{"bar", []string{"foo", "bar", "baz"}, true},
	{"bar33", []string{"foo", "bar", "baz"}, false},
	{3.14159, []float64{33.12, 12.333, 3.14159, 78.4429}, true},
	{3.141594, []float64{33.12, 12.333, 3.14159, 78.4429}, false},
}

func TestInArray(t *testing.T) {
	for _, object := range testInts {
		if pgo.InArray(object.val, object.slice) != object.result {
			t.Fatalf("Want: %v, got: %v", object.result, pgo.InArray(object.val, object.slice))
		}
	}
}
