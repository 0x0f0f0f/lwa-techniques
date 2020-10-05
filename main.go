package main

import (
	"fmt"
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

	opts := randtest.TestOptions{
		Dim:        4,
		NumSymbols: 2,
		NumSamples: 1000,
		MaxWeight:  3,
		Mode:       "nat",
	}

	rand.Seed(time.Now().UnixNano())

	numSamples := 1000
	numTests := 100000
	totalVerified := 0.0

	totbpr := 0.0
	tothkc := 0.0
	totnull := 0.0

	for i := 0; i < numTests; i++ {
		fmt.Printf("running test %20d...\r", i+1)

		results := randtest.RandTest(opts)
		totalVerified += results.Verified
		totbpr += results.Bprt
		tothkc += results.Hkct
		if results.Null {
			totnull++
		}
	}

	fmt.Printf("tests completed\n")

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
