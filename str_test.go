package pgo_test

import (
	"pgo"
	"testing"
)

func TestStrReplace(t *testing.T) {
	subject := "The quick brown fox jumped over the lazy dog"

	subject, _ = pgo.StrReplace("fox", "cat", subject)
	subject, _ = pgo.StrReplace("dog", "fox", subject)

	result := "The quick brown cat jumped over the lazy fox"

	if subject != result {
		t.Fatalf("want %s, got %s", result, subject)
	}
}

func TestStrIReplace(t *testing.T) {
	subject := "The quick brown fox jumped over the lazy dog"

	subject, _ = pgo.StrIReplace("Fox", "cat", subject)
	subject, _ = pgo.StrIReplace([]string{"DOG", "QuiCK"}, []string{"fox", "slow"}, subject)

	result := "The slow brown cat jumped over the lazy fox"

	if subject != result {
		t.Fatalf("want %s, got %s", result, subject)
	}
}

func TestStrReplaceCount(t *testing.T) {
	subject := "The quick brown fox jumped over the lazy fox or dog"

	str, _ := pgo.StrReplace("fox", "cat", subject, 1)

	result := "The quick brown cat jumped over the lazy fox or dog"

	if str != result {
		t.Fatalf("want %s, got %s", result, subject)
	}
}

func TestStrReplaceArray(t *testing.T) {
	subject := "The quick brown fox jumped over the lazy dog"

	str, _ := pgo.StrReplace([]string{"fox", "dog"}, []string{"cat", "elephant"}, subject)

	result := "The quick brown cat jumped over the lazy elephant"

	if str != result {
		t.Fatalf("want %s, got %s", result, subject)
	}
}

func TestStrReplaceErrs(t *testing.T) {
	subject := "The quick brown fox jumped over the lazy dog"

	str, err := pgo.StrReplace([]string{"fox", "dog"}, "", subject)

	if err == nil && str != subject {
		t.Fatalf("want %s, got %s", subject, str)
	}
}

func TestHTTPBuildQuery(t *testing.T) {
	queryStr := pgo.HTTPBuildQuery(map[string]string{
		"foo": "bar",
		"bar": "baz",
	})

	want := "bar=baz&foo=bar"
	if queryStr != want {
		t.Fatalf("want %s, got %s", want, queryStr)
	}

	queryStr2 := pgo.HTTPBuildQuery(map[string]string{})

	want2 := ""
	if queryStr2 != want2 {
		t.Fatalf("want %s, got %s", want2, queryStr)
	}
}
