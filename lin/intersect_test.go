// this file contains unit tests for intersection of linear subspaces

package lin

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gonum.org/v1/gonum/mat"
)

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

	base := Intersect(u, w)

	expected := mat.NewDense(5, 1, []float64{
		-0.447213595499957927703604809721582569181919097900390625,
		0.447213595499957872192453578463755548000335693359375,
		-0.44721359549995842730396589104202575981616973876953125,
		-0.447213595499957594636697422174620442092418670654296875,
		0.447213595499957927703604809721582569181919097900390625,
	})

	assert.True(t, mat.Equal(expected, base))
}
