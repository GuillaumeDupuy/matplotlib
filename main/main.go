package main

import (
	"github.com/GuillaumeDupuy/matplotlib"
	"math/rand"
)

// Create a new plot
func main() {
	// p := matplotlib.NewPlot()

	// // Set the title of the plot
	// p.SetTitle("Line plot")

	// // Set the xlabel of the plot
	// p.SetXLabel("X-axis")

	// // Set the ylabel of the plot
	// p.SetYLabel("Y-axis")

	// // Plot
	// p.LinePlot([]float64{1, 2, 3, 4, 5}, []float64{1, 4, 9, 16, 25})

	// // Show the plot
	// p.Show(0)

	// p2 := matplotlib.NewPlot()

	// // Set the title of the plot
	// p2.SetTitle("Scatter plot")

	// // Set the xlabel of the plot
	// p2.SetXLabel("X-axis")

	// // Set the ylabel of the plot
	// p2.SetYLabel("Y-axis")

	// // Plot
	// p2.ScatterPlot([]float64{1, 2, 3, 4, 5}, []float64{1, 4, 9, 16, 25})

	// // Show the plot
	// p2.Show(1)

	// p3 := matplotlib.NewPlot()

	// // Set the title of the plot
	// p3.SetTitle("Bar plot")

	// // Set the xlabel of the plot
	// p3.SetXLabel("X-axis")

	// // Set the ylabel of the plot
	// p3.SetYLabel("Y-axis")

	// // Plot
	// p3.BarPlot([]float64{1, 2, 3, 4, 5}, []float64{1, 4, 9, 16, 25})

	// // Show the plot
	// p3.Show(2)

	// p4 := matplotlib.NewPlot()

	// // Set the title of the plot
	// p4.SetTitle("Histogram")

	// // Set the xlabel of the plot
	// p4.SetXLabel("X-axis")

	// // Set the ylabel of the plot
	// p4.SetYLabel("Y-axis")

	// // Plot
	// p4.HistogramPlot([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5)

	// // Show the plot
	// p4.Show(3)

	// p5 := matplotlib.NewPlot()

	// // Set the title of the plot
	// p5.SetTitle("Box plot")

	// // Set the xlabel of the plot
	// p5.SetXLabel("X-axis")

	// // Set the ylabel of the plot
	// p5.SetYLabel("Y-axis")

	// // Plot
	// p5.BoxPlot([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	// // Show the plot
	// p5.Show(4)

	// p6 := matplotlib.NewPlot()

	// // Set the title of the plot
	// p6.SetTitle("Pie")

	// // Set the xlabel of the plot
	// p6.SetXLabel("X-axis")

	// // Set the ylabel of the plot
	// p6.SetYLabel("Y-axis")

	// values := []float64{10, 20, 30, 40}
	// labels := []string{"A", "B", "C", "D"}

	// // Plot
	// p6.PiePlot(values, labels)

	// // Show the plot
	// p6.Show(5)

	// p7 := matplotlib.NewPlot()

	// // Set the title of the plot
	// p7.SetTitle("Violin")

	// // Set the xlabel of the plot
	// p7.SetXLabel("X-axis")

	// // Set the ylabel of the plot
	// p7.SetYLabel("Y-axis")

	// // Plot
	// p7.ViolinPlot([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	// // Show the plot
	// p7.Show(6)

	p8 := matplotlib.NewPlot()

	// Set the title of the plot
	p8.SetTitle("Heatmap")

	// Set the xlabel of the plot
	p8.SetXLabel("X-axis")

	// Set the ylabel of the plot
	p8.SetYLabel("Y-axis")

	rows, cols := 10, 10
	values := make([][]float64, rows)
	for i := range values {
		values[i] = make([]float64, cols)
		for j := range values[i] {
			values[i][j] = rand.Float64() // Utilisez des valeurs aléatoires pour l'exemple
		}
	}

	// TODO: Fix HeatMap
	// Plot
	// p8.HeatMap(values)

	// Show the plot
	p8.Show(7)
}
