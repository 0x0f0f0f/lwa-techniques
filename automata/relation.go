package automata

import (
	"errors"

	"github.com/0x0f0f0f/lwa-techniques/lin"
	"gonum.org/v1/gonum/mat"
)

// Relation represents a relation between sets of vectors of R^n.
// the relation is a congruence if it is an equivalence
// and is closed under linear combinations.
type Relation struct {
	s    *mat.Dense // the set of pairs in the relations
	u    *mat.Dense // generating set for the congruence closure
	size int        // how many pairs
	dim  int        // number of rows
	tol  float64
}

// NewRelation creates a new empty relation
func NewRelation(tol float64, dim int) Relation {
	s := mat.NewDense(dim, 1, nil)
	u := mat.NewDense(dim, 1, nil)
	s.Reset()
	u.Reset()
	return Relation{tol: tol, dim: dim, s: s, u: u}
}

// GetPair returns the pair in the relation at index i
func (r Relation) GetPair(i int) (*mat.Dense, error) {
	if i < 0 || i >= r.size {
		return nil, errors.New("index out of bounds")
	}
	return r.s.Slice(0, r.dim, i*2, (i*2)+2).(*mat.Dense), nil
}

// Has returns true if the relation contains the given pair of vectors.
// Computes in O(n). Could be done better by ordering.
func (r Relation) Has(p *mat.Dense) bool {
	m, _ := r.s.Dims()
	if m != r.dim && !PairCheck(p) {
		panic(errors.New("dimension mismatch"))
	}
	for i := 0; i < r.size; i++ {
		p1, err := r.GetPair(i)
		if err != nil {
			panic(err)
		}
		if PairEqs(p1, p, r.tol) {
			return true
		}
	}
	return false
}

// Add adds a pair of vectors to the relation
func (r *Relation) Add(p *mat.Dense) {
	// if the set already contains the pair (v, v'), return
	if r.Has(p) {
		return
	}
	r.dim++
	// add the pair (v,v') to the set
	if r.s.IsEmpty() {
		r.s = mat.DenseCopyOf(p)
	} else {
		r.s.Augment(r.s, p)
	}

	// add (v - v') result to the closure generating set
	sub := PairSub(p)
	subInU := false

	if r.u.IsEmpty() {
		r.u = mat.DenseCopyOf(sub)
		return
	}

	_, un := r.u.Dims()
	for j := 0; j < un; j++ {
		v := r.u.ColView(j).(*mat.VecDense)
		if mat.EqualApprox(v, sub, r.tol) {
			subInU = true
		}
	}

	if !subInU {
		r.u.Augment(r.u, sub)
	}
}

// PairIsInCongruenceClosure checks
//  if a pair of vectors is in a relation's congruence closure.
func (r Relation) PairIsInCongruenceClosure(p *mat.Dense) bool {
	// sub = v - v'
	sub := PairSub(p)

	// (v, v') ∈ c(R) iff v - v' ∈ U_R
	if r.u.IsEmpty() {
		return false
	}

	_, un := r.u.Dims()
	for j := 0; j < un; j++ {
		v := r.u.ColView(j).(*mat.VecDense)
		if mat.EqualApprox(v, sub, r.tol) {
			return true
		}
	}

	return false
}

func (r Relation) String() string {
	return lin.StringMat(r.s) + lin.StringMat(r.u)
}
