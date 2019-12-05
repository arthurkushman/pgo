package pgo_test

import (
	"github.com/stretchr/testify/assert"
	"pgo"
	"strconv"
	"testing"
	"time"
)

func TestDate(t *testing.T) {
	assert.Equalf(t, time.Now().Format("2006-02-01 15:04:05"), pgo.Date("Y-m-d H:i:s"), "want %s, got %s", time.Now().Format("2006-02-01 15:04:05"), pgo.Date("Y-m-d H:i:s"))

	zone, _ := time.Now().Zone()
	assert.Equalf(t, time.Now().Format("2006-02-01"+zone+"15:04:05"), pgo.Date("Y-m-dTH:i:s"), "want %s, got %s", time.Now().Format("2006-02-01"+zone+"15:04:05"), pgo.Date("Y-m-dTH:i:s"))

	assert.Equal(t, time.Now().Format("Mon, Jan"), pgo.Date("D, M"), "Time formats of week days and month in map doesn't match")

	// test with unix timestamp passed as 2nd param
	assert.NotEmpty(t, pgo.Date("Y-m-d H:i:s", time.Now().Unix()), "Time didn't parsed properly with unix timestamp")
}

func TestSpecCases(t *testing.T) {
	assert.Equal(t, time.Now().Weekday().String(), pgo.Date("l"), "Weekday has not been matched")

	yearDay, err := strconv.Atoi(pgo.Date("z"))
	assert.NoError(t, err)
	assert.Equal(t, time.Now().YearDay(), yearDay, "Year day has not been matched")

	monthDay, err := strconv.Atoi(pgo.Date("j"))
	assert.NoError(t, err)
	assert.Equal(t, time.Now().Day(), monthDay, "Year day has not been matched")

	weekDay, err := strconv.Atoi(pgo.Date("N"))
	assert.NoError(t, err)
	assert.Equal(t, int(time.Now().Weekday()), weekDay, "Weekday has not been matched")

	quarter, err := strconv.Atoi(pgo.Date("Q"))
	assert.NoError(t, err)
	q := 1
	m := int(time.Now().Month())
	if m > 3 && m <= 6 {
		q = 2
	} else if m > 6 && m <= 9 {
		q = 3
	} else if m > 9 && m <= 12 {
		q = 4
	}
	assert.Equalf(t, q, quarter, "want: %s, got: %s", q, quarter)
}

func TestTime_Milliseconds(t *testing.T) {
	now := time.Now()
	nowMicro := pgo.Time(now.Add(time.Microsecond * 3)).Microseconds()
	nowPlus := (now.UnixNano() / (int64(time.Microsecond)/int64(time.Nanosecond))) + 3
	assert.Equal(t, nowMicro, nowPlus)
}
