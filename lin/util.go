package lin

import (
	"errors"
	"fmt"
	"math/rand"

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

// generate a normally distributed random n*n matrix
func RandDense(n int) *mat.Dense {
	data := make([]float64, n*n)
	for i := range data {
		data[i] = rand.NormFloat64()
	}
	return mat.NewDense(n, n, data)
}

// generate a normally distributed random integer n*n matrix
func RandIntDense(n, max int) *mat.Dense {
	data := make([]float64, n*n)
	for i := range data {
		data[i] = float64(rand.Intn(max))
	}
	return mat.NewDense(n, n, data)
}

// set z randomly chosen values in the matrix to 0
func PokeHoles(a *mat.Dense, z int) {
	m, n := a.Dims()
	for k := 0; k < z; k++ {
		i := rand.Intn(m)
		j := rand.Intn(n)

		a.Set(i, j, 0)
	}
}

// performs a linear combination of the columns of v and the coefficients
// in the elements of b
func LinearCombination(a *mat.Dense, b *mat.VecDense) *mat.VecDense {
	m, n := a.Dims()
	bn, _ := b.Dims()

	if n != bn {
		panic(errors.New("mat-vector dimension mismatch"))
	}

	res := mat.NewVecDense(m, nil)

	for i := 0; i < n; i++ {
		res.AddScaledVec(res, b.AtVec(i), a.ColView(i))
	}

	return res
}

// generate a normally distributed random n*1 vector
func RandVec(n int) *mat.VecDense {
	data := make([]float64, n)
	for i := range data {
		data[i] = rand.NormFloat64()
	}
	return mat.NewVecDense(n, data)
}

// generate a zero matrix of size m*n
func ZeroDense(m, n int) *mat.Dense {
	return mat.NewDense(m, n, nil)
}

// generate a zero vector of length n
func ZeroVec(n int) *mat.VecDense {
	return mat.NewVecDense(n, nil)
}

// string representation of a matrix
func StringMat(a mat.Matrix) string {
	return fmt.Sprintf("%.5g", mat.Formatted(a, mat.Squeeze()))
}

func PrintMat(a mat.Matrix) {
	fmt.Println(StringMat(a))
}

// set elements of a which are < tol to 0
func CleanTolDense(a *mat.Dense, tol float64) {
	m, n := a.Dims()
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if a.At(i, j) < tol {
				a.Set(i, j, 0)
			}
		}
	}
}
