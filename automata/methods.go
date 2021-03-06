// this file contains transition functions methods.

package automata

import (
	"gonum.org/v1/gonum/mat"
)

// GetOutput applies the output function to a given state vector (o * v)
func (a Automaton) GetOutput(v *mat.VecDense) float64 {
	res := mat.Dot(a.O, v)
	return res
}

// ApplyTransition applies (multiplies) a transition function for a given symbol s
// to a vector v
func (a Automaton) ApplyTransition(s string, v *mat.VecDense) *mat.VecDense {
	res := mat.VecDenseCopyOf(v)
	res.MulVec(a.T[s], v)
	return res
}

// ApplyTransposeTransition applies
// (multiplies) a transpose transition function for a given symbol s, to a vector v
func (a Automaton) ApplyTransposeTransition(s string, v *mat.VecDense) *mat.VecDense {
	res := mat.VecDenseCopyOf(v)
	res.MulVec(a.T[s].T(), v)
	return res
}

// ApplyTransposeTransitionBasis applies
// (multiplies) a transpose transition function for a given symbol s, to a
// matrix b, which column space spans a subspace
func (a Automaton) ApplyTransposeTransitionBasis(s string, b *mat.Dense) *mat.Dense {
	res := mat.DenseCopyOf(b)
	res.Mul(a.T[s].T(), b)
	return res
}
