package pgo_test

import (
	"math"
	"os"
	"testing"

	"github.com/arthurkushman/pgo"
	"github.com/stretchr/testify/assert"
)

const (
	strToWrite    = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."
	fileName      = "example.txt"
	defaultDomain = "http://google.com"
	file1         = "file1.txt"
	file2         = "file2.txt"
	dir1          = "dir1"
	dir2          = "dir2"
	symlink       = "symlink"
)

func TestFileGetContents(t *testing.T) {
	n, err := pgo.FilePutContents(fileName, strToWrite)
	assert.NoError(t, err)

	// base case - readAll
	strBase, err := pgo.FileGetContents(fileName)
	assert.NoError(t, err)
	assert.Equalf(t, len(strBase), n, "want %d bytes of data, got %d", n, len(strBase))

	// reading full file with limit
	str, err := pgo.FileGetContents(fileName, nil, 0, n)
	assert.NoError(t, err)
	assert.Equalf(t, len(str), n, "want %d bytes of data, got %d", n, len(str))

	// reading offset/limit
	off := n / 3
	lim := n/2 - off - 1
	ss, err := pgo.FileGetContents(fileName, nil, off, lim)
	assert.NoError(t, err)
	assert.Equalf(t, len(ss), lim, "want %d bytes of data, got %d", n, len(str))

	sOff, err := pgo.FileGetContents(fileName, nil, off)
	assert.NoError(t, err)
	assert.Equalf(t, len(sOff), n-off, "want %d bytes of data, got %d", n-off, len(sOff))
}

func TestFileGetContentsHttp(t *testing.T) {
	str, err := pgo.FileGetContents(defaultDomain)
	assert.NoError(t, err)
	assert.NotEmpty(t, str)
}

func TestFileGetContentsPanics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic with offset")
		}
	}()

	// panic with non existent file
	_, err := pgo.FileGetContents("non-existent.txt", nil, math.MaxInt64)
	assert.NoError(t, err)

	// panic with non existent, but with limit
	_, err = pgo.FileGetContents("non-existent.txt", nil, math.MaxInt64, math.MaxInt64)
	assert.NoError(t, err)

	// panic with existent but out of range offset
	_, err = pgo.FileGetContents(fileName, nil, math.MaxInt64)
	assert.NoError(t, err)

	// panic with existent but out of range offset and limit
	_, err = pgo.FileGetContents(fileName, nil, math.MaxInt64, math.MaxInt64)
	assert.NoError(t, err)
}

func TestFileGetContentsInvalidTypes(t *testing.T) {
	content, err := pgo.FileGetContents(fileName, nil, "", "")
	assert.Empty(t, content)
	assert.EqualError(t, err, "Error on passing params with wrong types to FileGetContents: offset string and limit string")

	content2, err2 := pgo.FileGetContents(fileName, nil, "")
	assert.Empty(t, content2)
	assert.EqualError(t, err2, "Error on passing params to FileGetContents: offset string")
}

func TestFileGetContentsHttpGetRequest(t *testing.T) {
	content, err := pgo.FileGetContents(defaultDomain, pgo.NewContext())
	assert.NoError(t, err)
	assert.NotEmpty(t, content)
}

func TestFileGetContentsHttpInvalidRequest(t *testing.T) {
	ctx := pgo.NewContext()
	ctx.RequestMethod = "INVALID()"

	_, err := pgo.FileGetContents(defaultDomain, ctx)
	assert.Errorf(t, err, "Request has not been failed with error")

	ctx.RequestMethod = "OPTIONS"
	_, er := pgo.FileGetContents("https://abrakadabra.comz.ru", ctx)
	assert.Errorf(t, er, "Request has not been failed with error")
}

func TestFilePutContents(t *testing.T) {
	// test write to file with append without options
	n1, err := pgo.FilePutContents(fileName, strToWrite)
	assert.NoError(t, err)
	assert.Equalf(t, n1, len(strToWrite), "want %d bytes of data, got %d", n1, len(strToWrite))

	// test write to file with append option
	n2, err := pgo.FilePutContents(fileName, strToWrite, pgo.FileAppend)
	assert.NoError(t, err)
	assert.Equalf(t, n2, len(strToWrite), "want %d bytes of data, got %d", n2, len(strToWrite))
}

func TestFilePutContentsErrors(t *testing.T) {
	n1, err := pgo.FilePutContents(fileName, strToWrite, "")
	assert.EqualError(t, err, "type of 3d parameter must be an int, got string")
	assert.Equal(t, n1, -1)

	n2, err := pgo.FilePutContents("fakefile.out", "", 0x1212) // setting fake flags to invoke error from os.OpenFile
	assert.Error(t, err)
	assert.Equal(t, n2, -1)
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
	_, err := pgo.FilePutContents(file1, "foo bar baz")
	assert.NoError(t, err)
	defer os.Remove(file1)

	f1 := pgo.FileExists(file1)
	assert.Truef(t, f1, "File exists and returning %v", f1)

	f2 := pgo.FileExists(file2)
	assert.Falsef(t, f2, "File doesn't exist and returning %v", f2)
}

func TestIsDir(t *testing.T) {
	err := os.Mkdir(dir1, 0644)
	assert.NoError(t, err)
	defer os.Remove(dir1)

	isDir := pgo.IsDir(dir1)
	assert.Truef(t, isDir, "Directory dir1 is an existent dir but IsDir returned %v", isDir)

	isDir2 := pgo.IsDir(dir2)
	assert.Falsef(t, isDir2, "Directory "+dir2+" is non-existent dir but IsDir returned %v", isDir)
}

func TestIsFile(t *testing.T) {
	_, err := pgo.FilePutContents(file1, "foo bar baz")
	assert.NoError(t, err)
	defer os.Remove(file1)

	isFile := pgo.IsFile(file1)
	assert.Truef(t, isFile, "File "+file1+" is a regular file, but IsFile returned %v", isFile)

	isFile2 := pgo.IsFile(file2)
	assert.Falsef(t, isFile2, "File "+file2+" is not a regular file, but IsFile returned %v", isFile2)
}

func TestIsLnik(t *testing.T) {
	_, err := pgo.FilePutContents(file1, "foo bar baz")
	assert.NoError(t, err)
	err = os.Symlink(file1, symlink)
	assert.NoError(t, err)

	defer os.Remove(file1)
	defer os.Remove(symlink)

	isSymlink := pgo.IsLink(symlink)
	assert.Truef(t, isSymlink, symlink+" is a symlinkr, but IsLink returned %v", isSymlink)

	isSymlink2 := pgo.IsLink(file1)
	assert.Falsef(t, isSymlink2, file1+" is a regular file, but IsLink returned %v", isSymlink2)
}
