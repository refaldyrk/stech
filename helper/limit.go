package helper

import (
	"math/rand"
	"time"
)

func GenerateRandomLimit() float64 {
	// Generate random value between 1,000,000 and 10,000,000
	rand.Seed(time.Now().UnixNano())
	return float64(rand.Intn(9000000) + 1000000)
}
