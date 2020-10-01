// This file contains functions for reading automatas transition matrices and output vectors

package automata

import (
	"errors"

	"gonum.org/v1/gonum/mat"
)

// ======================================================================

// Error helpers

func errRead(msg string) error { return errors.New("could not read automaton:" + msg) }

type Automaton struct {
	A   []string              // The input alphabet
	T   map[string]*mat.Dense // Transition matrices are maps from input strings to dense real valued matrices
	O   *mat.VecDense         // Output vector uses a dense real valued vector
	Dim int                   // Number of states/dimension of vector space V in LWA
}

// applies the output function to a given state vector (o * v)
func (a Automaton) GetOutput(v *mat.VecDense) float64 {
	return mat.Dot(a.O, v)
}

func (a Automaton) ApplyTransition(s string, v *mat.VecDense) *mat.VecDense {
	res := mat.VecDenseCopyOf(v)
	res.MulVec(a.T[s], v)
	return res
}

func (a Automaton) ApplyTransposeTransition(s string, v *mat.VecDense) *mat.VecDense {
	res := mat.VecDenseCopyOf(v)
	res.MulVec(a.T[s].T(), v)
	return res
}

func (a Automaton) ApplyTransposeTransitionBasis(s string, b *mat.Dense) *mat.Dense {
	res := mat.DenseCopyOf(b)
	res.Mul(a.T[s].T(), b)
	return res
}
