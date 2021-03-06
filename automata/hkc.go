// this file contains the implementation of the HKC procedure

package automata

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

// HKC checks the language equivalence of two state vectors for a
// given weighted automaton by building a congruence relation
func (a Automaton) HKC(v1, v2 *mat.VecDense) (bool, error) {
	rel := NewRelation(a.HKCTol, a.Dim)
	todo := NewPairStack()

	p, err := NewPair(v1, v2)
	if err != nil {
		return false, err
	}

	// insert (v1, v2) into the todo list
	todo = PairStackPush(todo, p)

	for !todo.IsEmpty() {
		// extract (v1', v2') from todo
		q, err := PairStackPop(todo)
		if err != nil {
			return false, err
		}

		if rel.PairIsInCongruenceClosure(q) {
			continue
		}

		o1 := a.GetOutput(PairLeft(q))
		o2 := a.GetOutput(PairRight(q))
		if math.Abs(o1-o2) > a.HKCTol {
			return false, nil
		}

		for _, sym := range a.A {

			w1 := a.ApplyTransition(sym, PairLeft(q))
			w2 := a.ApplyTransition(sym, PairRight(q))
			wp, err := NewPair(w1, w2)
			if err != nil {
				return false, err
			}

			PairStackPush(todo, wp)
		}

		// insert (v1', v2') into R
		rel.Add(q)
	}

	return true, nil
}
