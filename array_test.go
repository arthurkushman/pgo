package pgo_test

import (
	"math"
	"pgo"
	"reflect"
	"strings"
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

// type emptyMap map[interface{}]interface{}

var testArrayCombine = []struct {
	keys   interface{}
	values interface{}
	result interface{}
}{
	{[]int{11, 32, 13, 14, 51, 46, 17, 88}, []string{"foo", "bar", "baz", "fizz", "buzz", "mazz", "freez", "lorum"}, map[int]string{
		11: "foo",
		32: "bar",
		13: "baz",
		14: "fizz",
		51: "buzz",
		46: "mazz",
		17: "freez",
		88: "lorum",
	}},
	{[]string{"foo", "bar", "baz", "fizz", "buzz"}, []float64{11.32, 32.42, 13.246, 14.41, 51.98},
		map[string]float64{
			"foo":  11.32,
			"bar":  32.42,
			"baz":  13.246,
			"fizz": 14.41,
			"buzz": 51.98,
		}},
	{[]string{"foo", "bar", "baz", "buzz"}, []float64{11.32, 32.42, 13.246, 14.41, 51.98}, nil},
}

func TestArrayCombine(t *testing.T) {
	for _, object := range testArrayCombine {
		res := pgo.ArrayCombine(object.keys, object.values)

		m := reflect.ValueOf(object.result)
		for k, v := range res {
			if m.MapIndex(reflect.ValueOf(k)).Interface() != v {
				t.Fatalf("want %d, got %d", m.MapIndex(reflect.ValueOf(k)).Interface(), v)
			}
		}
	}
}

var testArrayCountValues = []struct {
	values interface{}
	result interface{}
}{
	{[]string{"foo", "bar", "foo", "baz", "bar", "bar"}, map[string]int{"foo": 2, "bar": 3, "baz": 1}},
	{[]int{3, 43, 8, 43, 8}, map[int]int{43: 2, 8: 2, 3: 1}},
	{[]float64{3.14159, 43.03, 8, 3.14159, 43.02, 8}, map[float64]int{3.14159: 2, 8: 2, 43.03: 1, 43.02: 1}},
}

func TestArrayCountValues(t *testing.T) {
	for _, object := range testArrayCountValues {
		res := pgo.ArrayCountValues(object.values)

		m := reflect.ValueOf(object.result)
		for k, v := range res {
			if m.MapIndex(reflect.ValueOf(k)).Interface() != v {
				t.Fatalf("want %d, got %d", m.MapIndex(reflect.ValueOf(k)).Interface(), v)
			}
		}
	}
}

var testArrayMapStrings = []struct {
	values   []string
	callback func(v string) string
	result   []string
}{
	{[]string{"foo", "bar", "baz"}, func(v string) string {
		return strings.ToUpper(v)
	}, []string{"FOO", "BAR", "BAZ"}},
	{[]string{"FOO", "BAR", "BAZ"}, func(v string) string {
		return strings.ToLower(v)
	}, []string{"foo", "bar", "baz"}},
}

var testArrayMapFloats = []struct {
	values   []float64
	callback func(v float64) float64
	result   []float64
}{
	{[]float64{1, 2, 3, 4, 5}, func(v float64) float64 {
		return math.Pow(v, 2)
	}, []float64{1, 4, 9, 16, 25}},
	{[]float64{1, 2, 3, 4, 5}, func(v float64) float64 {
		return math.Pow(v, 3)
	}, []float64{1, 8, 27, 64, 125}},
}

func TestArrayMap(t *testing.T) {
	for _, object := range testArrayMapStrings {
		res := pgo.ArrayMap(object.values, object.callback)

		for k, v := range res {
			if v != object.result[k] {
				t.Fatalf("want %v, got %v", v, object.values[k])
			}
		}
	}

	for _, object := range testArrayMapStrings {
		res := pgo.ArrayMap(object.values, object.callback)

		for k, v := range res {
			if v != object.result[k] {
				t.Fatalf("want %v, got %v", v, object.values[k])
			}
		}
	}
}

var testArrayFilterFloats = []struct {
	values   []float64
	callback func(v float64) bool
	result   []float64
}{
	{[]float64{1, 2, 3, 4, 5}, func(v float64) bool {
		return v > 2.718
	}, []float64{3, 4, 5}},
	{[]float64{1, 2, 3, 4, 5}, func(v float64) bool {
		return v < 2.718
	}, []float64{1, 2}},
}

func TestArrayFilter(t *testing.T) {
	for _, object := range testArrayFilterFloats {
		res := pgo.ArrayFilter(object.values, object.callback)

		for k, v := range res {
			if v != object.result[k] {
				t.Fatalf("want %v, got %v", v, object.values[k])
			}
		}
	}
}

var testArrayDiff = []struct {
	values interface{}
	diff   interface{}
	result interface{}
}{
	{[]string{"foo", "bar", "fizz", "baz"}, []string{"foo", "bar"}, []string{"fizz", "baz"}},
	{[]int{3, 43, 8, 4, 9}, []int{3, 8, 9, 4}, []int{43}},
	{[]float64{3.14159, 43.03, 3.14159, 43.02, 8.74}, []float64{3.14159, 43.03, 3.14159}, []float64{43.02, 8.74}},
	{[]int{3, 43, 8, 4, 9}, []int{}, []int{3, 43, 8, 4, 9}},
	{[]int{}, []int{3, 43, 8, 4, 9}, []int{}},
}

func TestArrayDiff(t *testing.T) {
	for _, object := range testArrayDiff {
		res := pgo.ArrayDiff(object.values, object.diff)

		s := reflect.ValueOf(object.result)
		len := s.Len()
		for i := 0; i < len; i++ {
			if s.Index(i).Interface() != res[i] {
				t.Fatalf("want %v, got %v", s.Index(i).Interface(), res[i])
			}
		}
	}
}

var testArrayKeys = []struct {
	values map[string]int
	result []string
}{
	{map[string]int{"foo": 1, "bar": 8, "fizz": 12, "baz": 0}, []string{"foo", "bar", "fizz", "baz"}},
	// {map[int]interface{}{3: 32.4, 43: "foo", 8: "bar", 4: 1, 9: 1}, []int{3, 43, 8, 4, 9}},
	// {map[interface{}]int{3.45: 32, "foo": 33, 8: 53, "bar": 1, 9: 1}, []interface{}{3.45, "foo", 8, "bar", 9}},
}

var testArraySum = []struct {
	values interface{}
	result float64
}{
	{[]int{3, 43, 8, 43, 8}, 105},
	{[]interface{}{3, "foo", 8, 43, 8}, 0},
	{[]float64{3.14159, 43.03, 8, 3.14159, 43.02, 8}, 108.33318},
}

func TestArraySum(t *testing.T) {
	for _, object := range testArraySum {
		res, _ := pgo.ArraySum(object.values)

		if res != object.result {
			t.Fatalf("want %v, got %v", object.result, res)
		}
	}
}

var testArrayIntersect = []struct {
	values interface{}
	diff   interface{}
	diff2  interface{}
	result interface{}
}{
	{[]int{12, 54, 32, 12, 33}, []int{3, 12, 54, 9}, []int{12, 33, 9}, []int{12, 54, 33}},
	{[]string{"foo", "bar", "baz", "fizz", "bazz", "fizz", "fizz"}, []string{"bar", "fizz"},
		[]string{"foo", "bar", "hey"}, []string{"foo", "bar", "fizz"}},
}

func TestArrayIntersect(t *testing.T) {
	for _, object := range testArrayIntersect {
		res := pgo.ArrayIntersect(object.values, object.diff, object.diff2)

		resVal := reflect.ValueOf(res)
		resLen := resVal.Len()

		s := reflect.ValueOf(object.result)

		for i := 0; i < resLen; i++ {
			if resVal.Index(i).Interface() != s.Index(i).Interface() {
				t.Fatalf("want %v, got %v", s.Index(i).Interface(), res[i])
			}
		}
	}
}

var testRange = []struct {
	min    int
	max    int
	result []int
}{
	{3, 9, []int{3, 4, 5, 6, 7, 8, 9}},
	{-3, 7, []int{-3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7}},
}

func TestRange(t *testing.T) {
	for _, object := range testRange {

		res := pgo.Range(object.min, object.max)
		for k, v := range res {
			if v != object.result[k] {
				t.Fatalf("want %v, got %v", object.result[k], v)
			}
		}
	}
}
