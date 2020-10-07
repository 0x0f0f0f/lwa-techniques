package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/0x0f0f0f/lwa-techniques/randtest"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// defer profile.Start(profile.MemProfile).Stop()
	rand.Seed(time.Now().UnixNano())

	aopts := &randtest.AutomatonTestOptions{
		NumStates:  4,
		NumSymbols: 2,
		NumSamples: 1000,
		MaxWeight:  2,
		Mode:       "nat",
	}

	bopts := &randtest.BatchTestOptions{
		AutOptions:  aopts,
		NumAutomata: 3000,
	}

	tols := []float64{1e-17, 1e-16, 1e-15, 1e-14, 1e-13, 1e-12, 1e-11, 1e-10, 1e-9, 1e-8, 1e-7, 1e-6}
	tolstr := []string{"1e-17", "1e-16", "1e-15", "1e-14", "1e-13", "1e-12", "1e-11", "1e-10", "1e-9", "1e-8", "1e-7", "1e-6"}

	// points on the graph
	pts := make(plotter.XYs, len(tols))

	// ================================================================================================

	for i, tol := range tols {
		aopts.BPRTol = tol
		aopts.HKCTol = tol

		res := randtest.BatchTest(bopts)
		res.Print()
		pts[i].X = float64(i)
		pts[i].Y = res.F1
	}

	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = fmt.Sprintf("Tests on %d automata, %d states, %d symbols, max |weight| = %d", bopts.NumAutomata, aopts.NumStates, aopts.NumSymbols, aopts.MaxWeight) +
		"\nVariating tolerance on both BPR and HKC"
	p.X.Label.Text = "Tolerance"
	p.Y.Label.Text = "F1 Score"
	//p.X.Scale = plot.LogScale{}
	//p.X.Tick.Marker = plot.LogTicks{}
	p.NominalX(tolstr...)
	p.X.Tick.Width = vg.Points(0.5)
	p.X.Tick.Length = vg.Points(8)
	p.X.Width = vg.Points(0.5)

	plotutil.AddLinePoints(p, pts)

	// Save the plot to a PNG file.
	if err := p.Save(6*vg.Inch, 6*vg.Inch, "f1-all-tol.png"); err != nil {
		panic(err)
	}

	// ================================================================================================

	aopts.BPRTol = 1e-13
	for i, tol := range tols {
		aopts.HKCTol = tol

		res := randtest.BatchTest(bopts)
		res.Print()
		pts[i].X = float64(i)
		pts[i].Y = res.F1
	}

	p, err = plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = fmt.Sprintf("Tests on %d automata, %d states, %d symbols, max |weight| = %d", bopts.NumAutomata, aopts.NumStates, aopts.NumSymbols, aopts.MaxWeight) +
		"\nVariating tolerance on HKC. BPR  tolerance set to 1e-13"
	p.X.Label.Text = "Tolerance"
	p.Y.Label.Text = "F1 Score"
	//p.X.Scale = plot.LogScale{}
	//p.X.Tick.Marker = plot.LogTicks{}
	p.NominalX(tolstr...)
	p.X.Tick.Width = vg.Points(0.5)
	p.X.Tick.Length = vg.Points(8)
	p.X.Width = vg.Points(0.5)

	plotutil.AddLinePoints(p, pts)

	// Save the plot to a PNG file.
	if err := p.Save(6*vg.Inch, 6*vg.Inch, "f1-hkc-tol.png"); err != nil {
		panic(err)
	}

	// ================================================================================================

	aopts.HKCTol = 1e-13
	for i, tol := range tols {
		aopts.BPRTol = tol

		res := randtest.BatchTest(bopts)
		res.Print()
		pts[i].X = float64(i)
		pts[i].Y = res.F1
	}

	p, err = plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = fmt.Sprintf("Tests on %d automata, %d states, %d symbols, max |weight| = %d", bopts.NumAutomata, aopts.NumStates, aopts.NumSymbols, aopts.MaxWeight) +
		"\nVariating tolerance on BPR. HKC tolerance set to 1e-13"
	p.X.Label.Text = "Tolerance"
	p.Y.Label.Text = "F1 Score"
	//p.X.Scale = plot.LogScale{}
	//p.X.Tick.Marker = plot.LogTicks{}
	p.NominalX(tolstr...)
	p.X.Tick.Width = vg.Points(0.5)
	p.X.Tick.Length = vg.Points(8)
	p.X.Width = vg.Points(0.5)

	plotutil.AddLinePoints(p, pts)

	// Save the plot to a PNG file.
	if err := p.Save(6*vg.Inch, 6*vg.Inch, "f1-bpr-tol.png"); err != nil {
		panic(err)
	}

}
