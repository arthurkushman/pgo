package pgo_test

import (
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/arthurkushman/pgo"
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

func TestConcatFast(t *testing.T) {
	tests := []struct {
		name   string
		s      []string
		result string
	}{
		{
			name:   "concat 3 strings",
			s:      []string{"foo", "bar", "bazzz"},
			result: "foobarbazzz",
		},
		{
			name:   "concat 0 strings",
			s:      []string{},
			result: "",
		},
		{
			name: "concat random strings",
			s: []string{"impEdfCJyek3jn5kj3nkj35nkj35nkj3nkj3n5kjn3kjn35kjn5", "IpDtUOSwMy", "sMIaQYdeON", "TZTwRNgZfx",
				"kybtlfzfJa", "UJQJXhknLe", "GKDmxroeFv",
				"ifguLESWvm333334241341231242414k12m4k1m24k1m2k4m1k24n1l2n41ln41lk2n4k12"},
			result: "impEdfCJyek3jn5kj3nkj35nkj35nkj3nkj3n5kjn3kjn35kjn5IpDtUOSwMysMIaQYdeONTZTwRNgZfxkybtlfzfJaUJQJXhknLeGKDmxroeFvifguLESWvm333334241341231242414k12m4k1m24k1m2k4m1k24n1l2n41ln41lk2n4k12",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			resStr := pgo.ConcatFast(tc.s...)
			require.Equal(t, tc.result, resStr)
		})
	}
}

func BenchmarkConcatFast(b *testing.B) {
	s := generateRandomSliceOfStrings()
	for i := 0; i < b.N; i++ {
		pgo.ConcatFast(s...)
	}
}

func BenchmarkConcatFast2(b *testing.B) {
	s := generateRandomSliceOfStrings()
	for i := 0; i < b.N; i++ {
		stringBuilder(s...)
	}
}

func generateRandomSliceOfStrings() []string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	rand.Seed(time.Now().UnixNano())
	s := make([]string, 15)
	for i := range s {
		bt := make([]byte, 10)
		for j := range bt {
			bt[j] = letterBytes[rand.Intn(len(letterBytes))]
		}
		s[i] = string(bt)
	}

	return s
}

func stringBuilder(s ...string) string {
	l := len(s)
	if l == 0 {
		return ""
	}

	b := strings.Builder{}
	n := 0
	for i := 0; i < l; i++ {
		n += len(s[i])
	}

	b.Grow(n)
	for i := 0; i < l; i++ {
		b.WriteString(s[i])
	}

	return b.String()
}
