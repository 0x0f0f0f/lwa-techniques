package randtest

import (
	"errors"
	"math/rand"

	"github.com/0x0f0f0f/lwa-techniques/automata"
	"github.com/0x0f0f0f/lwa-techniques/lin"
	"gonum.org/v1/gonum/mat"
)

const tol = 10e-12

type TestOptions struct {
	Dim        int
	NumSymbols int
	NumSamples int
	MaxWeight  int
	Mode       string
}

type TestResults struct {
	// number of samples on which language equivalence
	// is verified by both HKC and BPR
	Verified float64
	// number of samples on which bpr equivalence is verified
	Bprt float64
	// number of samples on which hkc equivalence is verified
	Hkct float64
	// true if the LLWB computed by BPR is empty
	Null bool
}

func RandTest(o TestOptions) TestResults {
	var az automata.Automaton

	switch o.Mode {
	case "real":
		az = automata.RandAutomaton(o.NumSymbols, o.Dim, float64(o.MaxWeight), tol)

	case "nat":
		az = automata.RandNatAutomaton(o.NumSymbols, o.Dim, o.MaxWeight, tol)

	default:
		panic(errors.New("unknown mode"))
	}

	// fmt.Println(az)
	az.BackwardsPartitionRefinement()

	lin.CleanTolDense(az.LLWBperp, tol)

	samples := make([]*mat.VecDense, o.NumSamples)
	llwb := lin.Complement(az.LLWBperp, tol).(*mat.Dense)
	_, dimLLWB := llwb.Dims()

	lin.CleanTolDense(llwb, tol)

	if mat.Equal(llwb, mat.NewDense(o.Dim, 1, nil)) {
		return TestResults{
			Verified: 0,
			Bprt:     0,
			Hkct:     0,
			Null:     true,
		}
	}

	//lin.PrintMat(llwb)

	results := TestResults{Null: false}

	for i := range samples {
		samples[i] = lin.LinearCombination(llwb, lin.RandVec(dimLLWB, 100))
	}

	for i := range samples {
		j := rand.Intn(o.NumSamples)
		for j == i {
			j = rand.Intn(o.NumSamples)
		}
		resBPR := az.BPREquivalence(samples[i], samples[j])
		resHKC, _ := az.HKC(samples[i], samples[j], 30)

		if resBPR {
			results.Bprt++
		}

		if resHKC {
			results.Hkct++
		}

		if resBPR == resHKC {
			results.Verified++
		}
	}

	return results
}
