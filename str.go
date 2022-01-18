package pgo

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
)

type replaceParams struct {
	search      interface{}
	replace     interface{}
	subject     string
	count       int
	countSlices int
}

// StrIReplace replaces all occurrences of the search case-insensitive string|slice with the replacement string
// If search and replace are arrays, then str_replace() takes a value from each array
// and uses them to search and replace on subject.
func StrIReplace(args ...interface{}) (string, error) {
	params := &replaceParams{}
	params.prepareParams(args...)

	if params.countSlices == 2 { // prepare an array to lower
		search := params.search.([]string)
		for k, v := range search {
			search[k] = strings.ToLower(v)
		}
		params.search = search

		return params.doReplaceSlices(), nil
	}

	// prepare string to lower
	val, _ := params.search.(string)
	search := strings.ToLower(val)
	params.search = search

	return params.doReplace(), nil
}

func (params *replaceParams) prepareParams(args ...interface{}) {
	params.count = -1
	params.countSlices = 0

	for i, p := range args {
		switch i {
		case 0:
			param, ok := p.(string)
			params.search = param
			if !ok {
				param, ok := p.([]string)
				isOk(ok, "You must provide search parameter as []string or string")
				params.countSlices++
				params.search = param
			}
		case 1:
			param, ok := p.(string)
			params.replace = param
			if !ok {
				param, ok := p.([]string)
				isOk(ok, "You must provide replace parameter as []string or string")
				params.countSlices++
				params.replace = param
			}
		case 2:
			param, ok := p.(string)
			params.subject = param
			isOk(ok, "3d parameter must be passed as string")
		case 3:
			param, ok := p.(int)
			params.count = param
			isOk(ok, "4th parameter must be passed as int")
		}
	}
}

// StrReplace replaces all occurrences of the search string|slice with the replacement string
// If search and replace are arrays, then str_replace() takes a value from each array
// and uses them to search and replace on subject.
func StrReplace(args ...interface{}) (string, error) {
	params := &replaceParams{}
	params.prepareParams(args...)

	if params.countSlices == 1 {
		return params.subject, errors.New("both slices must be provided for search and replace")
	}

	if params.countSlices == 2 {
		return params.doReplaceSlices(), nil
	}

	return params.doReplace(), nil
}

func (params *replaceParams) doReplace() string {
	return strings.Replace(params.subject, params.search.(string), params.replace.(string), params.count)
}

func (params *replaceParams) doReplaceSlices() string {
	search := params.search.([]string)
	replace := params.replace.([]string)

	for k, v := range search {
		params.subject = strings.Replace(params.subject, v, replace[k], params.count)
	}

	return params.subject
}

// HTTPBuildQuery Generates a URL-encoded query string from the associative map
func HTTPBuildQuery(pairs map[string]interface{}) string {
	q := url.Values{}
	for k, v := range pairs {
		switch val := v.(type) {
		case string:
			q.Add(k, v.(string))
		case []string:
			q[k] = v.([]string)
		case uint, uint8, uint16, uint32, uint64, int, int8, int16, int32, int64:
			q.Add(k, fmt.Sprintf("%d", val))
		case float32, float64:
			q.Add(k, fmt.Sprintf("%v", val))
		case bool:
			q.Add(k, fmt.Sprintf("%t", val))
		}
	}

	return q.Encode()
}
