// 1-2-5 多组数据的应用
package main

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// 初始化坐标点
	points1 := plotter.XYs{}
	points2 := plotter.XYs{}
	for i := 0; i < 9; i++ {
		points1 = append(points1, plotter.XY{X: float64(i), Y: float64(i * i)})
		points2 = append(points2, plotter.XY{X: float64(i), Y: float64(i * i * i)})
	}

	// 创建画布
	p := plot.New()
	p.Title.Text = "Squares and Cubes"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"
	p.Add(plotter.NewGrid())

	// 画线
	line1, err := plotter.NewLine(points1)
	if err != nil {
		panic(err)
	}
	p.Add(line1)

	line2, err := plotter.NewLine(points2)
	if err != nil {
		panic(err)
	}
	p.Add(line2)

	p.Save(10*vg.Centimeter, 10*vg.Centimeter, "squares&cubs.png")
}
