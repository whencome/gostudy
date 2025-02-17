package main

import (
	"math"

	"golang.org/x/image/colornames"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// create points
	n := 100
	sinDots := make(plotter.XYs, n)
	cosDots := make(plotter.XYs, n)
	for i := 0; i < n; i++ {
		x := 4 * math.Pi * float64(i) / float64(n-1)
		sinDots[i] = plotter.XY{X: x, Y: math.Sin(x)}
		cosDots[i] = plotter.XY{X: x, Y: math.Cos(x)}
	}

	// create plot
	p := plot.New()
	p.Title.Text = "正弦与余弦波形对比"
	p.X.Label.Text = "角度（弧度）"
	p.Y.Label.Text = "幅值"
	p.X.Min, p.X.Max = 0, 4*math.Pi // 设置X轴范围
	p.Y.Min, p.Y.Max = -1.2, 1.2    // 设置Y轴范围

	// 绘制正弦曲线
	sinLine, err := plotter.NewLine(sinDots)
	if err != nil {
		panic(err)
	}
	sinLine.Color = colornames.Red
	sinLine.Width = vg.Points(2)
	p.Add(sinLine)

	// 绘制余弦曲线
	cosLine, err := plotter.NewLine(cosDots)
	if err != nil {
		panic(err)
	}
	cosLine.Color = colornames.Blue
	cosLine.Width = vg.Points(2)
	cosLine.LineStyle.Dashes = []vg.Length{vg.Points(5), vg.Points(3)}
	p.Add(cosLine)

	p.Legend.Add("sin(x)", sinLine)
	p.Legend.Add("cos(x)", cosLine)
	p.Legend.Top = true

	// save plot
	p.Save(12*vg.Centimeter, 6*vg.Centimeter, "sin_cos.png")
}
