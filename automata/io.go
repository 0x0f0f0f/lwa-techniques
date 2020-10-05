package automata

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/mat"
)

func errRead(msg string) error { return errors.New("could not read automaton:" + msg) }

// Helper function to sort and deduplicate in-place a slice of strings
func dedupStr(in []string) []string {
	sort.Strings(in)

	j := 0
	for i := 1; i < len(in); i++ {
		if in[j] == in[i] {
			continue
		}
		j++
		in[j] = in[i]
	}

	return in[:j+1]
}

// Read a positive number
func readIntPos(scanner *bufio.Scanner) (n int, err error) {
	// Read the number of rows
	if !scanner.Scan() {
		err = errRead("found eof when reading number")
		return
	}
	n, err = strconv.Atoi(scanner.Text())
	if err != nil {
		return
	}
	if n <= 0 {
		err = errRead(fmt.Sprintf("number %d must be positive", n))
	}
	return
}

// Read a slice of floats positive number
func readFloat64Slice(scanner *bufio.Scanner) ([]float64, error) {
	if !scanner.Scan() {
		return nil, errRead("found eof when reading vector data")
	}
	fields := strings.Fields(scanner.Text())
	data := make([]float64, 0)
	for _, str := range fields {
		num, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return nil, errRead("could not read float64 value")
		}
		data = append(data, num)
	}
	return data, nil
}

// Helper function to read a dense float matrix
func readDense(scanner *bufio.Scanner, rows, cols int) (*mat.Dense, error) {
	data := make([]float64, 0)
	for i := 0; i < rows; i++ {
		// Read a matrix line
		row, err := readFloat64Slice(scanner)
		if err != nil {
			return nil, err
		}
		if len(row) != cols {
			return nil, errRead(fmt.Sprintf("number of elements read in line %d does not match number of columns", i))
		}
		// Append the row to the matrix
		data = append(data, row...)
	}
	return mat.NewDense(rows, cols, data), nil
}

// Read an linear weighted automaton from a reader. If prompt is true
// then questions are asked (printed to stderr) before scanning
func ReadAutomaton(r io.Reader, prompt bool) (*Automaton, error) {
	a := &Automaton{}
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	// Read the first line containing the alphabet
	if prompt {
		fmt.Fprintf(os.Stderr, "enter the automaton alphabet:\n")
	}
	if !scanner.Scan() {
		return nil, errRead("could not read alphabet")
	}
	a.A = strings.Fields(scanner.Text())
	// Deduplicate elements in the alphabet
	a.A = dedupStr(a.A)
	if len(a.A) < 1 {
		return nil, errRead("alphabet is empty")
	}

	// Read the number of states
	if prompt {
		fmt.Fprintf(os.Stderr, "enter the number of states:\n")
	}
	numStates, err := readIntPos(scanner)
	if err != nil {
		return nil, err
	}
	a.Dim = numStates

	// Read the output vector
	if prompt {
		fmt.Fprintf(os.Stderr, "enter the output weight vector:\n")
	}
	outv, err := readFloat64Slice(scanner)
	if err != nil {
		return nil, err
	}
	if len(outv) != numStates {
		return nil, errRead("output vector length must be equal to the number of states")
	}
	a.O = mat.NewVecDense(len(outv), outv)

	fmt.Println(a.O.T().Dims())

	// Read a transition matrix for each symbol in the alphabet
	a.T = make(map[string]*mat.Dense)
	for _, symb := range a.A {
		if prompt {
			fmt.Fprintf(os.Stderr, "enter a %dx%d matrix:\n", numStates, numStates)
		}
		m, err := readDense(scanner, numStates, numStates)
		if err != nil {
			return nil, err
		}
		a.T[symb] = m
	}

	return a, nil
}

// Decorated representation of an automaton
func (a Automaton) FancyString() string {
	var buf bytes.Buffer

	// Print the alphabet
	buf.WriteString("A = ")
	buf.WriteString(strings.Join(a.A, " "))
	buf.WriteString("\n")

	// Print the output vector
	fo := mat.Formatted(a.O, mat.Prefix("      "), mat.Squeeze())
	buf.WriteString(fmt.Sprintf("o = %.5g\n\n", fo))

	// Print the matrices
	for sym, m := range a.T {
		fo := mat.Formatted(m, mat.Prefix("      "), mat.Squeeze())
		buf.WriteString(fmt.Sprintf("T_%s = %.5g\n\n", sym, fo))
	}

	return buf.String()
}

// String representation of an automaton
func (a Automaton) String() string {
	var buf bytes.Buffer

	// Print the alphabet
	buf.WriteString(strings.Join(a.A, " "))
	buf.WriteString("\n")

	// Print the output vector
	fo := mat.Formatted(a.O, mat.Squeeze(), mat.FormatMATLAB())
	buf.WriteString(fmt.Sprintf("o = %.5g\n\n", fo))

	// Print the matrices
	for sym, m := range a.T {
		fo := mat.Formatted(m, mat.Prefix("      "), mat.Squeeze())
		buf.WriteString(fmt.Sprintf("T_%s = %.5g\n\n", sym, fo))
	}

	return buf.String()
}
