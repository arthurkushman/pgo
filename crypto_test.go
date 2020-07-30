package pgo

import (
	"crypto/md5"
	"crypto/sha256"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMd5(t *testing.T) {
	assert.Equal(t, "e99a18c428cb38d5f260853678922e03", Md5("abc123"))
}

func TestSha1(t *testing.T) {
	assert.Equal(t, "6367c48dd193d56ea7b0baad25b19455e529f5ee", Sha1("abc123"))
}

func TestSha2(t *testing.T) {
	assert.Equal(t, "6ca13d52ca70c883e0f0bb101e425a89e8624de51db2d2392593af6a84118090", Sha2("abc123"))
}

const hashFileName = "example_hash_file.txt"

func TestHashFile(t *testing.T) {
	_, err := FilePutContents(hashFileName, "foo bar baz")
	assert.NoError(t, err)
	hex, err := HashFile("sha1", hashFileName)
	assert.NoError(t, err)
	assert.Equal(t, "c7567e8b39e2428e38bf9c9226ac68de4c67dc39", hex)
	err = os.Remove(hashFileName)
	assert.NoError(t, err)
}

func TestHashHmac(t *testing.T) {
	hmac := HashHmac("foo bar baz", "secret", sha256.New)
	assert.True(t, IsValidMac("foo bar baz", hmac, "secret", sha256.New))

	hmacMd5 := HashHmac("foo bar baz", "secret", md5.New)
	assert.True(t, IsValidMac("foo bar baz", hmacMd5, "secret", md5.New))
}
