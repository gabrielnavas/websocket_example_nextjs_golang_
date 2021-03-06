package utils

import (
	"math/rand"
	"time"
)

func GetRandom() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

func GetRandomInt() int {
	return GetRandom().Int()
}
