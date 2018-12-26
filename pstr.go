package pgo

import "strings"

type ReplaceParams struct {
	search  interface{}
	replace interface{}
	subject string
}

func StrReplace(args ...interface{}) string {
	var replaceParams ReplaceParams

	countSlices := 0

	for i, p := range args {
		switch i {
		case 0:
			param, ok := p.(string)
			replaceParams.search = param
			if !ok {
				param, ok := p.([]string)
				isOk(ok, "You must provide search parameter as []string or string")
				countSlices++
				replaceParams.search = param
			}
			break
		case 1:
			param, ok := p.(string)
			replaceParams.replace = param
			if !ok {
				param, ok := p.([]string)
				isOk(ok, "You must provide replace parameter as []string or string")
				countSlices++
				replaceParams.replace = param
			}
			break
		case 2:
			param, ok := p.(string)
			replaceParams.subject = param
			isOk(ok, "3d parameter must be passed as string")
			break
		}
	}

	if countSlices == 1 {
		panic("Both slices must be provided for search and replace")
	}

	if countSlices == 2 {
		return replaceParams.doReplaceSlices()
	}

	return replaceParams.doReplace()
}

func (params *ReplaceParams) doReplace() string {


	return params.subject
}

func (params *ReplaceParams) doReplaceSlices() string {


	return params.subject
}