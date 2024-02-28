package matplotlib

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"

	"os/exec"
	"runtime"

	// "image/color"

	"math"
	"strconv"

	"fmt"
)

// Plot struct
type Plot struct {
	Title  string
	XLabel string
	YLabel string
	p      *plot.Plot
}

// A basic pie chart plotter.
type PieChart struct {
	Values plotter.Values
	Labels []string
}

// HeatMap struct
type matrixGrid struct {
	cols, rows int
	Matrix     [][]float64
}

// NewPlot creates a new plot
func NewPlot() *Plot {
	p := plot.New()

	return &Plot{p: p}
}

// Set the title of the plot
func (p *Plot) SetTitle(title string) {
	p.Title = title
	p.p.Title.Text = title
}

// Set the xlabel of the plot
func (p *Plot) SetXLabel(xlabel string) {
	p.XLabel = xlabel
	p.p.X.Label.Text = xlabel
}

// Set the ylabel of the plot
func (p *Plot) SetYLabel(ylabel string) {
	p.YLabel = ylabel
	p.p.Y.Label.Text = ylabel
}

// Line Plot
func (p *Plot) LinePlot(x, y []float64) {
	pts := make(plotter.XYs, len(x))
	for i := range pts {
		pts[i].X = x[i]
		pts[i].Y = y[i]
	}

	line, err := plotter.NewLine(pts)
	if err != nil {
		panic(err)
	}

	p.p.Add(line)
}

// Scatter Plot
func (p *Plot) ScatterPlot(x, y []float64) {
	pts := make(plotter.XYs, len(x))
	for i := range pts {
		pts[i].X = x[i]
		pts[i].Y = y[i]
	}

	s, err := plotter.NewScatter(pts)
	if err != nil {
		panic(err)
	}

	p.p.Add(s)
}

// Bar Plot
func (p *Plot) BarPlot(x, y []float64) {
	pts := make(plotter.Values, len(x))
	for i := range pts {
		pts[i] = y[i]
	}

	bars, err := plotter.NewBarChart(pts, vg.Points(50))
	if err != nil {
		panic(err)
	}

	bars.Horizontal = false
	p.p.Add(bars)
}

// Violin Plot
func (p *Plot) ViolinPlot(values []float64) {
	w := vg.Points(20)
	box, err := plotter.NewBoxPlot(w, 0, plotter.Values(values))
	if err != nil {
		panic(err)
	}
	p.p.Add(box)

	leftDensity := make(plotter.XYs, len(values))
	rightDensity := make(plotter.XYs, len(values))

	leftLine, err := plotter.NewLine(leftDensity)
	if err != nil {
		panic(err)
	}
	rightLine, err := plotter.NewLine(rightDensity)
	if err != nil {
		panic(err)
	}

	p.p.Add(leftLine)
	p.p.Add(rightLine)
}

// NewPieChart creates a new pie chart plotter.
func NewPieChart(values []float64, labels []string) (*PieChart, error) {
	if len(values) != len(labels) {
		return nil, fmt.Errorf("number of values and labels must be equal")
	}

	return &PieChart{
		Values: plotter.Values(values),
		Labels: labels,
	}, nil
}

