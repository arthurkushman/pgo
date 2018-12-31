package pgo

import "strings"

type replaceParams struct {
	search  interface{}
	replace interface{}
	subject string
	count   int
}

// Replace all occurrences of the search string with the replacement string
// If search and replace are arrays, then str_replace() takes a value from each array
// and uses them to search and replace on subject.
func StrReplace(args ...interface{}) string {
	var rParams replaceParams

	countSlices := 0
	rParams.count = -1
	for i, p := range args {
		switch i {
		case 0:
			param, ok := p.(string)
			rParams.search = param
			if !ok {
				param, ok := p.([]string)
				isOk(ok, "You must provide search parameter as []string or string")
				countSlices++
				rParams.search = param
			}
			break
		case 1:
			param, ok := p.(string)
			rParams.replace = param
			if !ok {
				param, ok := p.([]string)
				isOk(ok, "You must provide replace parameter as []string or string")
				countSlices++
				rParams.replace = param
			}
			break
		case 2:
			param, ok := p.(string)
			rParams.subject = param
			isOk(ok, "3d parameter must be passed as string")
			break
		case 3:
			param, ok := p.(int)
			rParams.count = param
			isOk(ok, "4th parameter must be passed as int")
			break
		}
	}

	if countSlices == 1 {
		panic("Both slices must be provided for search and replace")
	}

	if countSlices == 2 {
		return rParams.doReplaceSlices()
	}

	return rParams.doReplace()
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
