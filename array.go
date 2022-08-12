package pgo

import (
	"reflect"
	"sort"

	"github.com/arthurkushman/pgo/constraints"
)

// InArray checks if a value exists in an array
func InArray[T comparable](needle T, haystack []T) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}

	return false
}

// ArrayChunk split an array into chunks
func ArrayChunk[T comparable](array []T, size int) [][]T {
	var chunks [][]T

	l := len(array)
	var subChunk []T
	for i := 0; i < l; i++ {
		subChunk = append(subChunk, array[i])

		if (i+1)%size == 0 || i+1 == l {
			chunks = append(chunks, subChunk)
			subChunk = make([]T, 0)
		}
	}

	return chunks
}

// ArrayCombine creates an array by using one array for keys and another for its values
// returns map[key]value if both slices are equal and nil otherwise
func ArrayCombine[K, V comparable](keys []K, values []V) map[K]V {
	kLen := len(keys)
	vLen := len(values)
	if kLen != vLen {
		return nil
	}

	resultMap := make(map[K]V)
	for i := 0; i < kLen; i++ {
		resultMap[keys[i]] = values[i]
	}

	return resultMap
}

// ArrayCountValues counts all the values of an array/slice
func ArrayCountValues[T comparable](array []T) map[T]int {
	res := make(map[T]int)
	l := len(array)
	for i := 0; i < l; i++ {
		res[array[i]]++
	}

	return res
}

// ArrayMap applies the callback to the elements of the given arrays
func ArrayMap(array interface{}, callback interface{}) []interface{} {
	s := reflect.ValueOf(array)
	l := s.Len()

	funcValue := reflect.ValueOf(callback)

	var result []interface{}
	for i := 0; i < l; i++ {
		result = append(result, funcValue.Call([]reflect.Value{s.Index(i)})[0].Interface())
	}

	return result
}

// ArrayFilter filters elements of an array using a callback function
func ArrayFilter(array interface{}, callback interface{}) []interface{} {
	s := reflect.ValueOf(array)
	l := s.Len()

	funcValue := reflect.ValueOf(callback)

	var result []interface{}
	for i := 0; i < l; i++ {
		if funcValue.Call([]reflect.Value{s.Index(i)})[0].Bool() {
			result = append(result, s.Index(i).Interface())
		}
	}

	return result
}

// ArrayDiff compares array1 against one or more other arrays
// returns the values in array1 that are not present in any of the other arrays
func ArrayDiff(arrays ...interface{}) []interface{} {
	s := reflect.ValueOf(arrays[0])
	l := s.Len()

	var result []interface{}
	isFound := false

	for i := 0; i < l; i++ {
		needle := s.Index(i).Interface()

		for _, v := range arrays[1:] {
			switch reflect.TypeOf(v).Kind() {
			case reflect.Slice:
				ss := reflect.ValueOf(v)
				sLen := ss.Len()

				for j := 0; j < sLen; j++ {
					if needle == ss.Index(j).Interface() {
						isFound = true
					}
				}
			}
		}

		if !isFound {
			result = append(result, needle)
		}

		isFound = false
	}

	return result
}

// ArrayUdiff computes the difference of arrays by using a callback function for data comparison
func ArrayUdiff(uf func(interface{}, interface{}) int, arrays ...interface{}) []interface{} {
	var result []interface{}

	first := reflect.ValueOf(arrays[0])
	firstLen := first.Len()
	elementType := reflect.TypeOf(arrays[0]).Elem()

	originFirst := reflect.MakeSlice(reflect.SliceOf(elementType), firstLen, firstLen)
	reflect.Copy(originFirst, first)

	for _, ar := range arrays {
		sort.Slice(ar, func(i, j int) bool {
			first := reflect.ValueOf(ar)
			return 0 > uf(first.Index(i).Interface(), first.Index(j).Interface())
		})
	}

	tempResult := make(map[interface{}]interface{})

	for i := 0; i < firstLen; i++ {
		needle := first.Index(i).Interface()
		isFound := false

		for _, v := range arrays[1:] {
			second := reflect.ValueOf(v)
			secondLen := second.Len()

			for j := 0; j < secondLen; j++ {
				valSecond := second.Index(j).Interface()
				compareResult := uf(needle, valSecond)

				if compareResult < 0 {
					break
				}

				if compareResult == 0 {
					isFound = true
					break
				}
			}

			if isFound {
				break
			}
		}

		if !isFound {
			tempResult[needle] = needle
		}
	}

	originFirstLen := originFirst.Len()

	for i := 0; i < originFirstLen && len(tempResult) > 0; i++ {
		needle := originFirst.Index(i).Interface()

		val, ok := tempResult[needle]
		if ok {
			result = append(result, val)
			delete(tempResult, needle)
		}

	}

	return result
}

