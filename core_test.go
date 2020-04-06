package pgo_test

import (
	"github.com/arthurkushman/pgo"
	"testing"
)

func TestIsOk(t *testing.T) {
	pgo.Date(123) // check passing wrong typed parameter
}
