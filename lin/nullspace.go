package lin

import (
	"log"
	"math"

	"gonum.org/v1/gonum/mat"
)

const tol = 10e-13

// columns of the returned matrix form an orthonormal basis
// for the nullspace of matrix a, computed
// through svd decomposition. also returns the maximum residual
func Nullspace(a mat.Matrix) (mat.Matrix, float64) {
	// compute svd decomposition  O(n^3)
	var svd mat.SVD
	if ok := svd.Factorize(a, mat.SVDFullV); !ok {
		log.Fatal("failed to factorize A")
	}
	vt := mat.NewDense(1, 1, nil)
	vt.Reset()
	svd.VTo(vt)

	// residual
	res := 0.0

	// the (right) null space of A is the columns of vt corresponding to
	// singular values equal to zero.
	j := 0
	for _, σ := range svd.Values(nil) {
		if σ <= tol {
			break
		}
		j++
	}

	// compute the residuum
	for k := j; k < vt.RawMatrix().Cols; k++ {
		v := mat.NewVecDense(1, nil)
		v.Reset()
		v.MulVec(a, vt.ColView(k))
		// current residual
		currRes := mat.Norm(v, math.Inf(1))
		if currRes > res {
			res = currRes
		}
	}

	m, n := vt.Dims()
	if n == j {
		return ZeroDense(m, 1), 0
	}
	ker := vt.Slice(0, m, j, n)
	return ker, res
}
