package lin

import (
	"errors"
	"fmt"
	"math"

	"gonum.org/v1/gonum/mat"
)

type testCase struct {
	a [][]float64
	b []float64
	x []float64
}

// result from above test case turns out to be correct to this tolerance.
const ε = 1e-14

// gaussian elimination with scaled partial pivoting
func GaussPartialo(a0 [][]float64, b0 []float64) ([]float64, error) {
	m := len(b0)
	a := make([][]float64, m)
	for i, ai := range a0 {
		row := make([]float64, m+1)
		copy(row, ai)
		row[m] = b0[i]
		a[i] = row
	}
	for k := range a {
		iMax := 0
		max := -1.
		for i := k; i < m; i++ {
			row := a[i]
			// compute scale factor s = max abs in row
			s := -1.
			for j := k; j < m; j++ {
				x := math.Abs(row[j])
				if x > s {
					s = x
				}
			}
			// scale the abs used to pick the pivot.
			if abs := math.Abs(row[k]) / s; abs > max {
				iMax = i
				max = abs
			}
		}
		if a[iMax][k] == 0 {
			return nil, errors.New("singular")
		}
		a[k], a[iMax] = a[iMax], a[k]
		for i := k + 1; i < m; i++ {
			for j := k + 1; j <= m; j++ {
				a[i][j] -= a[k][j] * (a[i][k] / a[k][k])
			}
			a[i][k] = 0
		}
	}
	x := make([]float64, m)
	for i := m - 1; i >= 0; i-- {
		x[i] = a[i][m]
		for j := i + 1; j < m; j++ {
			x[i] -= a[i][j] * x[j]
		}
		x[i] /= a[i][i]
	}
	return x, nil
}

func GaussPartial(a0 mat.Matrix) (*mat.Dense, *mat.VecDense, error) {
	a := mat.DenseCopyOf(a0)
	a.Copy(a0)

	m, n := a.Dims()

	fmt.Printf("==========\n%.20g\n========\n", mat.Formatted(a, mat.Squeeze()))

	// permutation matrix
	p := mat.NewDense(m, m, nil)

	for k := 0; k < m; k++ {
		//fmt.Printf("before step %d:\n%.20g\n", k, mat.Formatted(a, mat.Squeeze()))

		iMax := 0
		max := -1.0
		for i := k; i < m; i++ {
			// compute scale factor s = max abs in row
			s := -1.0
			for j := k; j < n; j++ {
				x := math.Abs(a.At(i, j))
				if x > s {
					s = x
				}
			}
			// scale the abs used to pick the pivot
			if abs := math.Abs(a.At(i, k)) / s; abs > max {
				iMax = i
				max = abs
			}
		}
		pivot := a.At(iMax, k)
		if pivot == 0 {
			return a, nil, errors.New("singular matrix")
		}

		// make permutation matrix for row swap
		swaps := make([]int, m)
		for i := 0; i < m; i++ {
			swaps[i] = i
		}
		swaps[k] = iMax
		p.Permutation(m, swaps)

		//fmt.Printf("swap %d with %d\n", iMax, k)
		//fmt.Printf("perm %d:\n%.20g\n", k, mat.Formatted(p, mat.Squeeze()))

		a.Mul(p, a)

		for i := k + 1; i < m; i++ {
			for j := k + 1; j < n; j++ {
				new := a.At(i, j) - a.At(k, j)*(a.At(i, k)/a.At(k, k))
				a.Set(i, j, new)
			}
			a.Set(i, k, 0)
		}

		//fmt.Printf("step %d:\n%.20g\n", k, mat.Formatted(a, mat.Squeeze()))
	}
	// calculate the solution vector x
	x := make([]float64, m)
	for i := m - 1; i >= 0; i-- {
		x[i] = a.At(i, m)
		for j := i + 1; j < n; j++ {
			x[i] -= a.At(i, j) * x[j]
		}
		x[i] /= a.At(i, i)
	}

	res := mat.NewDense(m, n-1, nil)
	res.Copy(a)

	return a, mat.NewVecDense(m, x), nil
}
