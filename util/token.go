package util

import (
	"math/rand"
	"time"
)

func NewNumToken() int {
	s := rand.NewSource(time.Now().UnixNano())
	a := rand.New(s)
	return a.Intn(9999999)
}

