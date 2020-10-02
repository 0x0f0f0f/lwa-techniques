// This file contains functions for reading automatas transition matrices and output vectors

package automata

import (
	"errors"
	"math/rand"

	"github.com/0x0f0f0f/lwa-techniques/lin"
	"gonum.org/v1/gonum/mat"
)

// ======================================================================

// Error helpers

func errRead(msg string) error { return errors.New("could not read automaton:" + msg) }

type Automaton struct {
	A        []string              // The input alphabet
	T        map[string]*mat.Dense // Transition matrices are maps from input strings to dense real valued matrices
	O        *mat.VecDense         // Output vector uses a dense real valued vector
	Dim      int                   // Number of states/dimension of vector space V in LWA
	LLWB     *mat.Dense            // Col(LLWB) is a basis of the largest linear weighted bisimulation, a binary linear relation. vRw iff (v-w) in Ker(R)
	LLWBperp *mat.Dense
}

func RandAutomaton(syms, states int) Automaton {
	// create the alphabet
	A := make([]string, syms)
	if syms <= 57 {
		for i := 0; i < syms; i++ {
			A[i] = string(rune(i + 65))
		}
	}

	T := map[string]*mat.Dense{}
	for _, sym := range A {
		T[sym] = lin.RandIntDense(states, 2)
		lin.PokeHoles(T[sym], rand.Intn((states*states)/2))
	}

	aut := Automaton{
		A:   A,
		T:   T,
		O:   lin.RandVec(states),
		Dim: states,
	}

	return aut
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
