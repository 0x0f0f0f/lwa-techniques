package lin

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

// returns true if a vector is composed only of zero values
func IsZero(vec *mat.VecDense) bool {
	for _, el := range vec.RawVector().Data {
		if el != 0 {
			return false
		}
	}

	return true
}

// create a slice of m unitary vectors of R^n
func UnitaryVecs(n, m int) []*mat.VecDense {
	if m > n {
		panic(fmt.Errorf("cannot create more than %d unitary vectors of R^%d", n, n))
	}

	basis := make([]*mat.VecDense, m)
	for i := 0; i < m; i++ {
		// make an the i-th unitary vector
		unitary := make([]float64, n)
		unitary[i] = 1
		basis[i] = mat.NewVecDense(n, unitary)
	}

	return basis
}

// create an n*n identity matrix
func EyeDense(n int) *mat.Dense {
	units := UnitaryVecs(n, n)
	a := mat.NewDense(n, n, nil)
	for i, e := range units {
		a.SetRow(i, e.RawVector().Data)
	}
	return a
}

func PrintMat(a mat.Matrix) {
	fmt.Printf("%.90g\n", mat.Formatted(a, mat.Squeeze()))

}
