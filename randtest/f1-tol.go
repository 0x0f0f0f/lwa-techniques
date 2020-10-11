// contains test for F1 score related to tolerance values

package randtest

import (
	"fmt"
	"time"

	"github.com/alitto/pond"
	"github.com/jinzhu/copier"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

// F1TolTask runs a test suite on random automata, collects F1 in relation
// to tolerance values and saves a plot in PDF and PNG format
func F1TolTask(fixedtol float64) {

	start := time.Now()

	aopts := &AutomatonTestOptions{
		NumStates:  4,
		NumSymbols: 2,
		NumSamples: 1000,
		MaxWeight:  2,
		Mode:       "nat",
	}

	bopts := &BatchTestOptions{
		AutOptions:  aopts,
		NumAutomata: 3000,
		Verbose:     false,
	}

	tols := []float64{1e-22, 1e-21, 1e-20, 1e-19, 1e-18, 1e-17, 1e-16, 1e-15, 1e-14, 1e-13, 1e-12, 1e-11, 1e-10, 1e-9, 1e-8, 1e-7, 1e-6}
	tolstr := []string{"1e-22", "1e-21", "1e-20", "1e-19", "1e-18", "1e-17", "1e-16", "1e-15", "1e-14", "1e-13", "1e-12", "1e-11", "1e-10", "1e-9", "1e-8", "1e-7", "1e-6"}

	// points on the graph
	ptsBoth := make(plotter.XYs, len(tols))
	ptsBPR := make(plotter.XYs, len(tols))
	ptsHKC := make(plotter.XYs, len(tols))

	// ================================================================================================

	pool := pond.New(10, 100)

	// tests varying both tolerances
	pool.Submit(func() {
		fmt.Println("Running test varying both tolerances")
		for i, tol := range tols {
			aopts.BPRTol = tol
			aopts.HKCTol = tol

			res := BatchTest(bopts)
			res.Print()
			ptsBoth[i].X = float64(i)
			ptsBoth[i].Y = res.F1
		}
	})

	// tests varying HKC tolerance
	pool.Submit(func() {
		fmt.Println("Running test varying HKC tolerance")
		HKCAutomataOpts := &AutomatonTestOptions{}
		copier.Copy(HKCAutomataOpts, aopts)
		HKCAutomataOpts.BPRTol = fixedtol

		var HKCBatchOpts BatchTestOptions
		copier.Copy(&HKCBatchOpts, bopts)
		HKCBatchOpts.AutOptions = HKCAutomataOpts

		for i, tol := range tols {
			HKCAutomataOpts.HKCTol = tol

			res := BatchTest(&HKCBatchOpts)
			res.Print()
			ptsHKC[i].X = float64(i)
			ptsHKC[i].Y = res.F1
		}
	})

	// tests varying BPR tolerance
	pool.Submit(func() {
		fmt.Println("Running test varying BPR tolerance")

		BPRAutomataOpts := &AutomatonTestOptions{}
		copier.Copy(BPRAutomataOpts, aopts)
		BPRAutomataOpts.HKCTol = fixedtol

		var BPRBatchOpts BatchTestOptions
		copier.Copy(&BPRBatchOpts, bopts)
		BPRBatchOpts.AutOptions = BPRAutomataOpts

		for i, tol := range tols {
			BPRAutomataOpts.BPRTol = tol

			res := BatchTest(&BPRBatchOpts)
			res.Print()
			ptsBPR[i].X = float64(i)
			ptsBPR[i].Y = res.F1
		}
	})

	pool.StopAndWait()

	dur := time.Now().Sub(start)

	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = fmt.Sprintf("Tests on %d automata, %d states, %d symbols, max |weight| = %d",
		bopts.NumAutomata,
		aopts.NumStates,
		aopts.NumSymbols,
		aopts.MaxWeight) +
		"\nTest took " + dur.String()
	p.X.Label.Text = "Tolerance"
	p.Y.Label.Text = "F1 Score"
	//p.X.Scale = plot.LogScale{}
	//p.X.Tick.Marker = plot.LogTicks{}
	p.NominalX(tolstr...)
	p.X.Tick.Width = vg.Points(0.5)
	p.X.Tick.Length = vg.Points(8)
	p.X.Width = vg.Points(0.5)

	plotutil.AddLinePoints(p,
		"Varying tolerance on both BPR and HKC", ptsBoth,
		"Varying tolerance on BPR, HKC tolerance set to "+fmt.Sprintf("%g", fixedtol), ptsBPR,
		"Varying tolerance on HKC, BPR tolerance set to "+fmt.Sprintf("%g", fixedtol), ptsHKC)

	// Save the plot to a PNG file.
	if err := p.Save(7*vg.Inch, 6*vg.Inch, fmt.Sprintf("paper/plots/f1-tol-%g.png", fixedtol)); err != nil {
		panic(err)
	}

	if err := p.Save(7*vg.Inch, 6*vg.Inch, fmt.Sprintf("paper/plots/f1-tol-%g.pdf", fixedtol)); err != nil {
		panic(err)
	}
}
