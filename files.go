package pgo

import (
	"bufio"
	"errors"
	"fmt"
	"net/http"
	"os"
	"regexp"
)

// FileAppend const mapping php -> go
const (
	FileAppend = os.O_APPEND
)

// Context is the opts for http/net requests
type Context struct {
	Headers           map[string]string
	RequestMethod     string
	Req               *http.Request
	UploadMaxFileSize int64
}

// NewContext returns a new Context with preset default headers and request method
func NewContext() *Context {
	return &Context{
		Headers: map[string]string{
			"Accept":        "text/html",
			"Cache-Control": "max-age=0",
		},
		RequestMethod: "GET",
	}
}

// FileGetContents reads files, http requests streams
// path name of a file or domain from where to read data
// flags[0] - offset
// flags[1] - maxLen
func FileGetContents(path string, args ...interface{}) (string, error) {
	var argsLen = len(args)

	// determine if there is only an http get request
	match, _ := regexp.MatchString("http(s?)\\:", path)
	if argsLen == 0 && match {
		context := NewContext()

		return context.doRequest(path)
	}

	if argsLen > 0 && args[0] != nil { // context has been passed - send http request
		context, ok := args[0].(*Context)

		if !ok {
			printError("Context param must be of type Context, %T passed", args[0])
		}

		return context.doRequest(path)
	}

	// write with both offset/maxLen
	if argsLen == 3 {
		offset, okOff := args[1].(int)
		limit, okLim := args[2].(int)

		if !okOff || !okLim {
			errMsg := fmt.Sprintf("Error on passing params with wrong types to FileGetContents: offset %T and limit %T", args[1], args[2])
			return "", errors.New(errMsg)
		}

		reader, err := initReader(path)
		if err != nil {
			return "", err
		}

		_, err = reader.Discard(offset) // skipping an offset from user input
		if err != nil {
			return "", err
		}

		buf := make([]byte, limit)
		_, e := reader.Read(buf)

		return string(buf), e
	}

	// write with offset
	if argsLen == 2 {
		offset, ok := args[1].(int)

		if !ok {
			errMsg := fmt.Sprintf("Error on passing params to FileGetContents: offset %T", args[1])
			return "", errors.New(errMsg)
		}

		f, fErr := os.Open(path)

		if fErr != nil {
			printError(fErr.Error())
		}

		reader := bufio.NewReader(f)

		fInfo, fiErr := f.Stat()

		if fiErr != nil {
			return "", fiErr
		}

		buf := make([]byte, int(fInfo.Size())-offset)
		_, err := reader.Discard(offset) // skipping an offset from user input

		if err != nil {
			printError(err.Error())
		}

		_, e := reader.Read(buf)

		return string(buf), e
	}

	data, err := os.ReadFile(path)
	return string(data), err
}

func initReader(fileName string) (*bufio.Reader, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	return bufio.NewReader(f), err
}

// FilePutContents write files with offset/limit
// fileName name of file to where put data
// flags[0] - flags how to put this data FileAppend | LockEx
func FilePutContents(fileName, data string, flags ...interface{}) (int, error) {
	if len(flags) > 0 {
		v, ok := flags[0].(int)

		if !ok {
			return -1, fmt.Errorf("type of 3d parameter must be an int, got %T", flags[0])
		}

		f, err := os.OpenFile(fileName, v|os.O_WRONLY, 0644)
		defer func(f *os.File) {
			err = f.Close()
			if err != nil {
				fmt.Println(err)
			}
		}(f)

		if err != nil {
			return -1, err
		}

		return f.WriteString(data)
	}

	return len(data), os.WriteFile(fileName, []byte(data), os.FileMode(0644))
}

// MoveUploadedFile uploads file from fieldName to destination path
func (c *Context) MoveUploadedFile(fieldName, destination string) bool {
	return c.uploadFile(fieldName, destination)
}

// FileExists checks wether file or catalog exists or not
// returning true or false respectfully
func FileExists(fileName string) bool {
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return false
	}

	return err == nil
}

// IsDir tells whether the filename is a directory
func IsDir(fileName string) bool {
	fi, err := os.Stat(fileName)
	if err != nil {
		return false
	}

	return fi.Mode().IsDir()
}

// IsFile tells whether the filename is a regular file
func IsFile(fileName string) bool {
	fi, err := os.Stat(fileName)
	if err != nil {
		return false
	}

	return fi.Mode().IsRegular()
}

// IsLink tells whether the filename is a symbolic link
func IsLink(fileName string) bool {
	_, err := os.Readlink(fileName)
	return err == nil
}
