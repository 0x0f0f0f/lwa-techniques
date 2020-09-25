package automata

import (
	"gonum.org/v1/gonum/mat"
)

// represents a relation between sets of vectors of R^n.
// the relation is a congruence if it is an equivalence
// and is closed under linear combinations.
type Relation struct {
	s []*Pair         // the set of pairs in the relations
	u []*mat.VecDense // generating set for the congruence closure
}

func NewRelation() Relation {
	return Relation{}
}

// returns true if the relation contains the pair of vectors. O(n)
func (r Relation) Has(p *Pair) bool {
	for _, v := range r.s {
		if v.Eqs(p) {
			return true
		}
	}
	return false
}

// adds a pair of vectors to the relation
func (r *Relation) Add(p *Pair) {
	// if the set already contains the pair, return
	if r.Has(p) {
		return
	}
	// add the pair to the set
	r.s = append(r.s, p)
	// subtract the elements of the pair
	sub := mat.VecDenseCopyOf(p.Left)
	sub.SubVec(p.Left, p.Right)

	// add the subtraction result to the closure generating set
	subInU := false
	for _, v := range r.u {
		if mat.Equal(v, sub) {
			subInU = true
		}
	}
	if !subInU {
		r.u = append(r.u, sub)
	}
}

func (r Relation) PairIsInCongruenceClosure(p *Pair) bool {
	sub := mat.VecDenseCopyOf(p.Left)
	sub.SubVec(p.Left, p.Right)

	for _, v := range r.u {
		if mat.Equal(v, sub) {
			return true
		}
	}

	return false
}
