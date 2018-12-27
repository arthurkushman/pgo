package pgo_test

import (
	"testing"
	"pgo"
)

func TestStrReplace(t *testing.T) {
	subject := "The quick brown fox jumped over the lazy dog"

	subject = pgo.StrReplace("fox", "cat", subject)
	subject = pgo.StrReplace("dog", "fox", subject)

	result := "The quick brown cat jumped over the lazy fox"

	if subject != result {
		t.Fatalf("want %s, got %s", result, subject)
	}
}

func TestStrReplaceArray(t *testing.T) {

}