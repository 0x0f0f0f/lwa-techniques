// this file contains unit tests for intersection of linear subspaces

package lin

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gonum.org/v1/gonum/mat"
)

func TestIntersect(t *testing.T) {

	u1 := mat.NewVecDense(5, []float64{1.0, 2.0, 3.0, -1, 2})
	u2 := mat.NewVecDense(5, []float64{2.0, 4.0, 7.0, 2, -1})
	w1 := mat.NewVecDense(5, []float64{1, 2, 0, -2, -1})
	w2 := mat.NewVecDense(5, []float64{0, 1, 1, -1, -1})
	w3 := mat.NewVecDense(5, []float64{0, 1, -3, -6, 1})

	base := Intersect([]*mat.VecDense{u1, u2}, []*mat.VecDense{w1, w2, w3})

	result := []float64{-4.472135954999579277036048097215825691819190979003906250000000e-01,
		4.472135954999578721924535784637555480003356933593750000000000e-01,
		-4.472135954999584273039658910420257598161697387695312500000000e-01,
		-4.472135954999575946366974221746204420924186706542968750000000e-01,
		4.472135954999579277036048097215825691819190979003906250000000e-01}

	for i, v := range result {
		assert.Equal(t, v, base[0].AtVec(i))
	}
}
