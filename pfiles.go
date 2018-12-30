package pgo

import (
	"io/ioutil"
	"os"
	"bufio"
)

// const mapping php -> go
const (
	FileAppend = os.O_APPEND
	LockEx     = os.O_EXCL
)

// FileGetContents reads files, http requests streams
// fileName name of file to where put data
// flags[0] - offset
// flags[1] - maxLen
func FileGetContents(fileName string, args ...interface{}) (string, error) {
	var argsLen = len(args)

	// write with both offset/maxLen
	if argsLen == 2 && args[0].(int) > 0 && args[1].(int) > 0 {
		offset := args[0].(int)
		limit := args[1].(int)

		reader := initReader(fileName)
		_, err := reader.Discard(offset) // skipping an offset from user input

		if err != nil {
			panic(err)
		}

		buf := make([]byte, limit)
		_, e := reader.Read(buf)

		return string(buf), e
	}

	// write with offset
	if argsLen == 1 && args[0].(int) > 0 {
		offset := args[0].(int)

		f, fErr := os.Open(fileName)

		if fErr != nil {
			panic(fErr)
		}

		reader := bufio.NewReader(f)

		fInfo, fiErr := f.Stat()

		if fiErr != nil {
			panic(fiErr)
		}

		buf := make([]byte, int(fInfo.Size()) - offset)
		_, err := reader.Discard(args[0].(int)) // skipping an offset from user input

		if err != nil {
			panic(err)
		}

		_, e := reader.Read(buf)

		return string(buf), e
	}

	data, err := ioutil.ReadFile(fileName)
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
	if len(flags) > 0 && flags[0] != nil {
		f, err := os.OpenFile(fileName, flags[0].(int) | os.O_WRONLY, 0644)
		defer f.Close()

		if err != nil {
			panic(err)
		}

		return f.WriteString(data)
	}

	return len(data), ioutil.WriteFile(fileName, []byte(data), os.FileMode(0644))
}
