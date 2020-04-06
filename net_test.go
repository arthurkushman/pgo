package pgo_test

import (
	"github.com/arthurkushman/pgo"
	"github.com/stretchr/testify/assert"
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
		assert.NoError(t, err)
		assert.Equal(t, long, v.long, "want %d, but got %d", long, v.long)
	}

	// long to ip
	for _, v := range testIP2Long {
		ip := pgo.Long2ip(v.long)
		assert.Equal(t, ip, v.ip, "want %d, but got %d", ip, v.ip)
	}
}

func TestGetMxrr(t *testing.T) {
	isMx, mxs, err := pgo.GetMxrr(DefaultOuterDomain)
	assert.NoError(t, err)
	assert.True(t, isMx)
	assert.Greaterf(t, len(mxs), 0, "want len(mxs) > 0, got %d", len(mxs))
}
