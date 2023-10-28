package tests

import (
	"matching_prefixes/prefixes"
	"testing"
)

func TestFindLongestPrefix(t *testing.T) {
	var maxGoroutines int = 1000

	// Load prefixes for testing
	err := prefixes.LoadPrefixes("../prefixes.txt", maxGoroutines)
	if err != nil {
		t.Errorf("Error loading prefixes: %s", err)
	}

	// Test case 1: Test with a matching prefix
	input := "mDHjpabcderhjkvv"
	expected := "mDHjp"
	result := prefixes.FindLongestPrefix(input)
	if result != expected {
		t.Errorf("Test 1 failed, expected: '%s', got: '%s'", expected, result)
	}

	// Test case 2: Test with another matching prefix
	input = "ePX0Gqwavvjhhdtjgj"
	expected = "ePX0Gqwav"
	result = prefixes.FindLongestPrefix(input)
	if result != expected {
		t.Errorf("Test 2 failed, expected: '%s', got: '%s'", expected, result)
	}

	// Test case 3: Test with no matching prefix
	input = "apple"
	expected = "sorry! no records found :("
	result = prefixes.FindLongestPrefix(input)
	if result != expected {
		t.Errorf("Test 3 failed, expected: '%s', got: '%s'", expected, result)
	}

	// Test case 4 - Empty input
	input = ""
	expected = "sorry! no records found :("
	result = prefixes.FindLongestPrefix(input)
	if result != expected {
		t.Errorf("Test 4 failed, expected: '%s', got: '%s'", expected, result)
	}

	// Test case 5
	input = "abcdefghijklmnopqrstuvwxyz"
	expected = "sorry! no records found :("
	result = prefixes.FindLongestPrefix(input)
	if result != expected {
		t.Errorf("Test 5 failed, expected: '%s', got: '%s'", expected, result)
	}

}
