package util

import (
	"math/rand"
	"time"
)

func GetUUID() uint32 {
	rand.Seed(time.Now().UnixNano())
	return rand.Uint32()
}
