package pgo

import (
	"regexp"
	"strconv"
	"time"
)

var specCases map[string]interface{}

type goDate struct {
	parsedSymbols   []string
	inputDateFormat string
	t               time.Time
	unix            int64
	phpToGoFormat   map[string]string
}

const phpDateFormatSymbols = "[\\D]"

// Date returns formatted output of system data/time
// 1st argument formatted string e.g.: Y-m-d H:i:s
// 2nd argument int64 unix timestamp
func Date(args ...interface{}) string {
	var date goDate

	if len(args) == 0 {
		panic("At least 1st parameter format must be passed")
	}

	for i, p := range args {
		switch i {
		case 0:
			param, ok := p.(string)
			isOk(ok, "You must provide format string parameter as string")
			date.inputDateFormat = param
			break
		case 1:
			param, ok := p.(int64)
			isOk(ok, "You must provide timestamp as int64 type")
			date.unix = param
			break
		}
	}

	date.t = time.Now()
	if date.unix > 0 {
		// parse unix representation with int type to Time object
		date.t = time.Unix(date.unix, 0)
	}

	date.initMapping()
	return date.parse()
}

func (date *goDate) parse() string {
	date.convert()

	var convertedString string
	for _, v := range date.parsedSymbols {
		if val, ok := date.phpToGoFormat[v]; ok {
			convertedString += date.t.Format(val)
		} else if sVal, ok := specCases[v]; ok {
			v, ok := sVal.(int)
			if ok {
				convertedString += strconv.Itoa(v)
			} else {
				convertedString += sVal.(string)
			}
		} else {
			convertedString += date.t.Format(v)
		}
	}

	return convertedString
}

func (date *goDate) convert() {
	r, _ := regexp.Compile(phpDateFormatSymbols)

	for _, chSlice := range r.FindAllStringSubmatch(date.inputDateFormat, -1) {
		for _, ch := range chSlice {
			date.parsedSymbols = append(date.parsedSymbols, ch)
		}
	}
}

// initializes date formats mapping between go and php
func (date *goDate) initMapping() {
	date.phpToGoFormat = map[string]string{
		"Y": "2006",
		"m": "02",
		"d": "01",
		"H": "15",
		"i": "04",
		"s": "05",
		"D": "Mon",
		"M": "Jan",
		"r": time.RubyDate,
		"c": time.RFC3339,
	}

	_, isoWeek := date.t.ISOWeek()
	specCases = map[string]interface{}{
		"l": date.t.Weekday().String(),
		"N": int(date.t.Weekday()),
		"z": date.t.YearDay(),
		"j": date.t.Day(),
		"W": isoWeek,
	}
}
