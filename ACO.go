package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type ACO struct {
	f func(float64) float64 // 目標函數
	numberAnts int // 螞蟻的數量
	evaporation float64  // 費洛蒙蒸發率
	q float64  // 費洛蒙更新的強度
	xMax float64  // 解的上界
	xMin float64  // 解的下界
	toleranceError float64  // 容忍誤差
	iterate int  // 最大迭代次數
	position []float64  // 螞蟻當前的位置
	fitness []float64  // 螞蟻的適應度
	globalBestPosition float64  // 全域最佳解
	previousGlobalBestPosition float64  // 上一個全域最佳解
	p []float64  // 轉移機率
	globalBestFitness float64  // 全域最佳解的值
	pheromone []float64 // 費洛蒙的強度
}

// 定義f(x)
func f(x float64) float64 {
	return x * x - 2
}

// 更新費洛蒙
func updatePheromone(aco *ACO) {
	for i := 0; i < aco.numberAnts; i++ {
		// 費洛蒙蒸發
		aco.pheromone[i] = aco.pheromone[i] * (1 - aco.evaporation)

		// 根據螞蟻的適應度更新費洛蒙
		aco.pheromone[i] = aco.pheromone[i] + aco.q / (aco.fitness[i] + 0.0001)  // 是避免除以零
	}
}

// 初始化螞蟻的資訊
func initializeAnts(aco *ACO) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < aco.numberAnts; i++ {
		aco.position[i] = rand.Float64() * (aco.xMax - aco.xMin) + aco.xMin
		aco.fitness[i] = aco.f(aco.position[i])
		aco.pheromone[i] = 1.0
		aco.p[i] = 1.0
	}
}

// 計算最佳解
func updateACO(aco *ACO) {
	// 初始化螞蟻的資訊
	initializeAnts(aco)

	// 使用蟻群演算法求解
	for i := 0; i < aco.iterate; i++ {
		// 每隻螞蟻根據當前費洛蒙和位置選擇新的位置
		for j := 0; j < aco.numberAnts; j++ {
			// 螞蟻移動
			newPosition := aco.position[j] + (rand.Float64() * 2 - 1) * (aco.xMax - aco.xMin) * aco.pheromone[j]  // 根據費洛蒙計算位置
			
			// 確保新解在合法範圍內
			if newPosition > aco.xMax {
				newPosition = aco.xMax
			} else if newPosition < aco.xMin {
				newPosition = aco.xMin
			}

			// 計算適應度
			newFitness := aco.f(newPosition)

			// 更新最優解
			if newFitness < aco.fitness[j] {
				aco.fitness[j] = newFitness
				aco.position[j] = newPosition
			}

			// 更新全局最優解
			if newFitness < aco.globalBestFitness {
				aco.globalBestFitness = newFitness
				aco.previousGlobalBestPosition = aco.globalBestPosition
				aco.globalBestPosition = newPosition
			}
		}

		// 更新費洛蒙
		updatePheromone(aco)

		// 早期停止條件，如果解的誤差在容忍範圍內，則停止
		if 3 < i && math.Abs(aco.globalBestFitness - aco.previousGlobalBestPosition) < aco.toleranceError {
			break
		}
	}
}

// 初始化蟻群演算法會使用到的參數
func initializeACO(f func(float64) float64, numberAnts int, evaporation float64, q float64, pheromone float64, xMax float64, xMin float64, toleranceError float64, iterate int) *ACO {
	aco := &ACO {
		f: f,
		numberAnts: numberAnts,
		evaporation: evaporation,
		q: q,
		xMax: xMax,
		xMin: xMin,
		toleranceError: toleranceError,
		iterate: iterate,
		position: make([]float64, numberAnts),
		fitness: make([]float64, numberAnts),
		globalBestPosition: math.MaxFloat64,
		previousGlobalBestPosition: math.MaxFloat64,
		p: make([]float64, numberAnts),
		globalBestFitness: math.MaxFloat64,
		pheromone: make([]float64, numberAnts),
	}

	return aco
}

func main() {
	// 初始化蟻群演算法的參數
	aco := initializeACO(f, 50, 0.5, 100.0, 0.1, 10, -10, 1e-6, 1000)

	// 使用蟻群演算法求解
	updateACO(aco)

	// 輸出結果
	fmt.Printf("蟻群演算法求得的解: %v\n", aco.globalBestPosition)
	fmt.Printf("理論解: %v\n", 0)
}
