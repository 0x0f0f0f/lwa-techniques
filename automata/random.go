package automata

import (
	"math/rand"

	"github.com/0x0f0f0f/lwa-techniques/lin"
	"gonum.org/v1/gonum/mat"
)

// generate a random automaton on natural numbers
func RandNatAutomaton(syms, states, maxweight int, tol float64) Automaton {
	// create the alphabet
	A := make([]string, syms)
	if syms <= 57 {
		for i := 0; i < syms; i++ {
			A[i] = string(rune(i + 65))
		}
	}

	T := map[string]*mat.Dense{}
	for _, sym := range A {
		T[sym] = lin.RandIntDense(states, maxweight)
		lin.PokeHoles(T[sym], rand.Intn((states*states)/2))
	}

	O := lin.RandNatVec(states, maxweight)
	for lin.IsZero(O) {
		O = lin.RandNatVec(states, maxweight)
	}

	aut := Automaton{
		A:   A,
		T:   T,
		O:   O,
		Dim: states,
		Tol: tol,
	}

	return aut
}

// generate a random automaton on real numbers
func RandAutomaton(syms, states int, maxweight float64, tol float64) Automaton {
	// create the alphabet
	A := make([]string, syms)
	if syms <= 57 {
		for i := 0; i < syms; i++ {
			A[i] = string(rune(i + 65))
		}
	}

	T := map[string]*mat.Dense{}
	for _, sym := range A {
		T[sym] = lin.RandDense(states, maxweight)
		lin.PokeHoles(T[sym], rand.Intn((states*states)/2))
	}

	O := lin.RandVec(states, maxweight)
	for lin.IsZero(O) {
		O = lin.RandVec(states, maxweight)
	}

	aut := Automaton{
		A:   A,
		T:   T,
		O:   O,
		Dim: states,
		Tol: tol,
	}

	return aut
}
