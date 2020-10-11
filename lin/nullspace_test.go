// this file contains unit tests for nullspace calculation

package lin

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gonum.org/v1/gonum/mat"
)

func TestNullspace(t *testing.T) {
	a := mat.NewDense(3, 3, []float64{
		1.0, 2.0, 3.0,
		4.0, 5.0, 6.0,
		7.0, 8.0, 9.0,
	})
	ker, res := Nullspace(a, testTol)

	//fmt.Printf("%.60e\n", res)

	expres := 8.8817841970012523233890533447265625e-16
	assert.Equal(t, expres, res)

	expker := mat.NewDense(3, 1, []float64{
		-0.408248290463862406607375987732666544616222381591796875,
		0.81649658092772614548238152565318159759044647216796875,
		-0.408248290463863572341551844147033989429473876953125,
	})

	assert.True(t, mat.Equal(expker, ker))

}

func TestNullspace2(t *testing.T) {
	a := mat.NewDense(2, 3,
		[]float64{
			1, 7, 2,
			-2, 3, 1,
		})
	ker, _ := Nullspace(a, testTol)

	expker := mat.NewDense(3, 1,
		[]float64{
			0.056343616981901080420502836432206095196306705474853515625,
			-0.2817180849095055616970739720272831618785858154296875,
			0.9578414886923187765432885498739778995513916015625,
		})

	assert.True(t, mat.Equal(expker, ker))

}
