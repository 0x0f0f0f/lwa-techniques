package lin

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gonum.org/v1/gonum/mat"
)

const testTol = 10e-14

func TestInSubspace(t *testing.T) {
	m := mat.NewDense(3, 1, []float64{1, 1, 0})

	om := mat.DenseCopyOf(m)

	v := mat.NewVecDense(3, []float64{3, 3, 0})

	nv := mat.NewVecDense(3, []float64{3, 3, 1})

	assert.True(t, InSubspace(m, v, testTol))
	assert.False(t, InSubspace(m, nv, testTol))

	// matrix is not modified
	assert.True(t, mat.Equal(m, om))
}

func TestIntersect(t *testing.T) {

	u := mat.NewDense(5, 2, []float64{
		1, 2,
		2, 4,
		3, 7,
		-1, 2,
		2, -1,
	})
	w := mat.NewDense(5, 3, []float64{
		1, 0, 0,
		2, 1, 1,
		0, 1, -3,
		-2, -1, -6,
		-1, -1, 1,
	})

	base := Intersect(testTol, u, w)

	expected := mat.NewDense(5, 1, []float64{
		-0.447213595499957927703604809721582569181919097900390625,
		0.447213595499957872192453578463755548000335693359375,
		-0.44721359549995842730396589104202575981616973876953125,
		-0.447213595499957594636697422174620442092418670654296875,
		0.447213595499957927703604809721582569181919097900390625,
	})

	assert.True(t, mat.Equal(expected, base))
}

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

	wperp := Complement(wspan, testTol)

	assert.True(t, mat.Equal(expected, wperp))

	//PrintMat(wperp)
}

func TestUnion(t *testing.T) {
	wspan := mat.NewDense(3, 3, []float64{
		1, 2, 3,
		4, 5, 6,
		7, 8, 9,
	})

	uspan := mat.NewDense(3, 3, []float64{
		8, 3, 12,
		2, 5, 2,
		1, 4, 7,
	})
	//assert.Equal()

	expected := mat.NewDense(3, 6, []float64{
		1, 2, 3, 8, 3, 12,
		4, 5, 6, 2, 5, 2,
		7, 8, 9, 1, 4, 7,
	})

	un := Union(wspan, uspan)

	assert.True(t, mat.Equal(expected, un))

	//PrintMat(wperp)
}
