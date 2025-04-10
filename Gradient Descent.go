package main

import (
	"fmt"
	"math"
)

type Gradient struct {
	f func(float64) float64  // 目標函數
	x float64  // 初始值
	h float64  // 一個很小的步長
	learningRate float64  // 學習率
	toleranceError float64  // 容忍誤差
	iterate int      // 最大迭代次數
}

// 定義f(x)
func f(x float64) float64 {
	return x * x - 2
}

// 計算目標函數的導數，使用中心差商法
func centralDifference(f func(float64) float64, x float64, h float64) float64 {
	return (f(x + h) - f(x - h)) / (2 * h)
}

// 計算最佳解
func updateGradient(gradient *Gradient) {
	// 使用梯度下降求解
	for i := 0; i < gradient.iterate; i++ {
		// 計算f'(x)
		dfx := centralDifference(gradient.f, gradient.x, gradient.h)

		// 計算新的x
		gradient.x = gradient.x - gradient.learningRate * dfx

		// 檢查收斂條件
		if math.Abs(dfx) < gradient.toleranceError {
			break
		}
	}
}

// 初始化Gradient會使用到的參數
func initializeGradient(f func(float64) float64, x float64, h float64, learningRate float64, toleranceError float64, iterate int) *Gradient {
	gradient := &Gradient {
		f: f,
		x: x,
		h: h,
		learningRate: learningRate,
		toleranceError: toleranceError,
		iterate: iterate,
	}

	return gradient
}

func main() {
	// 初始化梯度下降的參數
	gradient := initializeGradient(f, 1.0, 1e-6, 0.03, 1e-6, 2000)

	// 使用梯度下降求解
	updateGradient(gradient)

	// 輸出結果
	fmt.Printf("梯度下降求得的解: %v\n", gradient.x)
	fmt.Printf("理論解: %v\n", 0)
}
