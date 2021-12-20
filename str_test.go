package pgo_test

import (
	"testing"

	"github.com/arthurkushman/pgo"
	"github.com/stretchr/testify/assert"
)

func TestStrReplace(t *testing.T) {
	subject := "The quick brown fox jumped over the lazy dog"

	subject, err := pgo.StrReplace("fox", "cat", subject)
	assert.NoError(t, err)
	subject, err = pgo.StrReplace("dog", "fox", subject)
	assert.NoError(t, err)

	result := "The quick brown cat jumped over the lazy fox"
	assert.Equalf(t, subject, result, "want %s, got %s", subject, result)
}

func TestStrIReplace(t *testing.T) {
	subject := "The quick brown fox jumped over the lazy dog"

	subject, err := pgo.StrIReplace("Fox", "cat", subject)
	assert.NoError(t, err)
	subject, err = pgo.StrIReplace([]string{"DOG", "QuiCK"}, []string{"fox", "slow"}, subject)
	assert.NoError(t, err)

	result := "The slow brown cat jumped over the lazy fox"
	assert.Equalf(t, subject, result, "want %s, got %s", subject, result)
}

func TestStrReplaceCount(t *testing.T) {
	subject := "The quick brown fox jumped over the lazy fox or dog"
	str, err := pgo.StrReplace("fox", "cat", subject, 1)
	assert.NoError(t, err)

	result := "The quick brown cat jumped over the lazy fox or dog"
	assert.Equalf(t, str, result, "want %s, got %s", result, str)
}

func TestStrReplaceArray(t *testing.T) {
	subject := "The quick brown fox jumped over the lazy dog"
	str, err := pgo.StrReplace([]string{"fox", "dog"}, []string{"cat", "elephant"}, subject)
	assert.NoError(t, err)

	result := "The quick brown cat jumped over the lazy elephant"
	assert.Equalf(t, str, result, "want %s, got %s", result, str)
}

func TestStrReplaceErrs(t *testing.T) {
	subject := "The quick brown fox jumped over the lazy dog"

	str, err := pgo.StrReplace([]string{"fox", "dog"}, "", subject)
	assert.Error(t, err)
	assert.Equalf(t, str, subject, "want %s, got %s", subject, str)
}

func TestHTTPBuildQuery(t *testing.T) {
	queryStr := pgo.HTTPBuildQuery(map[string]interface{}{
		"foo":      "bar",
		"bar":      "baz",
		"s":        []string{"1", "foo", "2", "bar", "3", "baz"},
		"num":      123,
		"bigNum":   int64(1238873737737737373),
		"amount":   623.937,
		"isActive": true,
	})

	want := "amount=623.937&bar=baz&bigNum=1238873737737737373&foo=bar&isActive=true&num=123&s=1&s=foo&s=2&s=bar&s=3&s=baz"
	assert.Equal(t, queryStr, want, "want %s, got %s", queryStr, want)

	queryStr2 := pgo.HTTPBuildQuery(map[string]interface{}{})
	assert.Empty(t, queryStr2, "built str from an empty map must be empty")
}
