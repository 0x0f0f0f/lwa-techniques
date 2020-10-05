// this file contains definitions for the backwards algorithm for
// computing the largest linear weighted bisimulation

package automata

import (
	"log"

	"github.com/0x0f0f0f/lwa-techniques/lin"
	"gonum.org/v1/gonum/mat"
)

// compute and store a basis for the largest linear weighted bisimulation of
// the linear weighted automaton
func (a *Automaton) BackwardsPartitionRefinement() {
	// i = 0
	lastBasis := mat.NewDense(a.Dim, 1, a.O.RawVector().Data)
	currBasis := lastBasis
	// index of the column of last basis that has already been computed
	// lastIndex := 0

	for i := 1; i <= a.Dim; i++ {
		// Σ_{a \in A} T_a^t(R_i)
		for _, sym := range a.A {
			newBasis := a.ApplyTransposeTransitionBasis(sym, lastBasis)
			currBasis = lin.Union(currBasis, newBasis)
			currBasis = lin.OrthonormalColumnSpaceBasis(currBasis, a.Tol).(*mat.Dense)
		}

		lastBasis = currBasis
	}

	a.LLWBperp = currBasis
	// we could compute the orthogonal complement to find a basis of LLWB:
	// a.LLWB = lin.Complement(currBasis).(*mat.Dense)
}

func (a Automaton) BPREquivalence(v1, v2 *mat.VecDense) bool {
	if a.LLWBperp == nil {
		log.Fatalln("largest linear weighted bisimulation not computed for automaton")
		return false
	}

	sub := mat.VecDenseCopyOf(v1)
	sub.SubVec(v1, v2)

	mul := mat.VecDenseCopyOf(sub)
	mul.Reset()
	mul.MulVec(a.LLWBperp.T(), sub)

	return lin.IsZeroTol(mul, a.Tol)
}
