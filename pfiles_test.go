package pgo_test

import (
	"testing"
	"pgo"
)

const (
	strToWrite = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."
	fileName   = "example.txt"
)

func TestFileGetContents(t *testing.T) {
	n, err := pgo.FilePutContents(fileName, strToWrite)

	if err != nil {
		panic(err)
	}

	// reading full file with limit
	str, er := pgo.FileGetContents(fileName, 0, n)

	if er != nil {
		panic(er)
	}

	if len(str) != n {
		t.Fatalf("want %d bytes of data, got %d", n, len(str))
	}

	// reading offset/limit
	off := n / 3
	lim := n/2 - off - 1

	ss, er := pgo.FileGetContents(fileName, off, lim)

	if er != nil {
		panic(er)
	}

	if len(ss) != lim {
		t.Fatalf("want %d bytes of data, got %d", n, len(str))
	}

	sOff, errOff := pgo.FileGetContents(fileName, off)

	if errOff != nil {
		panic(errOff)
	}

	if len(sOff) != n-off {
		t.Fatalf("want %d bytes of data, got %d", len(sOff), n-off)
	}
}

func TestFilePutContents(t *testing.T) {
	// test write to file with append without options
	n1, err := pgo.FilePutContents(fileName, strToWrite)

	if err != nil {
		t.Fatal(err)
	}

	if n1 != len(strToWrite) {
		t.Fatalf("want %d bytes of data, got %d", n1, len(strToWrite))
	}

	// test write to file with append option
	n2, err2 := pgo.FilePutContents(fileName, strToWrite, pgo.FileAppend)

	if err2 != nil {
		t.Fatal(err2)
	}

	if n2 == n1*2 {
		t.Fatalf("want %d bytes of data, got %d", n2, n1*2)
	}
}
