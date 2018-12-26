package pgo

type ReplaceParams struct {
	search  interface{}
	replace interface{}
	subject string
}

func StrReplace(args ...interface{}) {
	var replaceParams ReplaceParams

	for i, p := range args {
		switch i {
		case 0:
			param, ok := p.(string)
			replaceParams.search = param
			if !ok {
				param, ok := p.([]string)
				isOk(ok, "You must provide search parameter as []string or string")
				replaceParams.search = param
			}
			break
		case 1:
			param, ok := p.(string)
			replaceParams.replace = param
			if !ok {
				param, ok := p.([]string)
				isOk(ok, "You must provide replace parameter as []string or string")
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
}

func (params *ReplaceParams) doReplace() string {


	return params.subject
}