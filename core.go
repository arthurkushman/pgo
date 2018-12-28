package pgo

type PDate interface {
	Date(args ...interface{}) string
	parse() string
	convert()
}

type PStr interface {
	StrReplace(args ...interface{}) string
}

func isOk(ok bool, msg string) {
	if !ok {
		panic(msg)
	}
}