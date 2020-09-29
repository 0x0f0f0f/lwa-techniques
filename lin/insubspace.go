package lin

import (
	"gonum.org/v1/gonum/mat"
)

type MyDense struct {
	m *mat.Dense
}

func (m MyDense) T() mat.Matrix {
	return m.m.T()
}

func (m MyDense) At(i, j int) float64 {
	return m.m.At(i, j)
}

func (m MyDense) Dims() (int, int) {
	return m.m.Dims()
}

// returns true if vector b is contained in the subspace spanned by
// the columns of u

func (m MyDense) MulVecTo(dst *mat.VecDense, trans bool, x mat.Vector) {
	if trans {
		dst.MulVec(m.m.T(), x)
		return
	}
	dst.MulVec(m.m, x)

}

func InSubspace(u *mat.Dense, b *mat.VecDense) bool {

	var svd mat.SVD
	svd.Factorize(u, mat.SVDNone)

	ranku := svd.Rank(tol)

	r, c := u.Dims()
	aug := mat.NewDense(r, c+1, nil)

	aug.Augment(u, b)
	svd.Factorize(aug, mat.SVDNone)
	rankub := svd.Rank(tol)

	return ranku == rankub
}
