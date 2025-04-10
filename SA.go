package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type SA struct {
	f func(float64) float64  // 目標函數
	initialTemperature float64  // 初始溫度
	finalTemperature float64  // 終止溫度
	xMax float64  // 解的上界
	xMin float64  // 解的下界
	alpha float64  // 溫度衰減系數
	iterate int  // 每個溫度下的迭代次數
	currentSolution float64  // 每次迭代的解
	currentValue float64  // 每次迭代的值
	bestSolution float64  // 最佳解
	bestValue float64  // 最佳解的值
}

// 定義f(x)
func f(x float64) float64 {
	return x * x - 2
}

// 計算最佳解
func updateSA(sa *SA) {
	// 隨機數生成器
	rand.Seed(time.Now().UnixNano())

	// 使用模擬退火演算法求解
	for temperature := sa.initialTemperature; temperature > sa.finalTemperature; temperature = temperature * sa.alpha {
		for i := 0; i < sa.iterate; i++ {
			// 生成一個解，步長範圍和溫度相關
			step := (rand.Float64() * 2 - 1) * (sa.xMax - sa.xMin) * math.Exp(-temperature / 1000)
			newSolution := sa.currentSolution + step

			// 確保新解在合法範圍內
			if newSolution > sa.xMax {
				newSolution = sa.xMax
			} else if newSolution < sa.xMin {
				newSolution = sa.xMin
			}

			// 計算新解的值
			newValue := f(newSolution)

			// 判断是否接受新解
			if newValue < sa.currentValue || rand.Float64() < math.Exp((sa.currentValue - newValue) / temperature) {
				sa.currentSolution = newSolution
				sa.currentValue = newValue
			}

			// 更新最佳解
			if sa.currentValue < sa.bestValue {
				sa.bestSolution = sa.currentSolution
				sa.bestValue = sa.currentValue
			}
		}
	}
}

// 初始化模擬退火演算法會使用到的參數
func initializeSA(f func(float64) float64, initialTemperature float64, finalTemperature float64, xMax float64, xMin float64, alpha float64, iterate int) *SA {
	sa := &SA {
		f: f,
		initialTemperature: initialTemperature,
		finalTemperature: finalTemperature,
		xMax: xMax,
		xMin: xMin,
		alpha: alpha,
		iterate: iterate,
		currentSolution: rand.Float64() * (xMax - xMin) + xMin,
		currentValue: f(rand.Float64() * (xMax - xMin) + xMin),
		bestSolution: rand.Float64() * (xMax - xMin) + xMin,
		bestValue: f(rand.Float64() * (xMax - xMin) + xMin),
	}

	return sa
}

func main() {
	// 初始化模擬退火演算法的參數
	sa := initializeSA(f, 1000.0, 1.0, 10.0, -10.0, 0.995, 1000)

	// 使用模擬退火演算法求解
	updateSA(sa)

	// 輸出結果
	fmt.Printf("模擬退火演算法求得的解: %v\n", sa.bestSolution)
	fmt.Printf("理論解: %v\n", 0)
}
