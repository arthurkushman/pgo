package html_test

import (
	"pgo/html"
	"testing"
)

func TestStripTags(t *testing.T) {
	tests := []struct {
		in, out string
	}{
		{"<div>Lorem <span>ipsum dolor sit amet</span>, consectetur adipiscing elit, sed do eiusmod <a href=\"http://example.com\">tempor incididunt</a> ut labore et dolore magna aliqua.</div>",
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod <a href=\"http://example.com\">tempor incididunt</a> ut labore et dolore magna aliqua."},
	}

	for _, test := range tests {
		if got := html.StripTags(test.in, []string{"a"}); got != test.out {
			t.Fatalf("%q: want %q, got %q", test.in, test.out, got)
		}
	}
}
