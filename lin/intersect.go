package lin

import (
	"gonum.org/v1/gonum/mat"
)

// compute a basis for the intersection of many vector spaces
func Intersect(spansets ...mat.Matrix) mat.Matrix {
	// dimension of vector spaces

	a := mat.DenseCopyOf(spansets[0])
	for i, span := range spansets {
		if i == 0 {
			continue
		}

		ar, ac := a.Dims()
		_, sc := span.Dims()

		m := mat.NewDense(ar, ac+sc, nil)

		m.Augment(a, span)

		a = m
	}

	// compute the nullspace
	ker, _ := Nullspace(a)

	return ker
}
