package automata

import (
	"errors"
	"fmt"

	"github.com/0x0f0f0f/lwa-techniques/lin"
	"gonum.org/v1/gonum/mat"
)

// represents an unordered pair of vectors
type Pair struct {
	Left  *mat.VecDense
	Right *mat.VecDense
	Dim   int
}

// creates a new pair of vectors
func NewPair(l, r *mat.VecDense) (*Pair, error) {
	ld, _ := l.Dims()
	rd, _ := r.Dims()
	if ld != rd {
		return nil, errors.New("dimensions of vector do not match")
	}
	return &Pair{Left: l, Right: r, Dim: ld}, nil
}

// Returns true if two pairs equal each other
func (p Pair) Eqs(p1 *Pair, tol float64) bool {
	// if the dimensions do not match, pairs are not equal
	if p.Dim != p1.Dim {
		return false
	}
	eql := lin.EqVecTol(p.Left, p1.Left, tol) && lin.EqVecTol(p.Right, p1.Right, tol)
	eqr := lin.EqVecTol(p.Left, p1.Right, tol) && lin.EqVecTol(p.Right, p1.Left, tol)

	return eql || eqr
}

func (p Pair) String() string {
	fl := mat.Formatted(p.Left, mat.FormatMATLAB())
	fr := mat.Formatted(p.Right, mat.FormatMATLAB())

	return fmt.Sprintf("(%.20g, %.20g)", fl, fr)
}
