package main

import (
	"fmt"

	"github.com/0x0f0f0f/lwa-techniques/lin"
	"gonum.org/v1/gonum/mat"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	/* 		v1 := mat.NewVecDense(3, []float64{1.0, 2.0, 3.0})
	   	v2 := mat.NewVecDense(3, []float64{3.0, 2.0, 1.0})
	   	p1, err := aut.NewPair(v1, v2)
	   	check(err)

	   	v3 := mat.NewVecDense(3, []float64{4.0, 3.0, 5.0})
	   	v4 := mat.NewVecDense(3, []float64{6.0, 3.0, 3.0})
	   	p2, err := aut.NewPair(v3, v4)
	   	check(err)

	r := aut.NewRelation()
	r.Add(p1)

	fmt.Println(r.PairIsInCongruenceClosure(p2))
	*/

	// a, err := aut.ReadAutomaton(os.Stdin, true)
	// check(err)
	// fmt.Println(a.String())

	// v1 := mat.NewVecDense(3, []float64{0, 1, 0})
	// v2 := mat.NewVecDense(3, []float64{0, 0, 1})

	// res, err := a.HKC(v1, v2)
	// fmt.Println(res, err)

	w1 := mat.NewVecDense(3, []float64{1, 7, 2})
	w2 := mat.NewVecDense(3, []float64{-2, 3, 1})
	//assert.Equal()

	wperp := lin.Complement(3, []*mat.VecDense{w1, w2})

	fmt.Println(len(wperp))

}
