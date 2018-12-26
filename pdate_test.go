package pgo_test

import (
	"testing"
	"pgo"
	"time"
)

func TestDate(t *testing.T) {
	if time.Now().Format("2006-02-01 15:04:05") != pgo.Date("Y-m-d H:i:s") {
		t.Fatal("Time formats in map doesn't match")
	}

	if time.Now().Format("2006-02-01T15:04:05") != pgo.Date("Y-m-dTH:i:s") {
		t.Fatal("Time formats in map doesn't match")
	}
}

func TestTime(t *testing.T) {

}

func TestDateTime(t *testing.T) {

}