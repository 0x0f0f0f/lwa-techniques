// this file contains definitions for the backwards algorithm for
// computing the largest linear weighted bisimulation

package automata

import (
	"fmt"

	"github.com/0x0f0f0f/lwa-techniques/lin"
	"gonum.org/v1/gonum/mat"
)

func (a Automaton) BackwardsPartitionRefinement() *mat.Dense {
	// i = 0

	fmt.Println(a.FancyString())

	lastBasis := mat.NewDense(a.Dim, 1, a.O.RawVector().Data)
	currBasis := lastBasis
	// index of the column of last basis that has already been computed
	lastIndex := 0

	for i := 1; i < a.Dim; i++ {
		fmt.Printf("==============================\nstep %d\n", i)
		// Σ_{a \in A} T
		for _, sym := range a.A {
			_, n := lastBasis.Dims()
			toCompute := lastBasis.Slice(0, a.Dim, lastIndex, n).(*mat.Dense)
			newBasis := a.ApplyTransposeTransitionBasis(sym, toCompute)
			fmt.Printf("T_(%s)^t of \n%s \n = \n", sym, lin.StringMat(toCompute))
			lin.PrintMat(newBasis)
			currBasis = lin.Union(currBasis, newBasis)
			fmt.Println("curr basis = ")
			lin.PrintMat(currBasis)
			currBasis = lin.IndependentGauss(currBasis)
			fmt.Printf("independent columns of b = \n%s\n", lin.StringMat(currBasis))
		}
		_, lastIndex = lastBasis.Dims()
		lastBasis = currBasis
	}

	return currBasis
}

func (a Automaton) BPREquivalence(v1, v2 *mat.VecDense, b *mat.Dense) bool {
	borth := lin.Complement(b)
	bo := borth.(*mat.Dense)

	sub := mat.VecDenseCopyOf(v1)
	sub.SubVec(v1, v2)

	return lin.InSubspace(bo, sub)
}
