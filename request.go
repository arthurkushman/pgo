package pgo

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
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

	content, cErr := io.ReadAll(resp.Body)

	if cErr != nil {
		return string(content), cErr
	}

	return string(content), nil
}

func (c *Context) uploadFile(fieldName, filePath string) bool {
	var maxSize = int64(1 << 31)
	if c.UploadMaxFileSize > 0 {
		maxSize = c.UploadMaxFileSize
	}

	err := c.Req.ParseMultipartForm(maxSize)
	if err != nil {
		fmt.Println(err)
		return false
	}

	file, _, err := c.Req.FormFile(fieldName)
	if err != nil {
		fmt.Println(err)
		return false
	}

	defer file.Close()

	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return false
	}

	defer f.Close()
	_, err = io.Copy(f, file)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}
