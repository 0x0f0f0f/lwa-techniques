// this file contains unit tests for nullspace calculation

package lin

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gonum.org/v1/gonum/mat"
)

func TestNullspaceSVD(t *testing.T) {
	a := mat.NewDense(3, 3, []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0})
	ker, res := NullspaceSVD(a)

	//fmt.Printf("%.60e\n", res)

	expres := 8.8817841970012523233890533447265625e-16
	assert.Equal(t, expres, res)

	expker := []float64{
		-4.082482904638624066073759877326665446162223815917968750000000e-01,
		8.164965809277261454823815256531815975904464721679687500000000e-01,
		-4.082482904638635723415518441470339894294738769531250000000000e-01}
	for i, v := range expker {
		assert.Equal(t, v, ker[0].AtVec(i))
	}
}

func TestNullspaceSVD2(t *testing.T) {
	a := mat.NewDense(2, 3, []float64{1, 7, 2, -2, 3, 1})
	ker, _ := NullspaceSVD(a)

	//fmt.Printf("%.60e\n", res)

	expker := []float64{5.634361698190108e-02,
		-2.8171808490950556e-01,
		9.578414886923188e-01}

	assert.Len(t, ker, 1)

	for i, v := range expker {
		assert.Equal(t, v, ker[0].AtVec(i))
	}
}
