package pgo

import (
	"regexp"
	"strconv"
	"time"
)

const (
	DivideMicroseconds = 1e3
	DivideMilliseconds = 1e6
)

var specCases map[string]interface{}

type goDate struct {
	parsedSymbols   []string
	inputDateFormat string
	t               time.Time
	unix            int64
	phpToGoFormat   map[string]string
}

type Time time.Time

const phpDateFormatSymbols = "[\\D]"

// Date returns formatted output of system data/time
// 1st argument formatted string e.g.: Y-m-d H:i:s
// 2nd argument int64 unix timestamp
func Date(args ...interface{}) string {
	var date goDate

	if len(args) == 0 { // return default "2006-01-02 15:04:05.999999999 -0700 MST" formatted string
		return date.t.String()
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
		"m": "01",
		"d": "02",
		"H": "15",
		"i": "04",
		"s": "05",
		"D": "Mon",
		"l": "Monday",
		"M": "Jan",
		"F": "January",
		"r": time.RubyDate,
		"c": time.RFC3339,
		"A": "PM",
	}

	// todo: move all the spec cases to parser, because of waste of the resources
	_, isoWeek := date.t.ISOWeek()
	zone, offset := date.t.Zone()
	// determine quarter of a year
	q := 1
	m := int(date.t.Month())
	if m > 3 && m <= 6 {
		q = 2
	} else if m > 6 && m <= 9 {
		q = 3
	} else if m > 9 && m <= 12 {
		q = 4
	}
	specCases = map[string]interface{}{
		"l": date.t.Weekday().String(),
		"N": int(date.t.Weekday()),
		"z": date.t.YearDay(),
		"j": date.t.Day(),
		"W": isoWeek,
		"Z": offset,
		"T": zone,
		"Q": q,
	}
}

// Milliseconds from time.Time Go type
func (t Time) Milliseconds() int64 {
	return time.Time(t).UnixNano() / DivideMilliseconds
}

// Microseconds from time.Time Go type
func (t Time) Microseconds() int64 {
	return time.Time(t).UnixNano() / DivideMicroseconds
}

// UnixMilli milliseconds since unix epoch
func UnixMilli() int64 {
	return time.Now().UnixNano() / DivideMilliseconds
}

// UnixMicro microseconds since unix epoch
func UnixMicro() int64 {
	return time.Now().UnixNano() / DivideMicroseconds
}
