package automata

import (
	"errors"
	"fmt"

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
func (p Pair) Eqs(p1 *Pair) bool {
	// if the dimensions do not match, pairs are not equal
	if p.Dim != p1.Dim {
		return false
	}
	eql := mat.Equal(p.Left, p1.Left) && mat.Equal(p.Right, p1.Right)
	eqr := mat.Equal(p.Left, p1.Right) && mat.Equal(p.Right, p1.Left)

	return eql || eqr
}

func (p Pair) String() string {
	fl := mat.Formatted(p.Left, mat.FormatMATLAB())
	fr := mat.Formatted(p.Right, mat.FormatMATLAB())

	return fmt.Sprintf("(%.5g, %.5g)", fl, fr)
}
