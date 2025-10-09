package main

import (
	"reflect"
	"testing"
)

func Test_Pangram(t *testing.T) {
	result := Pangram("The quick brown fox jumps over the lazy dog")
	if result != "Pangram" {
		t.Errorf("Expected %s, but got %s", "pangram", result)
	}
	result = Pangram("Hello World")
	if result != "Not pangram" {
		t.Errorf("Expected %s, but got %s", "Not pangram", result)
	}
}

func Test_Palindrome(t *testing.T) {
	result := palindrome("level")
	if result != true {
		t.Errorf("Expected %t, but got %t", true, result)
	}
	result = palindrome("sky")
	if result != false {
		t.Errorf("Expected %t, but got %t", false, result)
	}
}

func Test_Abbreviate(t *testing.T) {
	result := Abbreviate("Automated Teller Machine")
	if result != "ATM" {
		t.Errorf("Expected %s but got %s", "ATM", result)
	}
	result = Abbreviate("Indian Institute of Management")
	if result != "IIM" {
		t.Errorf("Expected %s but got %s", "IIM", result)
	}
	result = Abbreviate("Central Processing Unit")
	if result != "CPU" {
		t.Errorf("Expected %s but got %s", "CPU", result)
	}

}

func Test_Frequency(t *testing.T) {
	result := Frequency("apple")
	expected := map[string]int{
		"a": 1,
		"e": 1,
		"l": 1,
		"p": 2,
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
	result = Frequency("banana")
	expected = map[string]int{
		"b": 1,
		"a": 3,
		"n": 2,
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

func Test_Fizzbizz(t *testing.T) {
	result := fizzbuzz(15)
	expected := []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
	result = fizzbuzz(4)
	expected = []string{"1", "2", "Fizz", "4"}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}

}
