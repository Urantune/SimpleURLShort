package utils

import "crypto/rand"

const charRoot = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateCode(length int) string {
	b := make([]byte, length)
	_, _ = rand.Read(b)

	for i := 0; i < length; i++ {
		b[i] = charRoot[int(b[i])%len(charRoot)]
	}
	return string(b)
}
