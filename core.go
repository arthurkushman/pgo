package pgo

type PDate interface {
	Date() string
	parse() string
	convert()
}

type PStr interface {
	StrReplace() string
}

type PRegExp interface {
	PregReplace() string
	PregMatch() []string
}

func isOk(ok bool, msg string) {
	if !ok {
		panic(msg)
	}
}