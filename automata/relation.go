package automata

import (
	"bytes"
	"fmt"

	"github.com/0x0f0f0f/lwa-techniques/lin"
	"gonum.org/v1/gonum/mat"
)

// represents a relation between sets of vectors of R^n.
// the relation is a congruence if it is an equivalence
// and is closed under linear combinations.
type Relation struct {
	s   []*Pair         // the set of pairs in the relations
	u   []*mat.VecDense // generating set for the congruence closure
	tol float64
}

// creates a new empty relation
func NewRelation(tol float64) Relation {
	return Relation{tol: tol}
}

// returns true if the relation contains the pair of vectors. O(n)
func (r Relation) Has(p *Pair) bool {
	for _, v := range r.s {
		if v.Eqs(p, r.tol) {
			return true
		}
	}
	return false
}

// adds a pair of vectors to the relation
func (r *Relation) Add(p *Pair) {
	// if the set already contains the pair (v, v'), return
	if r.Has(p) {
		return
	}
	// add the pair (v,v') to the set
	r.s = append(r.s, p)
	// sub = v - v'
	sub := mat.VecDenseCopyOf(p.Left)
	sub.SubVec(p.Left, p.Right)

	// add the subtraction result to the closure generating set
	subInU := false
	for _, v := range r.u {
		if lin.EqVecTol(v, sub, r.tol) {
			subInU = true
		}
	}
	if !subInU {
		r.u = append(r.u, sub)
	}
}

// check if a pair of vectors is in a relation's congruence closure.
func (r Relation) PairIsInCongruenceClosure(p *Pair) bool {
	// sub = v - v'
	sub := mat.VecDenseCopyOf(p.Left)
	sub.SubVec(p.Left, p.Right)

	// (v, v') ∈ c(R) iff v - v' ∈ U_R
	for _, v := range r.u {
		if lin.EqVecTol(v, sub, r.tol) {
			return true
		}
	}

	return false
}

func (r Relation) String() string {
	b := bytes.Buffer{}
	for _, uel := range r.u {
		b.WriteString(fmt.Sprintf("%.5g,", mat.Formatted(uel, mat.FormatMATLAB())))
	}
	return fmt.Sprintf("(s = %v, u = [%s])", r.s, b.String())
}
