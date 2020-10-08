package main

import (
	"math/rand"
	"time"

	"github.com/0x0f0f0f/lwa-techniques/randtest"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// defer profile.Start(profile.MemProfile).Stop()
	rand.Seed(time.Now().UnixNano())

	randtest.F1TolTask()
}
