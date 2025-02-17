// 1-2-1 画线基础实践
package main

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func main() {
	// 初始化坐标点
	points := plotter.XYs{}
	for i := 0; i < 9; i++ {
		points = append(points, plotter.XY{X: float64(i), Y: float64(i * i)})
	}

	p := plot.New()
	// 标题
	p.Title.Text = "Squares"
	// X轴坐标
	p.X.Label.Text = "X"
	// Y轴坐标
	p.Y.Label.Text = "Y"

	// 绘制线条
	line, err := plotter.NewLine(points)
	if err != nil {
		panic(err)
	}
	// 设置线条宽度
	line.LineStyle.Width = vg.Points(2)
	// 设置线条颜色
	line.LineStyle.Color = plotutil.Color(1)
	p.Add(line)

	// 添加网格背景
	g := plotter.NewGrid()
	p.Add(g)

	p.Save(10*vg.Centimeter, 10*vg.Centimeter, "squares.png")
}
