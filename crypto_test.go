package pgo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMd5(t *testing.T) {
	assert.Equal(t, "e99a18c428cb38d5f260853678922e03", Md5("abc123"))
}

func TestSha1(t *testing.T) {
	assert.Equal(t, "6367c48dd193d56ea7b0baad25b19455e529f5ee", Sha1("abc123"))
}
