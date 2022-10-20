package random

import (
	"math/rand"
	"os"
	"time"
)

func GenerateStringByN(n int) string {
	rand.Seed(time.Now().Add(time.Second * 1563).UnixNano())
	letterBytes := os.Getenv("RANDOM_LETTERS")
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
