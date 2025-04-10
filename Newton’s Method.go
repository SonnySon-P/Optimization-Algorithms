package main

import (
	"fmt"
	"math"
)

type Newton struct {
	f func(float64) float64  // 目標函數
	x float64  // 初始值
	h float64  // 一個很小的步長
	toleranceError float64  // 容忍誤差
	iterate int  // 最大迭代次數
}

// 定義f(x)
func f(x float64) float64 {
	return x * x - 2
}

// 中心差商法，計算f'(x)
func firstDerivative(f func(float64) float64, x float64, h float64) float64 {
	return (f(x + h) - f(x - h)) / (2 * h)
}

// 中心差商法，計算f''(x)
func secondDerivative(f func(float64) float64, x float64, h float64) float64 {
	return (f(x + h) - 2 * f(x) + f(x - h)) / (h * h)
}

// 計算最佳解
func updateNnewton(newton *Newton) {
	// 使用牛頓法求解
	for i := 0; i < newton.iterate; i++ {
		// 計算f(x)、f'(x)和f''(x)
		dfx := firstDerivative(f, newton.x, newton.h)
		df2x := secondDerivative(f, newton.x, newton.h)

		// 避免除以零的情況
		if math.Abs(df2x) < 1e-6 {
			break
		}

		// 計算新的x（牛頓法公式）
		xNew := newton.x - dfx / df2x

		// 檢查收斂條件
		if math.Abs(xNew - newton.x) < newton.toleranceError  {
			break
		}

		// 更新x
		newton.x = xNew
	}
}

// 初始化Newton會使用到的參數
func initializeNewton(f func(float64) float64, x float64, h float64, toleranceError float64, iterate int) *Newton {
	newton := &Newton {
		f: f,
		x: x,
		h: h,
		toleranceError: toleranceError,
		iterate: iterate,
	}

	return newton
}

func main() {
	newton := initializeNewton(f, 1.0, 1e-6, 1e-6, 100)
	
	// 使用牛頓法求解
	updateNnewton(newton)
	
	// 輸出結果
	fmt.Printf("牛頓法求得的解: %v\n", newton.x)
	fmt.Printf("理論解: %v\n", 0)
}
