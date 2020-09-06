package main

import (
	"fmt"
	"os"

	"github.com/0x0f0f0f/linear-weighted-automata-bisimulation/automata"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	a, err := automata.ReadAutomaton(os.Stdin, true)
	check(err)
	fmt.Println(a.String())
}