// Plot draws the pie chart to the given draw.Canvas
func (p *PieChart) Plot(c draw.Canvas, plt *plot.Plot) {
	// Calcul du rayon en fonction de la taille du canvas
	radius := vg.Length(math.Min(float64(c.Max.X-c.Min.X), float64(c.Max.Y-c.Min.Y)) / 2)

	// Calcul du centre du canvas
	centerX, centerY := vg.Length(float64(c.Min.X+c.Max.X)/2), vg.Length(float64(c.Min.Y+c.Max.Y)/2)

	total := 0.0
	for _, v := range p.Values {
		total += v
	}

	startAngle := 0.0
	for i, v := range p.Values {
		percentage := v / total
		sweep := percentage * 2 * math.Pi
		endAngle := startAngle + sweep

		path := c.ClipPolygonXY([]vg.Point{
			{X: centerX, Y: centerY},
		})

		for a := startAngle; a <= endAngle; a += 0.01 {
			x := centerX + vg.Length(float64(radius)*math.Cos(a))
			y := centerY + vg.Length(float64(radius)*math.Sin(a))
			path = append(path, vg.Point{X: x, Y: y})
		}

		path = append(path, vg.Point{X: centerX, Y: centerY})

		fillColor := plotutil.Color(i)
		c.SetColor(fillColor)
		c.FillPolygon(fillColor, path)

		// Calcul de la position du label en prenant en compte le centre du canvas
		midAngle := startAngle + (sweep / 2)
		labelX := centerX + vg.Length(0.7*float64(radius)*math.Cos(midAngle))
		labelY := centerY + vg.Length(0.7*float64(radius)*math.Sin(midAngle))
		c.FillText(plt.Title.TextStyle, vg.Point{X: labelX, Y: labelY}, p.Labels[i])

		startAngle = endAngle
	}
}

// Pie Plot
func (p *Plot) PiePlot(values []float64, labels []string) {
	pie, err := NewPieChart(values, labels)
	if err != nil {
		panic(err)
	}

	p.p.Add(pie)
}

// Box Plot
func (p *Plot) BoxPlot(values []float64) {
	vals := make(plotter.Values, len(values))

	copy(vals, values)

	box, err := plotter.NewBoxPlot(vg.Length(10), 0, vals)
	if err != nil {
		panic(err)
	}

	p.p.Add(box)
}

// Histogram Plot
func (p *Plot) HistogramPlot(values []float64, bins int) {
	hist, err := plotter.NewHist(plotter.Values(values), bins)
	if err != nil {
		panic(err)
	}

	p.p.Add(hist)
}

// Dims retourne les dimensions de la grille.
func (m matrixGrid) Dims() (c, r int) {
	return m.cols, m.rows
}

// Z retourne la valeur à la position (x, y).
func (m matrixGrid) Z(c, r int) float64 {
    if r >= 0 && r < len(m.Matrix) && c >= 0 && c < len(m.Matrix[r]) {
        return m.Matrix[r][c]
    }
    return 0
}

// X retourne la coordonnée x correspondant à l'indice de colonne c.
func (m matrixGrid) X(c int) float64 {
	return float64(c)
}

// Y retourne la coordonnée y correspondant à l'indice de ligne r.
func (m matrixGrid) Y(r int) float64 {
	return float64(r)
}

// TODO: Fix this
// HeatMap crée et ajoute une HeatMap à partir de valeurs bidimensionnelles.
// func (p *Plot) HeatMap(values [][]float64) {
// 	rows := len(values)
// 	cols := 0
// 	if rows > 0 {
// 		cols = len(values[0])
// 	}

// 	heatMap := plotter.NewHeatMap(matrixGrid{cols, rows, values}, nil)
// 	p.p.Add(heatMap)
// }

// Show the plot
func (p *Plot) Show(number int) {
	fileName := "plot" + strconv.Itoa(number) + ".png"

	if err := p.p.Save(4*vg.Inch, 4*vg.Inch, fileName); err != nil {
		panic(err)
	}

	switch runtime.GOOS {
	case "linux":
		err := exec.Command("xdg-open", fileName).Start()
		if err != nil {
			fmt.Println("Erreur lors de l'ouverture de l'image :", err)
		}
	case "windows":
		err := exec.Command("rundll32", "url.dll,FileProtocolHandler", fileName).Start()
		if err != nil {
			fmt.Println("Erreur lors de l'ouverture de l'image :", err)
		}
	case "darwin":
		err := exec.Command("open", fileName).Start()
		if err != nil {
			fmt.Println("Erreur lors de l'ouverture de l'image :", err)
		}
	default:
		fmt.Println("Plateforme non supportée pour afficher l'image automatiquement")
		fmt.Println("Vous pouvez ouvrir le fichier manuellement :", fileName)
	}
}
