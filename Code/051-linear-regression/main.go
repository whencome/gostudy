package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
	"gonum.org/v1/gonum/mat"
)

func main() {
	// 创建数据帧
	df := dataframe.New(
		series.New([]float64{50, 60, 80, 100, 120}, series.Float, "面积"), // 如果是series.Float类型，则提供的数据必须是[]float64，使用[]float32无法解析数据
		series.New([]float64{150, 180, 240, 300, 350}, series.Float, "房价"),
	)
	fmt.Printf("origin df: %#v\n", df.Records())

	// 对数据进行分割
	n := df.Nrow()
	// 创建索引切片并打乱
	indices := make([]int, n)
	for i, _ := range indices {
		indices[i] = i
	}
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	rng.Shuffle(n, func(i, j int) {
		indices[i], indices[j] = indices[j], indices[i]
	})
	splitIndex := int(float32(n) * 0.8)
	trainDF := df.Subset(indices[:splitIndex])
	testDF := df.Subset(indices[splitIndex:])
	fmt.Printf("train df: %#v\n", trainDF.Records())
	fmt.Printf("test df: %#v\n", testDF.Records())

	// 模型训练，进行线性回归分析
	theta, err := train(trainDF)
	if err != nil {
		log.Panicf("train fail: %v", err)
	}

	beta0 := theta.At(0, 0)
	beta1 := theta.At(1, 0)
	fmt.Printf("回归方程：y = %.4f + %.4fx\n", beta0, beta1)

	// 进行模型验证
	err = validate(testDF, theta)
	if err != nil {
		log.Panicf("validate fail: %v", err)
	}
}

// extract 对数据进行提取和处理
func extract(df dataframe.DataFrame) (*mat.Dense, *mat.Dense, error) {
	n := df.Nrow()
	// 进行线性回归分析
	labels := make([]float64, n)
	xData := mat.NewDense(n, 2, nil) // 特征矩阵
	for i, r := range df.Records() {
		if i == 0 { // 第一行是标题
			continue
		}
		fx, err := strconv.ParseFloat(r[0], 32)
		if err != nil {
			return nil, nil, err
		}
		fy, err := strconv.ParseFloat(r[1], 32)
		if err != nil {
			return nil, nil, err
		}
		xData.Set(i-1, 0, 1)
		xData.Set(i-1, 1, fx)
		labels[i-1] = fy
	}
	yData := mat.NewDense(n, 1, labels) // 标签矩阵
	// 返回结果
	return xData, yData, nil
}

// train 进行模型训练
func train(df dataframe.DataFrame) (*mat.Dense, error) {
	xData, yData, err := extract(df)
	if err != nil {
		return nil, err
	}

	var XTx, XTy, theta *mat.Dense = &mat.Dense{}, &mat.Dense{}, &mat.Dense{}
	XTx.Mul(xData.T(), xData)
	XTy.Mul(xData.T(), yData)

	// 求XTx的逆
	var invXTX *mat.Dense = &mat.Dense{}
	err = invXTX.Inverse(XTx)
	if err != nil {
		return nil, err
	}

	// 计算theta
	theta.Mul(invXTX, XTy)
	return theta, nil
}

// validate 模型验证
func validate(df dataframe.DataFrame, theta *mat.Dense) error {
	xData, yData, err := extract(df)
	if err != nil {
		return err
	}

	// 预测值
	var pred *mat.Dense = &mat.Dense{}
	pred.Mul(xData, theta)

	// 计算均方差
	var residuals *mat.Dense = &mat.Dense{}
	residuals.Sub(pred, yData)
	mse := 0.0
	n, _ := xData.Caps()
	for i := 0; i < n; i++ {
		mse += residuals.At(i, 0) * residuals.At(i, 0)
	}
	mse = mse / float64(n)
	fmt.Printf("均方差(MSE)：%.4f\n", mse)

	// 计算R^2
	rSquared := 1 - mat.Norm(residuals, 2)/mat.Norm(yData, 2)
	fmt.Printf("R^2：%.4f\n", rSquared)
	return nil
}
