package store

import (
	"math/rand/v2"
)

const characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenCode() string {
	const n = 8
	byts := make([]byte, n)
	for i := range n {
		byts[i] = characters[rand.IntN(len(characters))]
	}
	return string(byts)
}
