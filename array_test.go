package pgo_test

import (
	"pgo"
	"reflect"
	"testing"
)

var testInArray = []struct {
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
	for _, object := range testInArray {
		if pgo.InArray(object.val, object.slice) != object.result {
			t.Fatalf("Want: %v, got: %v", object.result, pgo.InArray(object.val, object.slice))
		}
	}
}

var testArrayChunk = []struct {
	array  interface{}
	size   int
	result interface{}
}{
	{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 2, [][]int{[]int{1, 2}, []int{3, 4}, []int{5, 6}, []int{7, 8}}},
	{[]string{"foo", "bar", "baz", "fizz", "buzz"}, 3, [][]string{[]string{"foo", "bar", "baz"}, []string{"fizz", "buzz"}}},
}

func TestArrayChunk(t *testing.T) {
	for _, object := range testArrayChunk {
		res := pgo.ArrayChunk(object.array, object.size)

		s := reflect.ValueOf(object.result)
		len := s.Len()
		for i := 0; i < len; i++ {
			array := s.Index(i).Interface()

			ss := reflect.ValueOf(array)
			arrLen := ss.Len()

			result := reflect.ValueOf(res[i])
			for j := 0; j < arrLen; j++ {
				if result.Index(j).Interface() != ss.Index(j).Interface() {
					t.Fatalf("Want: %v, got: %v", result.Index(j).Interface(), ss.Index(j).Interface())
				}
			}
		}
	}
}
