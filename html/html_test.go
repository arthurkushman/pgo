package html_test

import (
	"github.com/arthurkushman/pgo/html"
	"testing"
)

var tests = []struct {
	in, out string
}{
	{"<div>Lorem <span>ipsum dolor sit amet</span>, consectetur adipiscing elit, sed do eiusmod <a href=\"http://example.com\">tempor incididunt</a> ut labore et dolore magna aliqua.</div>",
		"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod <a href=\"http://example.com\">tempor incididunt</a> ut labore et dolore magna aliqua."},
}

func TestStripTags(t *testing.T) {
	for _, test := range tests {
		if got := html.StripTags(test.in, []string{"a"}); got != test.out {
			t.Fatalf("%q: want %q, got %q", test.in, test.out, got)
		}
	}
}
