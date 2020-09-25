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
	A []string              // The input alphabet
	T map[string]*mat.Dense // Transition matrices are maps from input symbols to dense real valued matrices
	O *mat.VecDense         // Output vector uses a dense real valued vector
}
