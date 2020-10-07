// This file contains weighted automata data structure definition
// and basic transition functions.

package automata

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

type Automaton struct {
	// The input alphabet
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

// applies the output function to a given state vector (o * v)
func (a Automaton) GetOutput(v *mat.VecDense, tol float64) float64 {
	res := mat.Dot(a.O, v)
	if math.Abs(res) < tol {
		res = 0.0
	}
	return res
}

// apply (multiply) a transition function for a given symbol s, to a vector v
func (a Automaton) ApplyTransition(s string, v *mat.VecDense) *mat.VecDense {
	res := mat.VecDenseCopyOf(v)
	res.MulVec(a.T[s], v)
	return res
}

// apply (multiply) a transpose transition function for a given symbol s, to a vector v
func (a Automaton) ApplyTransposeTransition(s string, v *mat.VecDense) *mat.VecDense {
	res := mat.VecDenseCopyOf(v)
	res.MulVec(a.T[s].T(), v)
	return res
}

// apply (multiply) a transpose transition function for a given symbol s, to a
// matrix b, which column space spans a subspace
func (a Automaton) ApplyTransposeTransitionBasis(s string, b *mat.Dense) *mat.Dense {
	res := mat.DenseCopyOf(b)
	res.Mul(a.T[s].T(), b)
	return res
}
