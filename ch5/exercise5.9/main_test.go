package main

import (
	"strings"
	"testing"
)

var cases = []struct {
	input    string
	f        func(string) string
	expected string
}{
	{"this is a $test", strings.ToUpper, "this is a TEST"},
	{"this is a $test", strings.ToLower, "this is a test"},
	{"--$test--", strings.ToUpper, "--TEST--"},
	{"--$test_test-test", strings.ToUpper, "--TEST_TEST-test"},
}

func TestExpand(t *testing.T) {
	for _, tc := range cases {
		result := expand(tc.input, tc.f)
		if result != tc.expected {
			t.Fatalf("Expected '%s' but got '%s'", tc.expected, result)
		}
	}
}
