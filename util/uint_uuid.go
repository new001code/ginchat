package util

import (
	"math/rand"
	"time"
)

func GetUUID() uint32 {
	rand.NewSource(time.Now().UnixNano())
	return rand.Uint32()
}
