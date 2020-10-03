package lin

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

// find which column vectors in m are linearly independent and
// return them
func Independent(a *mat.Dense, tol float64) *mat.Dense {
	var qr mat.QR
	qr.Factorize(a)
	m, n := a.Dims()
	r := mat.NewDense(m, m, nil)
	r.Reset()
	qr.RTo(r)

	rank := 0
	cols := []int{}
	for k := 0; k < n; k++ {
		if math.Abs(r.At(k, k)) > tol {
			rank++
			cols = append(cols, k)
		}
	}

	res := mat.NewDense(m, rank, nil)

	j := 0
	for _, k := range cols {
		for i := 0; i < m; i++ {
			res.Set(i, j, a.At(i, k))
		}
		j++
	}

	return res
}

// find which column vectors in m are linearly independent and
// return them
func IndependentGauss(a *mat.Dense, tol float64) *mat.Dense {
	m, n := a.Dims()
	ab := a.Grow(0, 1)

	red, _, err := GaussPartial(ab)
	if err != nil {
		panic(err)
	}

	red = red.Slice(0, m, 0, n).(*mat.Dense)

	rank := 0
	cols := []int{}
	for k := 0; k < n; k++ {
		if k < m {
			if math.Abs(red.At(k, k)) > tol {
				rank++
				cols = append(cols, k)
			}
		}
	}

	res := mat.NewDense(m, rank, nil)

	j := 0
	for _, k := range cols {
		for i := 0; i < m; i++ {
			res.Set(i, j, a.At(i, k))
		}
		j++
	}

	return res
}
