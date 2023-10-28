package tests

import (
	"matching_prefixes/prefixes"
	"testing"
)

func TestFindLongestPrefix(t *testing.T) {
	// Test cases with actual matches
	prefixes.LoadPrefixes("../prefixes.txt")

	// Test case 1
	input := "mDHjpabcderhjkvv"
	expected := "mDHjp"
	result := prefixes.FindLongestPrefix(input)
	if result != expected {
		t.Errorf("Test failed, expected: '%s', got: '%s'", expected, result)
	}

	// Test case 2
	input = "ePX0Gqwavvjhhdtjgj"
	expected = "ePX0Gqwav"
	result = prefixes.FindLongestPrefix(input)
	if result != expected {
		t.Errorf("Test failed, expected: '%s', got: '%s'", expected, result)
	}

	// Test cases without matches
	// Test case 3
	input = "apple"
	expected = ""
	result = prefixes.FindLongestPrefix(input)
	if result != expected {
		t.Errorf("Test failed, expected: '%s', got: '%s'", expected, result)
	}

	// Test case 4 - Empty input
	input = ""
	expected = ""
	result = prefixes.FindLongestPrefix(input)
	if result != expected {
		t.Errorf("Test failed, expected: '%s', got: '%s'", expected, result)
	}

	// Test case 5
	input = "abcdefghijklmnopqrstuvwxyz"
	expected = ""
	result = prefixes.FindLongestPrefix(input)
	if result != expected {
		t.Errorf("Test failed, expected: '%s', got: '%s'", expected, result)
	}

}
