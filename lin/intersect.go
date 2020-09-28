package lin

import "gonum.org/v1/gonum/mat"

// compute the intersection of two vector spaces
func Intersect(uspan, wspan []*mat.VecDense) []*mat.VecDense {
	// dimension of vector spaces
	m := len(uspan)
	k := len(wspan)

	if m == 0 || k == 0 {
		return []*mat.VecDense{}
	}

	// dimensions of vectors in U and W
	n, _ := uspan[0].Dims()

	// create block matrix a of the column vectors in U and W
	a := mat.NewDense(n, m+k, nil)
	var j int
	for j = 0; j < m; j++ {
		a.SetCol(j, uspan[j].RawVector().Data)
	}
	for j := 0; j < k; j++ {
		a.SetCol(m+j, wspan[j].RawVector().Data)
	}

	// compute the nullspace
	ker, _ := Nullspace(a)

	return ker
}
