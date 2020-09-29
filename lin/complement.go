package lin

import (
	"gonum.org/v1/gonum/mat"
)

// compute the orthogonal complement of a vector subspace of R^n
func Complement(spanset mat.Matrix) mat.Matrix {
	n, m := spanset.Dims()
	// the orthogonal complement of {} is R^n
	if m == 0 {
		return EyeDense(n)
	}

	basis, _ := Nullspace(spanset.T())

	return basis
}
