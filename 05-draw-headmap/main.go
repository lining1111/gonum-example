package main

import (
	"math"
	"math/rand"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/palette/moreland"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

func main() {
	points := plotter.XYs{}
	for i := 0; i < 10; i++ {
		points = append(points, plotter.XY{
			X: 100 * rand.Float64(),
			Y: 100 * rand.Float64(),
		})
	}

	scatter, err := plotter.NewScatter(points)
	if err != nil {
		panic(err)
	}
	scatter.Shape = draw.CircleGlyph{}

	heatmap := plotter.NewHeatMap(Heat(points), moreland.SmoothBlueRed().Palette(100))

	plt := plot.New()

	plt.Y.Min, plt.X.Min, plt.Y.Max, plt.X.Max = 0, 0, 100, 100

	plt.Add(heatmap, scatter)
	if err := plt.Save(5*vg.Inch, 5*vg.Inch, "05-draw-heatmap.png"); err != nil {
		panic(err)
	}
}

type Heat plotter.XYs

func (h Heat) Dims() (c, r int) { return 100, 100 }
func (h Heat) X(c int) float64  { return float64(c) }
func (h Heat) Y(r int) float64  { return float64(r) }
func (h Heat) Z(c, r int) float64 {
	var sum float64
	for _, p := range h {
		sum += math.Sqrt(math.Pow(p.X-h.X(c), 2) + math.Pow(p.Y-h.Y(r), 2))
	}
	return -sum
}
