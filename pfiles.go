package pgo

import "io/ioutil"

// fileName name of file to where put data
// flags[0] - offset
// flags[1] - maxLen
func FileGetContents(fileName string, args... interface{}) (string, error) {
	data, err := ioutil.ReadFile(fileName)
	return string(data), err
}

// fileName name of file to where put data
// flags[0] - flags how to put this data FileAppend | LockEx
func FilePutContents(fileName, date string, flags... interface{}) {

}