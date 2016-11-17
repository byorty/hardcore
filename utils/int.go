package utils

import (
	"math/rand"
	"time"
)

func RandomInt(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
