package utils

import (
	"math"
	"strings"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func Encode(number int64) string {
	if number == 0 {
		return string(charset[0])
	}

	var builder strings.Builder
	base := int64(len(charset))

	for number > 0 {
		builder.WriteByte(charset[number%base])
		number /= base
	}

	// Reverse the string
	s := builder.String()
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func Decode(code string) int64 {
	var number int64
	base := int64(len(charset))

	for i, char := range code {
		index := strings.IndexRune(charset, char)
		exponent := len(code) - 1 - i
		number += int64(index) * int64(math.Pow(float64(base), float64(exponent)))
	}

	return number
}
