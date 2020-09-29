package lin

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gonum.org/v1/gonum/mat"
)

func TestComplement(t *testing.T) {
	wspan := mat.NewDense(3, 3, []float64{
		1, 2, 3,
		4, 5, 6,
		7, 8, 9,
	})
	//assert.Equal()

	expected := mat.NewDense(3, 1, []float64{
		-0.408248290463863294785795687857898883521556854248046875,
		0.81649658092772614548238152565318159759044647216796875,
		-0.4082482904638629062077370690531097352504730224609375,
	})

	wperp := Complement(wspan)

	assert.True(t, mat.Equal(expected, wperp))

	//PrintMat(wperp)

}
