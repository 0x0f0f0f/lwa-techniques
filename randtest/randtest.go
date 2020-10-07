package randtest

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/0x0f0f0f/lwa-techniques/automata"
	"github.com/0x0f0f0f/lwa-techniques/lin"
	"gonum.org/v1/gonum/mat"
)

func BatchTest(opt *BatchTestOptions) BatchResult {
	batchResults := BatchResult{opt: opt}

	for i := 0; i < opt.NumAutomata; i++ {
		fmt.Printf("testing automata %20d...\r", i)

		batchResults.Accumulate(TestRandAutomaton(opt.AutOptions))
	}

	fmt.Println()
	batchResults.ComputeStats()
	return batchResults

}

func TestRandAutomaton(o *AutomatonTestOptions) AutomatonResult {
	var az automata.Automaton

	switch o.Mode {
	case "real":
		az = automata.RandAutomaton(o.NumSymbols, o.NumStates, float64(o.MaxWeight))

	case "nat":
		az = automata.RandNatAutomaton(o.NumSymbols, o.NumStates, o.MaxWeight)

	default:
		panic(errors.New("unknown mode"))
	}
	az.BPRTol = o.BPRTol
	az.HKCTol = o.HKCTol

	// fmt.Println(az)
	az.BackwardsPartitionRefinement()

	samples := make([]*mat.VecDense, o.NumSamples)
	randoms := make([]*mat.VecDense, o.NumSamples)
	llwb := lin.Complement(az.LLWBperp, o.BPRTol).(*mat.Dense)
	_, dimLLWB := llwb.Dims()

	lin.CleanTolDense(llwb, o.BPRTol)

	if mat.Equal(llwb, mat.NewDense(o.NumStates, 1, nil)) {
		return AutomatonResult{
			Null: true,
		}
	}

	//lin.PrintMat(llwb)

	autResult := AutomatonResult{Null: false}

	for i := range samples {
		samples[i] = lin.LinearCombination(llwb, lin.RandVec(dimLLWB, 100))
		randoms[i] = lin.RandVec(az.Dim, 100)
	}

	for i := range samples {
		j := rand.Intn(o.NumSamples)
		// test for vectors in span of LLWB
		for j == i {
			j = rand.Intn(o.NumSamples)
		}
		autResult.Accumulate(TestSamplePair(az, samples[i], samples[j]))
		// test for totally random vectors
		autResult.Accumulate(TestSamplePair(az, samples[i], samples[j]))

	}

	return autResult
}

func TestSamplePair(az automata.Automaton, v1, v2 *mat.VecDense) int {
	BPReq := az.BPREquivalence(v1, v2)
	HKCeq, _ := az.HKC(v1, v2)

	if BPReq && HKCeq {
		return TP // true positive
	} else if !BPReq && !HKCeq {
		return TN // true negative
	} else if !BPReq && HKCeq {
		return FP // false negative
	} else {
		return FN // false positive
	}
}
