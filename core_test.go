package pgo_test

import (
	"pgo"
	"testing"
)

func TestIsOk(t *testing.T) {
	pgo.Date(123) // check passing wrong typed parameter
}
