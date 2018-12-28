package pgo

import (
	"io/ioutil"
	"os"
)

// const mapping php -> go
const (
	FileAppend = os.ModeAppend
	LockEx     = os.ModeExclusive
)

// fileName name of file to where put data
// flags[0] - offset
// flags[1] - maxLen
func FileGetContents(fileName string, args ...interface{}) (string, error) {
	// write with both offset/maxLen
	if args[0].(int) > 0 && args[1].(int) > 0 {

	}

	// write with offset
	if args[0].(int) > 0 && args[1] == nil {

	}

	// write with maxLen
	if args[1].(int) > 0 && args[0] == nil {

	}

	data, err := ioutil.ReadFile(fileName)
	return string(data), err
}

// fileName name of file to where put data
// flags[0] - flags how to put this data FileAppend | LockEx
func FilePutContents(fileName, data string, flags ...interface{}) (int, error) {
	if flags[0] != nil {
		f, err := os.OpenFile(fileName, flags[0].(int), 0644)
		if err != nil {
			panic(err)
		}

		return f.WriteString(data)
	}

	return len(data), ioutil.WriteFile(fileName, []byte(data), os.FileMode(0644))
}
