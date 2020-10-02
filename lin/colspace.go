package lin

import (
	"log"

	"gonum.org/v1/gonum/mat"
)

// compute an orthonormal basis of the column space of a through
// svd decomposition. Cost is O(n^3)
func OrthonormalColumnSpaceBasis(a mat.Matrix) mat.Matrix {
	var svd mat.SVD
	if ok := svd.Factorize(a, mat.SVDFullU); !ok {
		log.Fatal("failed to factorize A")
	}
	u := mat.NewDense(1, 1, nil)
	u.Reset()
	svd.UTo(u)

	// The column space of A is spanned by the first r columns of U.
	j := 0
	for _, σ := range svd.Values(nil) {
		if σ <= tol {
			break
		}
		j++
	}

	m, _ := u.Dims()
	//fmt.Println(m, n, j)
	basis := u.Slice(0, m, 0, j)
	return basis
}