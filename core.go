package pgo

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"fmt"
	"os"
)

func isOk(ok bool, msg string, args ...interface{}) {
	if !ok {
		printError(msg, args)
	}
}

func printError(msg string, args ...interface{}) {
	fmt.Printf(msg, []interface{}(args)) // strange fix but it didn't work on go version go1.11.4 darwin/amd64 with args...
	if os.Getenv("PGO_ENV") != "dev" {
		os.Exit(1)
	}
}

// Serialize encodes Go code entities to string for e.g.: later storage
func Serialize(val interface{}) (string, error) {
	b := bytes.Buffer{}
	err := gob.NewEncoder(&b).Encode(val)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(b.Bytes()), nil
}

// Unserialize decodes string back into Go code representation
func Unserialize(val string, v interface{}) error {
	by, err := base64.StdEncoding.DecodeString(val)
	if err != nil {
		return err
	}

	b := new(bytes.Buffer)
	b.Write(by)
	err = gob.NewDecoder(b).Decode(v)
	if err != nil {
		return err
	}

	return nil
}
