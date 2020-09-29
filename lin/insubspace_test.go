package lin

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gonum.org/v1/gonum/mat"
)

func TestInSubspace(t *testing.T) {
	m := mat.NewDense(3, 1, []float64{1, 1, 0})

	v := mat.NewVecDense(3, []float64{3, 3, 0})

	nv := mat.NewVecDense(3, []float64{3, 3, 1})

	assert.True(t, InSubspace(m, v))

	assert.False(t, InSubspace(m, nv))

}
