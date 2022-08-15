// have tests up when writing and refactoring code
// make some notes about a testing strategy (what is it about this code that should be tested)
// think about some interesting cases (how will your code work if there are not letters, does it work for empty strings, big strings, etc.)
// go test package

package main

import (
	"fmt"
	"testing"
)

var tests = []struct {
	input string
	want  string
	err   error
}{
	{input: "hello", want: "uryyb", err: nil},
	{input: "HELLO", want: "URYYB", err: nil},
	{input: "", want: "", err: nil},
	{input: "incomprehensibility", want: "vapbzcerurafvovyvgl", err: nil},
	{input: "Hello there", want: "Uryyb gurer", err: nil},
	{input: "Hello there!", want: "", err: fmt.Errorf("word contains an invalid character")},
	{input: "y2k", want: "", err: fmt.Errorf("word contains an invalid character")},
}

func TestConvert(t *testing.T) {
	k := makeKey()
	for _, test := range tests {
		got, got_err := k.encrypt(test.input, k.encryptionKey)
		if got != test.want {
			t.Errorf("encrypt(%q) = %q, expected: %q", test.input, got, test.want)
		}

		if test.err == nil && got_err == nil {
			continue
		} else if test.err == nil && got_err != nil {
			t.Errorf("Expected no error, but got %#v", got_err)
		} else if test.err != nil && got_err == nil {
			t.Errorf("Expected error %#v, but got nil", test.err)
		} else if test.err.Error() != got_err.Error() {
			t.Errorf("Expected error %#v but got %#v", test.err, got_err)
		}
	}
}