// ArraySum calculate the sum of values in an array
func ArraySum[T constraints.IntFloat](array []T) (T, error) {
	l := len(array)

	var amount T
	for i := 0; i < l; i++ {
		amount += array[i]
	}

	return amount, nil
}

// ArrayIntersect computes the intersection of arrays
func ArrayIntersect(arrays ...interface{}) []interface{} {
	s := reflect.ValueOf(arrays[0])
	l := s.Len()

	var result []interface{}
	isFound := false

	intersected := make(map[interface{}]bool)
	for i := 0; i < l; i++ {
		needle := s.Index(i).Interface()

		for _, v := range arrays[1:] {
			switch reflect.TypeOf(v).Kind() {
			case reflect.Slice:
				ss := reflect.ValueOf(v)
				sLen := ss.Len()

				for j := 0; j < sLen; j++ {
					if needle == ss.Index(j).Interface() && !intersected[needle] {
						isFound = true
						intersected[needle] = true // del op is more expensive for slices
						goto out                   // it is stupid to iterate O(n^2) if found
					}
				}
			}
		}

	out:
		if isFound {
			result = append(result, needle)
		}

		isFound = false
	}

	return result
}

// Range creates int slice of min to max range
// If a step value is given, it will be used as the increment between elements in the sequence.
// step should be given as a positive number.
// If not specified, step will default to 1.
func Range(min, max int, step ...int) []int {
	var slice []int

	var argsLen = len(step)

	var stepp = 1
	if argsLen > 0 && step[0] > 1 {
		stepp = step[0]
	}

	for i := min; i <= max; i += stepp {
		slice = append(slice, i)
	}

	return slice
}

// EqualSlices compares two slices and returns true if they are equal, false otherwise
// in case of passing wrong (non-slice) arguments error will be returned
func EqualSlices[T comparable](a, b []T) bool {
	la := len(a)
	lb := len(b)
	if la != lb {
		return false
	}

	for i := 0; i < la; i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

// ArrayMax finds maximum value from []T
func ArrayMax[T constraints.IntFloat](a []T) T {
	var max T
	if len(a) > 0 {
		max = a[0]
	}

	for _, v := range a {
		if max < v {
			max = v
		}
	}

	return max
}

// ArrayMin finds minimum value from []T
func ArrayMin[T constraints.IntFloat](a []T) T {
	var min T
	if len(a) > 0 {
		min = a[0]
	}

	for _, v := range a {
		if min > v {
			min = v
		}
	}

	return min
}

// ArrayUnique returns unique values form []T
func ArrayUnique[T comparable](a []T) []T {
	m := map[T]struct{}{}
	for _, v := range a {
		if _, ok := m[v]; !ok {
			m[v] = struct{}{}
		}
	}

	s := make([]T, len(m))
	i := 0
	for k := range m {
		s[i] = k
		i++
	}

	return s
}

// ArrayValues Return all the values of map as a slice of values with corresponding type
func ArrayValues[K, V comparable](a map[K]V) []V {
	s := make([]V, len(a))
	var i uint64
	for _, v := range a {
		s[i] = v
		i++
	}

	return s
}

// ArrayReverse reverse slice passed as an argument
func ArrayReverse[V comparable](s []V) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
