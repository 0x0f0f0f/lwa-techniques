package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/0x0f0f0f/lwa-techniques/automata"
	"github.com/0x0f0f0f/lwa-techniques/lin"
	"gonum.org/v1/gonum/mat"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func randTest(dim, numSyms, numSamples int) bool {
	az := automata.RandAutomaton(numSyms, dim)

	az.BackwardsPartitionRefinement()

	if mat.Equal(az.LLWB, mat.NewDense(dim, 1, nil)) {
		return false
	}

	fmt.Println(az)
	lin.PrintMat(az.LLWB)
	_, dimLLWB := az.LLWB.Dims()

	samples := make([]*mat.VecDense, numSamples)
	for i := range samples {
		samples[i] = lin.LinearCombination(az.LLWB, lin.RandVec(dimLLWB))
	}

	for i := range samples {
		j := rand.Intn(numSamples)
		for j == i {
			j = rand.Intn(numSamples)
		}
		resBPR := az.BPREquivalence(samples[i], samples[j])
		resHKC, _ := az.HKC(samples[i], samples[j])

		if resBPR == resHKC {
			fmt.Println(resBPR, resHKC)
			lin.PrintMat(samples[i])
			lin.PrintMat(samples[j])
		}
	}
	return false
}

func main() {
	//defer profile.Start(profile.MemProfile).Stop()

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 100; i++ {
		randTest(3, 1, 10)
	}
}
