package pgo_test

import (
	"pgo"
	"testing"
)

const DefaultOuterDomain = "google.com"

var testIP2Long = []struct {
	long uint32
	ip   string
}{
	{2130706433, "127.0.0.1"},
	{3927057275, "234.18.35.123"},
	{2956665461, "176.59.34.117"},
}

func TestIP2long(t *testing.T) {

	// ip to long
	for _, v := range testIP2Long {
		long, err := pgo.IP2long(v.ip)

		if err != nil {
			t.Error(err)
		}

		if long != v.long {
			t.Fatalf("want %d, but got %d", long, v.long)
		}
	}

	// long to ip
	for _, v := range testIP2Long {
		ip := pgo.Long2ip(v.long)

		if ip != v.ip {
			t.Fatalf("want %s, but got %s", ip, v.ip)
		}
	}
}

func TestGetMxrr(t *testing.T) {
	isMx, mxs, err := pgo.GetMxrr(DefaultOuterDomain)
	if err != nil {
		t.Error(err)
	}
	if !isMx && len(mxs) <= 0 {
		t.Fatalf("want true, got %t and want len(mxs) > 0, got %d", isMx, len(mxs))
	}

}