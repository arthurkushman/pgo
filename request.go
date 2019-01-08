package pgo

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

const (
	Post = "POST"
	Get  = "GET"
)

func (c *Context) doRequest(path string) (string, error) {
	req, reqErr := http.NewRequest(c.RequestMethod, path, bytes.NewBuffer([]byte{}))

	if reqErr != nil {
		return "", reqErr
	}

	// setting headers
	for k, v := range c.Headers {
		req.Header.Set(k, v)
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	content, cErr := ioutil.ReadAll(resp.Body)

	if cErr != nil {
		return string(content), cErr
	}

	return string(content), nil
}
