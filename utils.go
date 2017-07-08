package main

import "math/rand"

var alphabet = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func randInt(min, max int) int {
	return rand.Intn(max-min) + min
}

// RandString func.
func RandString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = alphabet[randInt(0, len(alphabet))]
	}
	return string(b)
}
