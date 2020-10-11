package automata

import (
	"fmt"
	"testing"

	"github.com/0x0f0f0f/lwa-techniques/lin"
	"github.com/stretchr/testify/assert"
	"gonum.org/v1/gonum/mat"
)

// uses automata seen in bonchi 2012
func TestBackwardsPartitionRefinement(t *testing.T) {
	a := Automaton{
		A: []string{"a", "b"},
		T: map[string]*mat.Dense{
			"a": mat.NewDense(3, 3, []float64{
				1, 0.33333333333333333333333333333333333333, 1,
				0, 0, 0,
				0, 0, 0,
			}),
			"b": mat.NewDense(3, 3, []float64{
				1, 0, 0,
				0, 0, 3,
				0, 0.33333333333333333333333333333333333333, 0,
			}),
		},
		Dim:    3,
		O:      mat.NewVecDense(3, []float64{2, 1, 1}),
		BPRTol: 10e-14,
		HKCTol: 10e-14,
	}

	a.BackwardsPartitionRefinement()
	//lin.PrintMat(a.LLWBperp)

	v1 := mat.NewVecDense(3, []float64{1, 0, 0})
	v2 := mat.NewVecDense(3, []float64{0, 1.5, 0.5})

	// confront BPR and HKC results
	resBPR := a.BPREquivalence(v1, v2)
	resHKC, _ := a.HKC(v1, v2)
	fmt.Println(resBPR, resHKC)
	assert.Equal(t, resHKC, resBPR)

	v1 = mat.NewVecDense(3, []float64{1, 0, 1})
	v2 = mat.NewVecDense(3, []float64{2, 4, 0.5})

	resBPR = a.BPREquivalence(v1, v2)
	resHKC, _ = a.HKC(v1, v2)
	fmt.Println(resBPR, resHKC)
	assert.Equal(t, resHKC, resBPR)

}

// uses automata seen in boreale 09
func TestBackwardsPartitionRefinement2(t *testing.T) {
	a := Automaton{
		A: []string{"a"},
		T: map[string]*mat.Dense{
			"a": mat.NewDense(3, 3, []float64{
				0, 1, 1,
				0, 1, 0,
				0, 0, 1,
			}),
		},
		Dim:    3,
		O:      mat.NewVecDense(3, []float64{1, 1, 1}),
		HKCTol: 10e-14,
		BPRTol: 10e-14,
	}

	a.BackwardsPartitionRefinement()
	lin.PrintMat(a.LLWBperp)

	v1 := mat.NewVecDense(3, []float64{1, 0, 0})
	v2 := mat.NewVecDense(3, []float64{1, 1, -1})

	// confront BPR and HKC results
	resBPR := a.BPREquivalence(v1, v2)
	resHKC, _ := a.HKC(v1, v2)
	fmt.Println(resBPR, resHKC)
	assert.Equal(t, resHKC, resBPR)

	v1 = mat.NewVecDense(3, []float64{1, 0, 1})
	v2 = mat.NewVecDense(3, []float64{2, 4, 0.5})

	// confront BPR and HKC results
	resBPR = a.BPREquivalence(v1, v2)
	resHKC, _ = a.HKC(v1, v2)
	fmt.Println(resBPR, resHKC)
	assert.Equal(t, resHKC, resBPR)

}
