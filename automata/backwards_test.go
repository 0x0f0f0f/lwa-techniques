package automata

import (
	"fmt"
	"testing"

	"github.com/0x0f0f0f/lwa-techniques/lin"
	"gonum.org/v1/gonum/mat"
)

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
		Dim: 3,
		O:   mat.NewVecDense(3, []float64{2, 1, 1}),
	}

	b := a.BackwardsPartitionRefinement()
	lin.PrintMat(b)

	v1 := mat.NewVecDense(3, []float64{1, 0, 0})
	v2 := mat.NewVecDense(3, []float64{0, 1.5, 0.5})

	res := a.BPREquivalence(v1, v2, b)
	fmt.Println(res)

	v1 = mat.NewVecDense(3, []float64{1, 0, 1})
	v2 = mat.NewVecDense(3, []float64{2, 4, 0.5})

	res = a.BPREquivalence(v1, v2, b)
	fmt.Println(res)
}
