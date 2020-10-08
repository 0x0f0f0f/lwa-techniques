// this file contains data structures relevant to tests

package randtest

import "fmt"

// Represent a sample result
const (
	TP int = iota // true positive
	TN            // true negative
	FP            // false positive
	FN            // false negative
)

type BatchTestOptions struct {
	AutOptions  *AutomatonTestOptions // initial automaton test options
	NumAutomata int                   // number of automata to generate and test
	Verbose     bool
}

func (opt BatchTestOptions) Print() {
	if opt.Verbose {
		fmt.Println("========= BATCH OPTIONS ===========")
		fmt.Println(opt.NumAutomata, "automata")
		opt.AutOptions.Print()
	}
}

// BatchResult represents the result of many weighted
// language equivalence sample tests on many automata
type BatchResult struct {
	TP        float64
	TN        float64
	FP        float64
	FN        float64
	Null      float64 // number of automata where ker(LLWB) is null
	Total     float64 // Total number of tested automata
	Accuracy  float64 // (TP+TN)/(TP+TN+FP+FN). percent of correctness
	Recall    float64
	Precision float64
	F1        float64
	opt       *BatchTestOptions
}

// Accumulate adds results of an automaton test to the results of a batch test
func (r *BatchResult) Accumulate(ar AutomatonResult) {
	r.TP += ar.TP
	r.TN += ar.TN
	r.FP += ar.FP
	r.FN += ar.FN
	if ar.Null {
		r.Null++
	}
	r.Total++
}

// ComputeStats computes relevant statistics on a batch test
func (r *BatchResult) ComputeStats() {
	T := r.TP + r.TN
	r.Accuracy = T / (T + r.FP + r.FN)
	r.Recall = r.TP / (r.TP + r.FN)
	r.Precision = r.TP / (r.TP + r.FP)
	r.F1 = 2 * ((r.Precision * r.Recall) / (r.Precision + r.Recall))
}

func (r BatchResult) Print() {
	if r.opt.Verbose {
		r.opt.Print()
		fmt.Println("========= RESULTS ===========")
		fmt.Println("LLWP is not empty for", float64(r.opt.NumAutomata)-r.Null, "automata")
		fmt.Printf("TP: %10d TN: %10d\n", int(r.TP), int(r.TN))
		fmt.Printf("FP: %10d FN: %10d\n", int(r.FP), int(r.FN))
		fmt.Printf("accuracy:  %.20g\n", r.Accuracy)
		fmt.Printf("precision: %.20g\n", r.Precision)
		fmt.Printf("recall:    %.20g\n", r.Recall)
		fmt.Printf("F1:        %.20g\n", r.F1)
	}
}

type AutomatonTestOptions struct {
	NumStates  int     // Number of states in automata
	NumSymbols int     // Number of symbols in alphabet
	NumSamples int     // Number of samples
	MaxWeight  int     // Max modulo of the weight
	Mode       string  // Either "real" or "nat"
	BPRTol     float64 // tolerance for BPR
	HKCTol     float64 // tolerance for HKC
}

func (opt AutomatonTestOptions) Print() {
	fmt.Println("weight kind:", opt.Mode)
	fmt.Println("max weight in modulo:", opt.MaxWeight)
	fmt.Println("number of samples per automata:", opt.NumSamples)
	fmt.Println("number of states:", opt.NumStates)
	fmt.Println("number of symbols:", opt.NumSymbols)
	fmt.Println("BPR tolerance:", opt.BPRTol)
	fmt.Println("HKC tolerance:", opt.HKCTol)
}

// AutomatonResult represents the result of many weighted
// language equivalence sample tests on a single automaton
type AutomatonResult struct {
	TP   float64
	TN   float64
	FP   float64
	FN   float64
	Null bool
}

// Accumulate adds results of a sample test to the results on an automaton test
func (r *AutomatonResult) Accumulate(kind int) {
	switch kind {
	case TP:
		r.TP++
	case TN:
		r.TN++
	case FP:
		r.FP++
	case FN:
		r.FN++
	}
}
