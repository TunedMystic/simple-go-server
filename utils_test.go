package main

import (
	"regexp"
	"testing"
)

func TestRandomInt(t *testing.T) {
	length := 10
	min := 2
	max := 10
	for i := 0; i < length; i++ {
		num := randInt(min, max)
		if !((min <= num) && (num < max)) {
			t.Errorf("Value %v not between %v and %v", num, min, max)
		}
	}
}

func TestRandString(t *testing.T) {
	isAlphaNum := regexp.MustCompile(`^[A-Za-z0-9]+$`).MatchString
	length := 10
	stringLength := 32
	for i := 0; i < length; i++ {
		randomString := RandString(stringLength)
		// Check string length.
		if !(len(randomString) == stringLength) {
			t.Errorf("Random string is of length %v. Expected %v", len(randomString), stringLength)
		}
		// Check if alphanumeric.
		if !isAlphaNum(randomString) {
			t.Errorf("Random string is not alphanumeric. Got %v", randomString)
		}
	}
}
