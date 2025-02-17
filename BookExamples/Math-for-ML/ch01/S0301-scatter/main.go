package main

import (
	"golang.org/x/image/colornames"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// 绘制点
	points := plotter.XYs{}
	for i := 0; i < 100; i++ {
		points = append(points, plotter.XY{X: float64(i), Y: float64(i * i)})
	}

	// 创建画布
	p := plot.New()
	p.Title.Text = "Scatters"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	// 绘制散列点
	scatter, err := plotter.NewScatter(points)
	if err != nil {
		panic(err)
	}
	scatter.GlyphStyle.Radius = vg.Points(0.5)
	scatter.Color = colornames.Blue
	p.Add(scatter)

	// 保存图片
	p.Save(10*vg.Centimeter, 10*vg.Centimeter, "scatters.png")
}
