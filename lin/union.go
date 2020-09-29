package lin

import "gonum.org/v1/gonum/mat"

// compute a basis for the union of many vector spaces
// TODO matrix form
func Union(spansets ...[]*mat.VecDense) []*mat.VecDense {
	basis := []*mat.VecDense{}
	for _, span := range spansets {
		basis = append(basis, span...)
	}

	return basis
}
