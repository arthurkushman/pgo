package pgo

import (
	"io/ioutil"
	"os"
	"bufio"
	"regexp"
)

// const mapping php -> go
const (
	FileAppend = os.O_APPEND
	LockEx     = os.O_EXCL
)

// Context is the opts for http/net requests
type Context struct {
	Headers       map[string]string
	RequestMethod string
}

// FileGetContents reads files, http requests streams
// fileName name of file to where put data
// flags[0] - offset
// flags[1] - maxLen
func FileGetContents(path string, args ...interface{}) (string, error) {
	var argsLen = len(args)

	// determine if there is only an http get request
	match, _ := regexp.MatchString("http(s?)\\:", path)
	if argsLen == 0 && match {
		context := &Context{
			RequestMethod:Get,
		}

		return context.doRequest(path)
	}

	if argsLen > 0 && args[0] != nil { // context has been passed - send http request
		context, ok := args[0].(Context)

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
			printError("Error on passing params with wrong types to FileGetContents: offset %T and limit %T", args[1], args[2])
		}

		reader := initReader(path)
		_, err := reader.Discard(offset) // skipping an offset from user input

		if err != nil {
			panic(err)
		}

		buf := make([]byte, limit)
		_, e := reader.Read(buf)

		return string(buf), e
	}

	// write with offset
	if argsLen == 2 {
		offset, ok := args[1].(int)

		if !ok {
			printError("Error on passing params to FileGetContents: offset %v", ok)
		}

		f, fErr := os.Open(path)

		if fErr != nil {
			printError(fErr.Error())
		}

		reader := bufio.NewReader(f)

		fInfo, fiErr := f.Stat()

		if fiErr != nil {
			panic(fiErr)
		}

		buf := make([]byte, int(fInfo.Size())-offset)
		_, err := reader.Discard(offset) // skipping an offset from user input

		if err != nil {
			printError(err.Error())
		}

		_, e := reader.Read(buf)

		return string(buf), e
	}

	data, err := ioutil.ReadFile(path)
	return string(data), err
}

func initReader(fileName string) *bufio.Reader {
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	return bufio.NewReader(f)
}

// FilePutContents write files with offset/limit
// fileName name of file to where put data
// flags[0] - flags how to put this data FileAppend | LockEx
func FilePutContents(fileName, data string, flags ...interface{}) (int, error) {
	if len(flags) > 0 {
		v, ok := flags[0].(int)

		isOk(ok, "Type of 3d parameter must be an int, got %T", flags[0])

		f, err := os.OpenFile(fileName, v|os.O_WRONLY, 0644)
		defer f.Close()

		if err != nil {
			panic(err)
		}

		return f.WriteString(data)
	}

	return len(data), ioutil.WriteFile(fileName, []byte(data), os.FileMode(0644))
}
