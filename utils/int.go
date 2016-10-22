package utils

import (
	"time"
	"math/rand"
)

func RandomInt(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max - min) + min
}
