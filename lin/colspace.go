package lin

import (
	"log"
	"math"

	"gonum.org/v1/gonum/mat"
)

// OrthonormalColumnSpaceBasis computes an orthonormal basis of the column space of a through
// svd decomposition. Cost is O(n^3). Also returns the condition number
func OrthonormalColumnSpaceBasis(a mat.Matrix, tol float64) (mat.Matrix, float64) {
	var svd mat.SVD
	if ok := svd.Factorize(a, mat.SVDFullU); !ok {
		log.Fatal("failed to factorize A")
	}
	u := mat.NewDense(1, 1, nil)
	u.Reset()
	svd.UTo(u)
	//fmt.Println(mat.Cond(a, 2))
	//fmt.Println(mat.Cond(u, 2))

	// The column space of A is spanned by the first r columns of U.
	j := 0
	for _, σ := range svd.Values(nil) {
		if math.Abs(σ) <= tol {
			break
		}
		j++
	}

	m, _ := u.Dims()
	//fmt.Println(m, n, j)
	basis := u.Slice(0, m, 0, j)
	return basis, svd.Cond()
}
