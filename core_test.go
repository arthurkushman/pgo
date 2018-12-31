package pgo_test

import (
	"testing"
	"pgo"
)

func TestIsOk(t *testing.T) {
	pgo.Date(123) // check passing wrong typed parameter
}