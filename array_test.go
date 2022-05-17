package pgo_test

import (
	"math"
	"reflect"
	"strings"
	"testing"

	"github.com/arthurkushman/pgo"
	"github.com/stretchr/testify/assert"
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
		switch object.slice.(type) {
		case []int:
			isInArray := pgo.InArray(object.val.(int), object.slice.([]int))
			assert.Equalf(t, isInArray, object.result, "Want: %v, got: %v", object.result, isInArray)
		case []float64:
			isInArray := pgo.InArray(object.val.(float64), object.slice.([]float64))
			assert.Equalf(t, isInArray, object.result, "Want: %v, got: %v", object.result, isInArray)
		case []string:
			isInArray := pgo.InArray(object.val.(string), object.slice.([]string))
			assert.Equalf(t, isInArray, object.result, "Want: %v, got: %v", object.result, isInArray)
		}
	}
}

func TestArrayChunkInts(t *testing.T) {
	res := pgo.ArrayChunk([]int{1, 2, 3, 4, 5, 6, 7, 8}, 2)

	s := [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}
	l := len(s)
	for i := 0; i < l; i++ {
		array := s[i]
		arrLen := len(array)

		result := res[i]
		for j := 0; j < arrLen; j++ {
			assert.Equalf(t, result[j], array[j], "Want: %v, got: %v", result, array[j])
		}
	}
}

func TestArrayChunkStrings(t *testing.T) {
	res := pgo.ArrayChunk([]string{"foo", "bar", "baz", "fizz", "buzz"}, 3)

	s := [][]string{{"foo", "bar", "baz"}, {"fizz", "buzz"}}
	l := len(s)
	for i := 0; i < l; i++ {
		array := s[i]
		arrLen := len(array)

		result := res[i]
		for j := 0; j < arrLen; j++ {
			assert.Equalf(t, result[j], array[j], "Want: %v, got: %v", result, array[j])
		}
	}
}

