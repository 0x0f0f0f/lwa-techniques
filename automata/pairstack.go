package automata

import (
	"errors"
	"fmt"
)

// stack data structure for real valued vector pairs, used in the HKC algorithm
type PairStack struct {
	stack []Pair
}

// create a new empty pair stack
func NewPairStack() PairStack {
	return PairStack{stack: []Pair{}}
}

// returns true if the stack is empty, false otherwise
func (ps PairStack) Empty() bool {
	return len(ps.stack) == 0
}

// returns the size
func (ps PairStack) Size() int {
	return len(ps.stack)
}

// push a pair into the stack
func (ps *PairStack) Push(p Pair) {
	ps.stack = append(ps.stack, p)
}

func (ps PairStack) String() string {
	return fmt.Sprintf("%v", ps.stack)
}

// pop a pair from the stack
func (ps *PairStack) Pop() (*Pair, error) {
	siz := ps.Size()
	if siz == 0 {
		return nil, errors.New("stack is empty")
	}
	el := ps.stack[siz-1]
	ps.stack = ps.stack[:siz-1]

	return &el, nil
}
