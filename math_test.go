package pgo_test

import (
	"github.com/arthurkushman/pgo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRand(t *testing.T) {
	r1 := pgo.Rand(1, 100)
	assert.GreaterOrEqual(t, r1, int64(1))
	assert.LessOrEqual(t, r1, int64(100))
}
