package lin

import (
	"log"
	"math"

	"gonum.org/v1/gonum/mat"
)

const tol = 10e-13

// computes a basis for the nullspace of a matrix a through svd decomposition.
// also returns the maximum residual
func Nullspace(a mat.Matrix) ([]*mat.VecDense, float64) {
	// compute svd decomposition
	var svd mat.SVD
	if ok := svd.Factorize(a, mat.SVDFullV); !ok {
		log.Fatal("failed to factorize A")
	}
	vt := mat.NewDense(1, 1, nil)
	vt.Reset()
	svd.VTo(vt)

	ker := []*mat.VecDense{}

	// residual
	res := 0.0

	// the (right) null space of A is the columns of vt corresponding to
	// singular values equal to zero.
	for j, σ := range svd.Values(nil) {
		if σ <= tol {
			v := mat.NewVecDense(1, nil)
			v.Reset()
			v.MulVec(a, vt.ColView(j))
			// current residual
			currRes := mat.Norm(v, math.Inf(1))
			if currRes > res {
				res = currRes
			}

			ker = append(ker, vt.ColView(j).(*mat.VecDense))
		}
	}

	return ker, res

}
