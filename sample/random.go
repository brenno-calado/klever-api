package sample

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomStringFromSet(a ...string) string {
	n := len(a)

	if n == 0 {
		return ""
	}

	return a[rand.Intn(n)]
}

func randomFloat32() float32 {
	return rand.Float32()
}

func randomCoinName() string {
	return randomStringFromSet(
		"BTC",
		"DGC",
		"KLV",
	)
}
