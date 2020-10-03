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

const tol = 10e-12

func randTest(dim, numSyms, numSamples, maxweight int) (verified, bprt, hkct, null float64) {
	az := automata.RandAutomaton(numSyms, dim, maxweight, tol)

	az.BackwardsPartitionRefinement()

	//fmt.Println(az)
	lin.CleanTolDense(az.LLWBperp, tol)

	samples := make([]*mat.VecDense, numSamples)
	llwb := lin.Complement(az.LLWBperp, tol).(*mat.Dense)
	_, dimLLWB := llwb.Dims()

	lin.CleanTolDense(llwb, tol)

	if mat.Equal(llwb, mat.NewDense(dim, 1, nil)) {
		return 0, 0, 0, 1
	}

	//lin.PrintMat(llwb)

	for i := range samples {
		samples[i] = lin.LinearCombination(llwb, lin.RandVec(dimLLWB))
	}

	for i := range samples {
		j := rand.Intn(numSamples)
		for j == i {
			j = rand.Intn(numSamples)
		}
		resBPR := az.BPREquivalence(samples[i], samples[j])
		resHKC, _ := az.HKC(samples[i], samples[j], 30)

		if resBPR {
			bprt++
		}

		if resHKC {
			hkct++
		}

		if resBPR == resHKC {
			verified++
			//fmt.Println(resBPR, resHKC)
			//lin.PrintMat(samples[i])
			//lin.PrintMat(samples[j])
		}
	}

	return
}

func main() {
	//defer profile.Start(profile.MemProfile).Stop()

	rand.Seed(time.Now().UnixNano())

	numSamples := 1000
	numTests := 100
	totalVerified := 0.0

	totbpr := 0.0
	tothkc := 0.0
	totnull := 0.0

	for i := 0; i < numTests; i++ {

		currVerified, bprverified, hkcverified, null := randTest(4, 1, numSamples, 2)
		// percentile := (100.0 * currVerified) / float64(numSamples)
		totalVerified += currVerified
		totbpr += bprverified
		tothkc += hkcverified
		totnull += null
	}

	notnull := float64(numTests) - totnull

	fmt.Println()
	fmt.Println(numSamples, "samples per test")
	fmt.Println(numTests, "tests")
	fmt.Println("LLWP is not empty in", notnull, "tests")
	fmt.Println("language equivalence HKC = BPR for", totalVerified, "samples")
	percent := (totalVerified * 100) / (notnull * float64(numSamples))

	fmt.Println("BPR equivalence verified for ", totbpr, "samples")
	fmt.Println("HKC equivalence verified for ", tothkc, "samples")
	fmt.Printf("total success rate is %.20g%%\n", percent)

}