func TestArrayCombine(t *testing.T) {
	// ints + strings
	res := pgo.ArrayCombine([]int{11, 32, 13, 14, 51, 46, 17, 88}, []string{"foo", "bar", "baz", "fizz", "buzz", "mazz", "freez", "lorum"})
	m := map[int]string{
		11: "foo",
		32: "bar",
		13: "baz",
		14: "fizz",
		51: "buzz",
		46: "mazz",
		17: "freez",
		88: "lorum",
	}
	for k, v := range res {
		assert.Equalf(t, m[k], v, "want %d, got %d", m[k], v)
	}

	// string + float
	resSf := pgo.ArrayCombine([]string{"foo", "bar", "baz", "fizz", "buzz"}, []float64{11.32, 32.42, 13.246, 14.41, 51.98})
	mSf := map[string]float64{
		"foo":  11.32,
		"bar":  32.42,
		"baz":  13.246,
		"fizz": 14.41,
		"buzz": 51.98,
	}
	for k, v := range resSf {
		assert.Equalf(t, mSf[k], v, "want %d, got %d", mSf[k], v)
	}

	// non-equal = nil
	resSf = pgo.ArrayCombine([]string{"foo", "bar", "baz", "buzz"}, []float64{11.32, 32.42, 13.246, 14.41, 51.98})
	assert.Nil(t, resSf)
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
		switch object.values.(type) {
		case []int:
			res := pgo.ArrayCountValues(object.values.([]int))
			m := reflect.ValueOf(object.result)
			for k, v := range res {
				assert.Equalf(t, m.MapIndex(reflect.ValueOf(k)).Interface(), v, "want %d, got %d", m.MapIndex(reflect.ValueOf(k)).Interface(), v)
			}
		case []float64:
			res := pgo.ArrayCountValues(object.values.([]float64))
			m := reflect.ValueOf(object.result)
			for k, v := range res {
				assert.Equalf(t, m.MapIndex(reflect.ValueOf(k)).Interface(), v, "want %d, got %d", m.MapIndex(reflect.ValueOf(k)).Interface(), v)
			}
		case []string:
			res := pgo.ArrayCountValues(object.values.([]string))
			m := reflect.ValueOf(object.result)
			for k, v := range res {
				assert.Equalf(t, m.MapIndex(reflect.ValueOf(k)).Interface(), v, "want %d, got %d", m.MapIndex(reflect.ValueOf(k)).Interface(), v)
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
			assert.Equalf(t, v, object.result[k], "want %v, got %v", v, object.values[k])
		}
	}

	for _, object := range testArrayMapFloats {
		res := pgo.ArrayMap(object.values, object.callback)
		for k, v := range res {
			assert.Equalf(t, v, object.result[k], "want %v, got %v", v, object.values[k])
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
			assert.Equalf(t, v, object.result[k], "want %v, got %v", v, object.values[k])
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

var testArrayUdiff = []struct {
	values interface{}
	diff   interface{}
	result interface{}
	uf     func(a interface{}, b interface{}) int
}{
	{[]string{"foo", "bar", "fizz", "baz"}, []string{"foo", "bar"}, []string{"fizz", "baz"}, func(a interface{}, b interface{}) int {
		if a.(string) > b.(string) {
			return 1
		} else if a.(string) < b.(string) {
			return -1
		}

		return 0
	}},
	{[]int{3, 43, 8, 4, 9}, []int{3, 8, 9, 4}, []int{43}, func(a interface{}, b interface{}) int {
		if a.(int) > b.(int) {
			return 1
		} else if a.(int) < b.(int) {
			return -1
		}

		return 0
	}},
	{[]float64{3.14159, 43.03, 3.14159, 43.02, 8.74}, []float64{3.14159, 43.03, 3.14159}, []float64{43.02, 8.74}, func(a interface{}, b interface{}) int {
		if a.(float64) > b.(float64) {
			return 1
		} else if a.(float64) < b.(float64) {
			return -1
		}

		return 0
	}},
	{[]int{3, 43, 8, 4, 9}, []int{}, []int{3, 43, 8, 4, 9}, func(a interface{}, b interface{}) int {
		if a.(int) > b.(int) {
			return 1
		} else if a.(int) < b.(int) {
			return -1
		}

		return 0
	}},
	{[]int{}, []int{3, 43, 8, 4, 9}, []int{}, func(a interface{}, b interface{}) int {
		if a.(int) > b.(int) {
			return 1
		} else if a.(int) < b.(int) {
			return -1
		}

		return 0
	}},
}

func TestArrayDiff(t *testing.T) {
	for _, object := range testArrayDiff {
		res := pgo.ArrayDiff(object.values, object.diff)

		s := reflect.ValueOf(object.result)
		len := s.Len()
		for i := 0; i < len; i++ {
			assert.Equalf(t, s.Index(i).Interface(), res[i], "want %v, got %v", s.Index(i).Interface(), res[i])
		}
	}
}

func TestArrayUDiff(t *testing.T) {
	for _, object := range testArrayUdiff {
		ov := object.values
		od := object.diff
		uf := object.uf
		res := pgo.ArrayUdiff(uf, ov, od)

		s := reflect.ValueOf(object.result)
		len := s.Len()
		for i := 0; i < len; i++ {
			assert.Equalf(t, s.Index(i).Interface(), res[i], "want %v, got %v", s.Index(i).Interface(), res[i])
		}
	}
}

var testArraySum = []struct {
	values interface{}
	result float64
}{
	{[]int{3, 43, 8, 43, 8}, 105},
	{[]float64{3.14159, 43.03, 8, 3.14159, 43.02, 8}, 108.33318},
}

func TestArraySum(t *testing.T) {
	for _, object := range testArraySum {
		switch object.values.(type) {
		case []int:
			res, _ := pgo.ArraySum(object.values.([]int))
			assert.Equalf(t, res, int(object.result), "want %v, got %v", object.result, res)
		case []float64:
			res, _ := pgo.ArraySum(object.values.([]float64))
			assert.Equalf(t, res, object.result, "want %v, got %v", object.result, res)
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
			assert.Equalf(t, resVal.Index(i).Interface(), s.Index(i).Interface(), "want %v, got %v", s.Index(i).Interface(), res[i])
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

var testRangeWithStep = []struct {
	min    int
	max    int
	step   int
	result []int
}{
	{3, 9, 2, []int{3, 5, 7, 9}},
	{-3, 7, 5, []int{-3, 2, 7}},
}

func TestRange(t *testing.T) {
	for _, object := range testRange {
		res := pgo.Range(object.min, object.max)
		for k, v := range res {
			assert.Equalf(t, v, object.result[k], "want %v, got %v", object.result[k], v)
		}
	}

	for _, object := range testRangeWithStep {
		res := pgo.Range(object.min, object.max, object.step)
		for k, v := range res {
			assert.Equalf(t, v, object.result[k], "want %v, got %v", object.result[k], v)
		}
	}
}

var testEqual = []struct {
	a   interface{}
	b   interface{}
	res bool
}{
	{[]int{1, 2, 3}, []int{1, 2, 3}, true},
	{[]int{1, 2, 3}, []int{1, 3, 2}, false},
	{[]int{1, 2, 3}, []int{}, false},
	{[]int{}, []int{}, true},
	{[]string{"foo"}, []string{"bar"}, false},
	{[]string{"foo"}, []string{"foo"}, true},
	{[]float64{123.33, 22}, []float64{123.33, 22}, true},
	{[]float64{123.33, 22}, []float64{123.33, 22.1111}, false},
	{[]bool{true, false}, []bool{true, false}, true},
	{[]bool{true, false}, []bool{false, false}, false},
	{[]byte{0, 123, 1}, []byte{0, 123, 1}, true},
	{[]byte{0, 123, 1}, []byte{0, 123, 133}, false},
}

func TestEqual(t *testing.T) {
	for _, v := range testEqual {
		switch v.a.(type) {
		case []int:
			res := pgo.EqualSlices(v.a.([]int), v.b.([]int))
			assert.Equal(t, v.res, res)
		case []bool:
			res := pgo.EqualSlices(v.a.([]bool), v.b.([]bool))
			assert.Equal(t, v.res, res)
		case []float64:
			res := pgo.EqualSlices(v.a.([]float64), v.b.([]float64))
			assert.Equal(t, v.res, res)
		case []string:
			res := pgo.EqualSlices(v.a.([]string), v.b.([]string))
			assert.Equal(t, v.res, res)
		case []byte:
			res := pgo.EqualSlices(v.a.([]byte), v.b.([]byte))
			assert.Equal(t, v.res, res)
		}
	}
}

func TestArrayMax(t *testing.T) {
	res := pgo.ArrayMax([]int{})
	assert.Equal(t, 0, res)

	res = pgo.ArrayMax([]int{3, 1, 2, 9})
	assert.Equal(t, 9, res)

	res = pgo.ArrayMax([]int{-3, -1, -2, -9})
	assert.Equal(t, -1, res)

	res = pgo.ArrayMax([]int{-3, -1, -2, -9, 3, 1, 2, 9})
	assert.Equal(t, 9, res)

	resF := pgo.ArrayMax([]float64{})
	assert.Equal(t, float64(0), resF)

	resF = pgo.ArrayMax([]float64{3.2, 1.0837, 2.123, 9.87})
	assert.Equal(t, 9.87, resF)

	resF = pgo.ArrayMax([]float64{-3.12, -1.678, -2.01, -9.007})
	assert.Equal(t, -1.678, resF)

	resF = pgo.ArrayMax([]float64{3.2, 1.0837, 2.123, 9.87, -3.12, -1.678, -2.01, -9.007})
	assert.Equal(t, 9.87, resF)
}

func TestArrayMin(t *testing.T) {
	res := pgo.ArrayMin([]int{})
	assert.Equal(t, 0, res)

	res = pgo.ArrayMin([]int{3, 1, 2, 9})
	assert.Equal(t, 1, res)

	res = pgo.ArrayMin([]int{-3, -1, -2, -9})
	assert.Equal(t, -9, res)

	res = pgo.ArrayMin([]int{-3, -1, -2, -9, 3, 1, 2, 9})
	assert.Equal(t, -9, res)

	resF := pgo.ArrayMin([]float64{})
	assert.Equal(t, float64(0), resF)

	resF = pgo.ArrayMin([]float64{3.2, 1.0837, 2.123, 9.87})
	assert.Equal(t, 1.0837, resF)

	resF = pgo.ArrayMin([]float64{-3.12, -1.678, -2.01, -9.007})
	assert.Equal(t, -9.007, resF)

	resF = pgo.ArrayMin([]float64{-3.12, -1.678, -2.01, -9.007, 3.2, 1.0837, 2.123, 9.87})
	assert.Equal(t, -9.007, resF)
}
