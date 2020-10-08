// This file contains weighted automata data structure definition

package automata

import (
	"gonum.org/v1/gonum/mat"
)

type Automaton struct {
	// The input alphabet slice
	A []string
	// Transition matrices are maps from input strings
	// to dense real valued matrices
	T map[string]*mat.Dense
	// Output vector uses a dense real valued vector
	O *mat.VecDense
	// Number of states/dimension of vector space V in LWA
	Dim int
	// Col(LLWB) is a basis of the largest linear weighted bisimulation,
	// a binary linear relation. vRw iff (v-w) in Ker(R)
	// with LLWBperp, we denote a basis of the orthogonal
	// complement of the basis LLWB
	LLWBperp *mat.Dense
	BPRTol   float64 // tolerance value for BPR
	HKCTol   float64 // tolerance value for HKC
}
