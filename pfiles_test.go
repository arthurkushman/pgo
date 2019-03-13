package pgo_test

import (
	"math"
	"os"
	"pgo"
	"testing"
)

const (
	strToWrite    = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."
	fileName      = "example.txt"
	defaultDomain = "http://localhost"
	file1         = "file1.txt"
	file2         = "file2.txt"
	dir1          = "dir1"
	dir2          = "dir2"
	symlink       = "symlink"
)

func TestFileGetContents(t *testing.T) {
	n, err := pgo.FilePutContents(fileName, strToWrite)

	if err != nil {
		panic(err)
	}

	// base case - readAll
	strBase, _ := pgo.FileGetContents(fileName)

	if len(strBase) != n {
		t.Fatalf("want %d bytes of data, got %d", n, len(strBase))
	}

	// reading full file with limit
	str, er := pgo.FileGetContents(fileName, nil, 0, n)

	if er != nil {
		panic(er)
	}

	if len(str) != n {
		t.Fatalf("want %d bytes of data, got %d", n, len(str))
	}

	// reading offset/limit
	off := n / 3
	lim := n/2 - off - 1

	ss, er := pgo.FileGetContents(fileName, nil, off, lim)

	if er != nil {
		panic(er)
	}

	if len(ss) != lim {
		t.Fatalf("want %d bytes of data, got %d", n, len(str))
	}

	sOff, errOff := pgo.FileGetContents(fileName, nil, off)

	if errOff != nil {
		panic(errOff)
	}

	if len(sOff) != n-off {
		t.Fatalf("want %d bytes of data, got %d", len(sOff), n-off)
	}
}

func TestFileGetContentsHttp(t *testing.T) {
	str, err := pgo.FileGetContents(defaultDomain)

	if err != nil {
		panic(err.Error())
	}

	if str == "" {
		t.Fatalf("want non-empty string, got %s", str)
	}
}

func TestFileGetContentsPanics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic with offset")
		}
	}()

	// panic with non existent file
	pgo.FileGetContents("non-existent.txt", nil, math.MaxInt64)

	// panic with non existent, but with limit
	pgo.FileGetContents("non-existent.txt", nil, math.MaxInt64, math.MaxInt64)

	// panic with existent but out of range offset
	pgo.FileGetContents(fileName, nil, math.MaxInt64)

	// panic with existent but out of range offset and limit
	pgo.FileGetContents(fileName, nil, math.MaxInt64, math.MaxInt64)
}

func TestFileGetContentsInvalidTypes(t *testing.T) {
	content, err := pgo.FileGetContents(fileName, nil, "", "")
	if content != "" && err.Error() != "Error on passing params with wrong types to FileGetContents: offset string and limit string" {
		t.Fatalf("Want an empty string, got: %s", content)
	}

	content2, err2 := pgo.FileGetContents(fileName, nil, "")
	if content2 != "" && err2.Error() != "Error on passing params to FileGetContents: offset string" {
		t.Fatalf("Want an empty string, got: %s", content2)
	}
}

func TestFileGetContentsHttpGetRequest(t *testing.T) {
	content, err := pgo.FileGetContents(defaultDomain, pgo.NewContext())

	if err != nil {
		t.Fatalf("Request failed with content: %s", content)
	}
}

func TestFileGetContentsHttpInvalidRequest(t *testing.T) {
	ctx := pgo.NewContext()
	ctx.RequestMethod = "INVALID()"

	_, err := pgo.FileGetContents(defaultDomain, ctx)

	if err == nil {
		t.Fatal("Request has not been failed with error")
	}

	ctx.RequestMethod = "OPTIONS"
	_, er := pgo.FileGetContents("https://abrakadabra.comz.ru", ctx)

	if er == nil {
		t.Fatal("Request has not been failed with error")
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

func TestFilePutContentsErrors(t *testing.T) {
	n1, err1 := pgo.FilePutContents(fileName, strToWrite, "")

	if n1 != -1 && err1.Error() != "Type of 3d parameter must be an int, got string" {
		t.Fatal("execution of FilePutContents has not been failed with error")
	}

	n2, err2 := pgo.FilePutContents("fakefile.out", "", 0x1212) // setting fake flags to invoke error from os.OpenFile

	if n2 != -1 && err2 == nil {
		t.Fatal("execution of FilePutContents has not been failed with error")
	}
}

// todo: Fix to working version of test for MoveUploadedFile
//func TestFilePutContents2(t *testing.T) {
//	fileHeader := make([]*multipart.FileHeader, 10)
//
//	req := http.Request{
//		Method: "POST",
//		MultipartForm: &multipart.Form{
//			Value: map[string][]string{
//				"foo": {"bar"},
//			},
//			File: map[string][]*multipart.FileHeader{
//				"foo": fileHeader,
//			},
//		},
//	}
//
//	c := &pgo.Context{
//		Req: &req,
//	}
//
//	c.MoveUploadedFile("foo", "./baz.txt")
//}

func TestFileExists(t *testing.T) {
	pgo.FilePutContents(file1, "foo bar baz")
	defer os.Remove(file1)

	f1 := pgo.FileExists(file1)
	if f1 != true {
		t.Fatalf("File exists and returning %v", f1)
	}

	f2 := pgo.FileExists(file2)
	if f2 != false {
		t.Fatalf("File doesn't exist and returning %v", f2)
	}
}

func TestIsDir(t *testing.T) {
	os.Mkdir(dir1, 0644)
	defer os.Remove(dir1)

	isDir := pgo.IsDir(dir1)
	if isDir != true {
		t.Fatalf("Directory dir1 is an existent dir but IsDir returned %v", isDir)
	}

	isDir2 := pgo.IsDir(dir2)
	if isDir2 != false {
		t.Fatalf("Directory "+dir2+" is non-existent dir but IsDir returned %v", isDir)
	}
}

func TestIsFile(t *testing.T) {
	pgo.FilePutContents(file1, "foo bar baz")
	defer os.Remove(file1)

	isFile := pgo.IsFile(file1)
	if isFile != true {
		t.Fatalf("File "+file1+" is a regular file, but IsFile returned %v", isFile)
	}

	isFile2 := pgo.IsFile(file2)
	if isFile2 != false {
		t.Fatalf("File "+file2+" is not a regular file, but IsFile returned %v", isFile)
	}
}

func TestIsLnik(t *testing.T) {
	pgo.FilePutContents(file1, "foo bar baz")
	os.Symlink(file1, symlink)
	defer os.Remove(file1)
	defer os.Remove(symlink)

	isSymlink := pgo.IsLink(symlink)
	if isSymlink != true {
		t.Fatalf(symlink+" is a symlinkr, but IsLink returned %v", isSymlink)
	}

	isSymlink2 := pgo.IsLink(file1)
	if isSymlink2 != false {
		t.Fatalf(file1+" is a regular file, but IsLink returned %v", isSymlink2)
	}
}
