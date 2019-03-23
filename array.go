package pgo

import (
	"reflect"
)

// InArray checks if a value exists in an array
func InArray(needle interface{}, haystack interface{}) bool {
	return search(needle, haystack)
}

func search(needle interface{}, haystack interface{}) bool {
	switch reflect.TypeOf(haystack).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(haystack)
		len := s.Len()

		for i := 0; i < len; i++ {
			if needle == s.Index(i).Interface() {
				return true
			}
		}
	}

	return false
}

// ArrayChunk split an array into chunks
func ArrayChunk(array interface{}, size int) []interface{} {
	var chunks []interface{}

	s := reflect.ValueOf(array)
	len := s.Len()

	var subChunk []interface{}
	for i := 0; i < len; i++ {
		subChunk = append(subChunk, s.Index(i).Interface())

		if (i+1)%size == 0 || i+1 == len {
			chunks = append(chunks, subChunk)
			subChunk = make([]interface{}, 0)
		}
	}

	return chunks
}

// ArrayCombine creates an array by using one array for keys and another for its values
// returns map[key]value if both slices are equal and nil otherwise
func ArrayCombine(keys interface{}, values interface{}) map[interface{}]interface{} {
	s := reflect.ValueOf(keys)
	len := s.Len()

	ss := reflect.ValueOf(values)
	ssLen := ss.Len()

	if len != ssLen {
		return nil
	}

	resultMap := make(map[interface{}]interface{})
	for i := 0; i < len; i++ {
		resultMap[s.Index(i).Interface()] = ss.Index(i).Interface()
	}

	return resultMap
}

// ArrayCountValues counts all the values of an array/slice
func ArrayCountValues(array interface{}) map[interface{}]int {
	res := make(map[interface{}]int)

	s := reflect.ValueOf(array)
	len := s.Len()
	for i := 0; i < len; i++ {
		res[s.Index(i).Interface()]++
	}

	return res
}

// ArrayMap applies the callback to the elements of the given arrays
func ArrayMap(array interface{}, callback interface{}) []interface{} {
	s := reflect.ValueOf(array)
	len := s.Len()

	funcValue := reflect.ValueOf(callback)

	var result []interface{}
	for i := 0; i < len; i++ {
		result = append(result, funcValue.Call([]reflect.Value{s.Index(i)})[0].Interface())
	}

	return result
}

// ArrayFilter filters elements of an array using a callback function
func ArrayFilter(array interface{}, callback interface{}) []interface{} {
	s := reflect.ValueOf(array)
	len := s.Len()

	funcValue := reflect.ValueOf(callback)

	var result []interface{}
	for i := 0; i < len; i++ {
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
	len := s.Len()

	var result []interface{}
	isFound := false

	for i := 0; i < len; i++ {
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

		if isFound == false {
			result = append(result, needle)
		}

		isFound = false
	}

	return result
}

// ArrayKeys return all the keys or a subset of the keys of an array
func ArrayKeys(array map[interface{}]interface{}) []interface{} {
	m := reflect.ValueOf(array)

	var result []interface{}
	for k := range array {
		result = append(result, m.MapIndex(reflect.ValueOf(k)).Interface())
	}

	return result
}
