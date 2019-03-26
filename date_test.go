package pgo_test

import (
	"pgo"
	"strconv"
	"testing"
	"time"
)

func TestDate(t *testing.T) {
	if time.Now().Format("2006-02-01 15:04:05") != pgo.Date("Y-m-d H:i:s") {
		t.Fatalf("want %s, got %s", time.Now().Format("2006-02-01 15:04:05"), pgo.Date("Y-m-d H:i:s"))
	}

	zone, _ := time.Now().Zone()
	if time.Now().Format("2006-02-01"+zone+"15:04:05") != pgo.Date("Y-m-dTH:i:s") {
		t.Fatalf("want %s, got %s", time.Now().Format("2006-02-01"+zone+"15:04:05"), pgo.Date("Y-m-dTH:i:s"))
	}

	if time.Now().Format("Mon, Jan") != pgo.Date("D, M") {
		t.Fatal("Time formats of week days and month in map doesn't match")
	}

	// test with unix timestamp passed as 2nd param
	if pgo.Date("Y-m-d H:i:s", time.Now().Unix()) == "" {
		t.Fatal("Time didn't parsed properly with unix timestamp")
	}
}

func TestSpecCases(t *testing.T) {
	if time.Now().Weekday().String() != pgo.Date("l") {
		t.Fatal("Weekday has not been matched")
	}

	yearDay, _ := strconv.Atoi(pgo.Date("z"))
	if time.Now().YearDay() != yearDay {
		t.Fatal("Year day has not been matched")
	}

	monthDay, _ := strconv.Atoi(pgo.Date("j"))
	if time.Now().Day() != monthDay {
		t.Fatal("Year day has not been matched")
	}

	weekDay, _ := strconv.Atoi(pgo.Date("N"))
	if int(time.Now().Weekday()) != weekDay {
		t.Fatal("Weekday has not been matched")
	}
}
