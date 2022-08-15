// have tests up when writing and refactoring code
// make some notes about a testing strategy (what is it about this code that should be tested)
// think about some interesting cases (how will your code work if there are not letters, does it work for empty strings, big strings, etc.)
// go test package

package main

import (
	"testing"
)

var tests = []struct {
	input string
	want  string
}{
	{"hello", "uryyb"},
	{"HELLO", "URYYB"},
	{"", ""},
	{"incomprehensibility", "vapbzcerurafvovyvgl"},
	{"Hello there", "Uryyb gurer"},
	{"Hello there!", "word contains an invalid character"},
	{"y2k", "word contains an invalid character"},
}

func TestConvert(t *testing.T) {
	for _, test := range tests {
		if got := convert(test.input, makeKey()); got != test.want {
			t.Errorf("convert(%q) = %q", test.input, got)
		}
	}
}
