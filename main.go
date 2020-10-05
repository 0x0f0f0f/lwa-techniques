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
	fmt.Println(numSamples, "samples per automaton")
	fmt.Println(numTests, "automata")
	fmt.Println("LLWP is not empty for", notnull, "automata")
	fmt.Printf("language equivalence HKC = BPR for %d samples\n", int(totalVerified))
	percent := (totalVerified * 100) / (notnull * 2 * float64(numSamples))

	fmt.Printf("BPR equivalence verified for %d samples\n", int(totbpr))
	fmt.Printf("HKC equivalence verified for %d samples\n", int(tothkc))
	fmt.Printf("confidence is %.20g%%\n", percent)

}
