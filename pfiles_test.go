package pgo_test

import (
	"testing"
	"pgo"
)

func TestFileGetContents(t *testing.T) {

}

func TestFilePutContents(t *testing.T) {
	str := "foo bar baz"
	n, err := pgo.FilePutContents("example.txt", str)

	if n != len(str) {
		t.Fatalf("want %d bytes of data, got %d", n, len(str))
	}

	if err != nil {
		t.Fatal(err)
	}
}