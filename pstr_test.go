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
