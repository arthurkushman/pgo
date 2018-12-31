package pgo

import (
	"net/http"
	"bytes"
	"io/ioutil"
)

const (
	Post = "POST"
	Get  = "GET"
)

func (c *Context) doRequest(path string) (string, error) {
	req, reqErr := http.NewRequest(c.RequestMethod, path, bytes.NewBuffer([]byte{}))

	if reqErr != nil {
		printError(reqErr.Error())
	}

	// setting headers
	for k, v := range c.Headers {
		req.Header.Set(k, v)
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		printError(err.Error())
	}

	defer resp.Body.Close()

	content, cErr := ioutil.ReadAll(resp.Body)

	if cErr != nil {
		printError(cErr.Error())
	}

	return string(content), nil
}
