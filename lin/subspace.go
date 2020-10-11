package lin

import (
	"gonum.org/v1/gonum/mat"
)

// InSubspace returns true iff the vector b is included in U's column space
// to do so we check if the matrix U and [U|b] have the same rank
func InSubspace(u *mat.Dense, b *mat.VecDense, tol float64) bool {
	var svd mat.SVD
	svd.Factorize(u, mat.SVDNone)

	ranku := svd.Rank(tol)

	// we use the same backing data for the matrix U and [U|b]
	// we then check the rank with SVD factorization
	r, c := u.Dims()
	u = u.Grow(0, 1).(*mat.Dense)
	u.SetCol(c, b.RawVector().Data)
	svd.Factorize(u, mat.SVDNone)
	rankub := svd.Rank(tol)
	u = u.Slice(0, r, 0, c).(*mat.Dense)

	return ranku == rankub
}

// Intersect computes a basis for the intersection of many vector spaces
func Intersect(tol float64, spansets ...mat.Matrix) mat.Matrix {
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
	ker, _ := Nullspace(a, tol)
	return ker
}

// Union computes a basis for the union of many vector spaces
func Union(spansets ...*mat.Dense) *mat.Dense {
	n, j := spansets[0].Dims()

	basis := mat.DenseCopyOf(spansets[0])
	for k := 1; k < len(spansets); k++ {
		r, c := spansets[k].Dims()
		if r != n {
			panic("mat dimension mismatch")
		}
		newbasis := mat.NewDense(n, j+c, nil)
		newbasis.Augment(basis, spansets[k])
		j = j + c
		basis = newbasis
	}

	return basis
}

// Complement computes the orthogonal complement of a vector subspace of R^n
func Complement(spanset mat.Matrix, tol float64) mat.Matrix {
	n, m := spanset.Dims()
	// the orthogonal complement of {} is R^n
	if m == 0 {
		return EyeDense(n)
	}

	basis, _ := Nullspace(spanset.T(), tol)
	return basis
}
