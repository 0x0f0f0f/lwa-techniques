// this file contains methods for handling vector pairs, exploiting the mat.Dense
// matrix type

package automata

import (
	"errors"

	"gonum.org/v1/gonum/mat"
)

// NewPair creates a new pair of vectors
func NewPair(l, r *mat.VecDense) (*mat.Dense, error) {
	ld, _ := l.Dims()
	rd, _ := r.Dims()
	if ld != rd {
		return nil, errors.New("dimensions of vector do not match")
	}
	m := mat.NewDense(ld, 2, nil)

	for i := 0; i < ld; i++ {
		m.Set(i, 0, l.AtVec(i))
		m.Set(i, 1, r.AtVec(i))
	}
	return m, nil
}

// PairLeft returns left element of a vector pair
func PairLeft(p *mat.Dense) *mat.VecDense {
	return p.ColView(0).(*mat.VecDense)
}

// PairRight returns right element of a vector pair
func PairRight(p *mat.Dense) *mat.VecDense {
	return p.ColView(1).(*mat.VecDense)
}

// PairSub returns the subtraction of the elements of a vector pair
func PairSub(p *mat.Dense) *mat.VecDense {
	m, _ := p.Dims()
	sub := mat.NewVecDense(m, nil)
	sub.SubVec(PairLeft(p), PairRight(p))
	return sub
}

// PairCheck returns true if a matrix is a vector pair
func PairCheck(p *mat.Dense) bool {
	_, n := p.Dims()
	return n == 2
}

// PairEqs returns true if two pairs equal each other
func PairEqs(p, p1 *mat.Dense, tol float64) bool {
	m, _ := p.Dims()
	m1, _ := p1.Dims()
	// if the dimensions do not match, pairs are not equal
	if m != m1 || !PairCheck(p) || !PairCheck(p1) {
		return false
	}
	eq := true
	for i := 0; i < 2; i++ {
		eq = eq && mat.EqualApprox(p.ColView(i), p1.ColView(i), tol)
		eq = eq && mat.EqualApprox(p.ColView(i), p1.ColView(2-i), tol)
	}

	return eq
}
