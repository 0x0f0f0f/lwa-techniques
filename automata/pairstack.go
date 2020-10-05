// stack data structure for real valued vector pairs, used in the HKC algorithm

package automata

import (
	"errors"

	"gonum.org/v1/gonum/mat"
)

func NewPairStack() *mat.Dense {
	m := mat.NewDense(1, 1, nil)
	m.Reset()
	return m
}

// returns the size
func PairStackSize(s *mat.Dense) int {
	if s.IsEmpty() {
		return 0
	}
	_, n := s.Dims()
	return n
}

// push a pair into the stack
func PairStackPush(s *mat.Dense, p *mat.Dense) *mat.Dense {
	if s.IsEmpty() {
		return mat.DenseCopyOf(p)
	}
	s.Augment(s, p)
	return s
}

// pop a pair from the stack
func PairStackPop(s *mat.Dense) (*mat.Dense, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	m, n := s.Dims()
	if n%2 != 0 {
		return nil, errors.New("inconsisten stack: odd number of elements")
	}

	pair := s.Slice(0, m, n-2, n).(*mat.Dense)

	if n-2 > 0 {
		s = s.Slice(0, m, 0, n-2).(*mat.Dense)
	} else {
		s.Reset()
	}

	return pair, nil
}
