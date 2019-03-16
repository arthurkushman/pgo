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
