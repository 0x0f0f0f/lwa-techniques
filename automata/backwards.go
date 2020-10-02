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
		// fmt.Printf("==============================\nstep %d\n", i)
		// Σ_{a \in A} T_a^t(R_i)
		for _, sym := range a.A {
			//_, n := lastBasis.Dims()
			// if lastIndex == n {
			// break
			// }
			// toCompute := lastBasis.Slice(0, a.Dim, lastIndex, n).(*mat.Dense)
			newBasis := a.ApplyTransposeTransitionBasis(sym, lastBasis)
			// fmt.Printf("T_(%s)^t of \n%s \n = \n", sym, lin.StringMat(lastBasis))
			// lin.PrintMat(newBasis)
			currBasis = lin.Union(currBasis, newBasis)
			// fmt.Println("curr basis = ")
			// lin.PrintMat(currBasis)

			currBasis = lin.OrthonormalColumnSpaceBasis(currBasis).(*mat.Dense)
			// fmt.Printf("orthonormal basis of col space of b = \n%s\n", lin.StringMat(currBasis))
		}

		//if i > 1 && mat.Equal(lastBasis, currBasis) {
		//	break
		//}
		// _, lastIndex = lastBasis.Dims()
		// _, currSize := currBasis.Dims()
		// fmt.Printf("B_%d has size %d \n", i-1, lastIndex)
		// fmt.Printf("B_%d has size %d \n", i, currSize)
		lin.CleanTolDense(currBasis, 10e-12)
		lastBasis = currBasis
	}

	a.LLWBperp = currBasis
	a.LLWB = lin.Complement(currBasis).(*mat.Dense)
}

func (a Automaton) BPREquivalence(v1, v2 *mat.VecDense) bool {
	if a.LLWB == nil {
		log.Fatalln("largest linear weighted bisimulation not computed for automaton")
		return false
	}

	sub := mat.VecDenseCopyOf(v1)
	sub.SubVec(v1, v2)

	return lin.InSubspace(a.LLWB, sub)
}
