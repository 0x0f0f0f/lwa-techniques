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

	randtest.F1TolTask(1e-6)
	randtest.F1TolTask(1e-13)
	randtest.F1TolTask(1e-15 / 2)
	randtest.F1TolTask(1e-16)
	randtest.F1TolTask(1e-20)
}
