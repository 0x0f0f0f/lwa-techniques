package lin

import (
	"gonum.org/v1/gonum/mat"
)

// compute a basis for the intersection of many vector spaces
func Intersect(spansets ...[]*mat.VecDense) []*mat.VecDense {
	// dimension of vector spaces
	lengths := make([]int, len(spansets))
	m := 0
	for i, span := range spansets {
		lengths[i] = len(span)
		if lengths[i] == 0 {
			return []*mat.VecDense{}
		}
		m += lengths[i]
	}

	// dimensions of vectors in subspaces
	n, _ := spansets[0][0].Dims()

	// create block matrix a of the column vectors in U and W
	a := mat.NewDense(n, m, nil)

	j := 0
	for _, span := range spansets {
		for _, vec := range span {
			a.SetCol(j, vec.RawVector().Data)
			j++
		}
	}

	// compute the nullspace
	ker, _ := NullspaceSVD(a)

	return ker
}
