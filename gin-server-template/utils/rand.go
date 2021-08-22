package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano() + rand.Int63n(1000000))
}

// RandSeq generates a random string to serve as dummy data
func RandNumber(n int) string {

	letters := []rune("1234567890")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func RandomString(n int) string {

	letters := []rune("1234567890abcdefghijklmnopqrstuvwxyz")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func RandomString1(n int) string {

	letters := []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func RandomTimeString(n int) string {
	str := fmt.Sprintf("%d%s", time.Now().Unix(), RandomString1(n))
	return str
}
