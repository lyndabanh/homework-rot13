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
		got, got_err := k.encrypt(test.input)
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

	for _, test := range tests {
		got, got_err := k.decrypt(test.input)
		if got != test.want {
			t.Errorf("decrypt(%q) = %q, expected %q", test.input, got, test.want)
		}

		if test.err == nil && got_err == nil {
			continue
		} else if test.err == nil && got_err != nil {
			t.Errorf("Expected no error, but got %#v", got_err)
		} else if test.err != nil && got_err == nil {
			t.Errorf("Expected error %#v, but got nill", test.err)
		} else if test.err.Error() != got_err.Error() {
			t.Errorf("Expected error %#v but got %#v", test.err, got_err)
		}
	}
}
