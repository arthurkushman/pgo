package pgo_test

import (
	"testing"
	"pgo"
)

func TestIsOk(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("The code did not panic from isOk")
		}
	}()

	pgo.Date(123) // inteniously pass int to call isOk()
}