package random

import (
	"math/rand"
	"time"
)

func GenerateStringByN(n int) string {
	rand.Seed(time.Now().Add(time.Second * 1563).UnixNano())
	//"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const letterBytes = "3RHgQYeaJGm6SpX0oEywurfn4zVkKcLCF2j1x89MUTbN5ZDOlthqBdAWs7viIP"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
