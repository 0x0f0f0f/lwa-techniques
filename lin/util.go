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

// create an n*n identity matrix
func EyeDense(n int) *mat.Dense {
	a := mat.NewDense(n, n, nil)
	for i := 0; i < n; i++ {
		a.Set(i, i, 0)
	}
	return a
}

func StringMat(a mat.Matrix) string {
	return fmt.Sprintf("%.5g", mat.Formatted(a, mat.Squeeze()))
}

func PrintMat(a mat.Matrix) {
	fmt.Println(StringMat(a))
}
