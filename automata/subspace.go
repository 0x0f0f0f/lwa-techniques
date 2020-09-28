// this file contains method for calculating the basis of the intersection and sum
// of linear subspaces of R^n

package automata

import (
	"github.com/0x0f0f0f/lwa-techniques/lin"
	"gonum.org/v1/gonum/mat"
)

//  the Zassenhaus algorithm is a method to calculate a basis for the intersection
// and sum of two subspaces of a vector space.
// func Zassenhaus(uspan, wspan []*mat.VecDense) []*mat.VecDense {
// m := len(uspan)
// k := len(wspan)
// create the block matrix of size ((m+k) * (2m))
//
// }

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
	ker, _ := lin.Nullspace(a)

	return ker
}
