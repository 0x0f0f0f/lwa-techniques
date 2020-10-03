package automata

import (
	"errors"

	"gonum.org/v1/gonum/mat"
)

// checks the language equivalence of two state vectors for a given weighted automaton
func (a Automaton) HKC(v1, v2 *mat.VecDense, maxiter int) (bool, error) {
	rel := NewRelation(0)
	todo := NewPairStack()

	p, err := NewPair(v1, v2)
	if err != nil {
		return false, err
	}

	// fmt.Println("p = ", p)

	// insert (v1, v2) into the todo list
	todo.Push(*p)

	i := 0
	for !todo.Empty() {
		if i >= maxiter {
			return false, errors.New("maximum iteration exceeded")
		}
		// fmt.Println("Step", i)

		// fmt.Println("todo =", todo)

		// extract (v1', v2') from todo
		q, err := todo.Pop()
		if err != nil {
			return false, err
		}
		//fmt.Println("Extracted pair is", q)
		//fmt.Println("todo =", todo)

		if rel.PairIsInCongruenceClosure(q) {
			// fmt.Println("(v1', v2') \\in c(R)")
			continue
		}

		o1 := a.GetOutput(q.Left)
		o2 := a.GetOutput(q.Right)
		// fmt.Println("o(v1) =", o1)
		// fmt.Println("o(v2) =", o2)
		if o1 != o2 {
			// fmt.Println("o(v1) =/= o(v2)")
			return false, nil
		}

		for _, sym := range a.A {

			w1 := a.ApplyTransition(sym, q.Left)
			w2 := a.ApplyTransition(sym, q.Right)
			wp, err := NewPair(w1, w2)
			if err != nil {
				return false, err
			}

			todo.Push(*wp)
		}

		// insert (v1', v2') into R
		rel.Add(q)

		// fmt.Println("R = ", rel)
		i++
	}

	return true, nil
}
