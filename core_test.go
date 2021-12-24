package pgo_test

import (
	"testing"

	"github.com/arthurkushman/pgo"
	"github.com/stretchr/testify/assert"
)

func TestIsOk(t *testing.T) {
	pgo.Date(123) // check passing wrong typed parameter
}

func TestSerialize(t *testing.T) {
	m := make(map[int]string)
	m[0] = "abc"

	str, err := pgo.Serialize(m)
	assert.NoError(t, err)
	assert.Equal(t, str, "Dv+BBAEC/4IAAQQBDAAACf+CAAEAA2FiYw==")

	unserMap := make(map[int]string)
	err = pgo.Unserialize(str, &unserMap)
	assert.NoError(t, err)
	assert.Equal(t, m, map[int]string{
		0: "abc",
	})
}
