package common

import "math/rand"

const LETTERS = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GenerateRandomString(rangeNum int) string {
	var randomString string

	for i := 0; i < rangeNum; i++ {
		randomString += string(LETTERS[rand.Int()%len(LETTERS)])
	}

	return randomString
}
