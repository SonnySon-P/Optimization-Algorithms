package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type PSO struct {
	f func(float64) float64 // 目標函數
	dimensions int  // 變數的數量
	particles int  // 粒子數量
	iterate int  // 最大迭代次數
	xMax float64  // 解的上界
	xMin float64  // 解的下界
	wMax float64  // 最大慣性權重
	wMin float64  // 最小慣性權重
	c1 float64  // 學習因子
	c2 float64  // 學習因子
	toleranceError float64  // 容忍誤差
	position []float64  // 粒子位置
	velocity []float64  // 粒子速度
	individualBestSolution []float64  // 局部最佳解
	individualBestValue []float64  // 局部最佳值
	globalBestSolution float64  // 全域最佳解
	globalBestValue float64  // 全域最佳解的值
}

// 定義f(x)
func f(x float64) float64 {
	return x * x - 2
}

// 計算最佳解
func updatePSO(pso *PSO) {
	// PSO算法的迭代
	for i := 0; i < pso.iterate; i++ {
		// 根據當前迭代次數調整慣性權重w
		w := pso.wMax - float64(pso.iterate) / float64(pso.iterate) * (pso.wMax - pso.wMin)

		for j := 0; j < pso.particles; j++ {
			// 更新速度
			pso.velocity[j] = w * pso.velocity[j] + pso.c1 * rand.Float64() * (pso.individualBestSolution[j] - pso.position[j]) + pso.c2 * rand.Float64() * (pso.globalBestSolution - pso.position[j])

			// 更新位置
			pso.position[j] = pso.position[j] + pso.velocity[j]

			// 確保新解在合法範圍內
			if pso.position[j] < pso.xMin {
				pso.position[j] = pso.xMin
			} else if pso.xMax < pso.position[j] {
				pso.position[j] = pso.xMax
			}

			// 計算目標式結果
			objectiveValue := pso.f(pso.position[j])

			// 更新粒子的局部最佳解
			if objectiveValue < pso.individualBestValue[j] {
				pso.individualBestValue[j] = objectiveValue
				pso.individualBestSolution[j] = pso.position[j]
			}

			// 更新全域最佳解
			if objectiveValue < pso.globalBestValue {
				pso.globalBestValue = objectiveValue
				pso.globalBestSolution = pso.position[j]
			}
		}
	}
}

// 初始化粒子群最佳化參數
func initializePSO(f func(float64) float64, particles int, iterate int, xMax float64, xMin float64, wMax float64, wMin float64, c1 float64, c2 float64, toleranceError float64) *PSO {
	// 創建並初始化粒子群最佳化結構體
	pso := &PSO {
		f: f,
		particles: particles,
		iterate: iterate,
		xMax: xMax,
		xMin: xMin,
		wMax: wMax,
		wMin: wMin,
		c1: c1,
		c2: c2,
		toleranceError:  toleranceError,
		position: make([]float64, particles),
		velocity: make([]float64, particles),
		individualBestSolution: make([]float64, particles),
		individualBestValue: make([]float64, particles),
		globalBestSolution: math.MaxFloat64,
		globalBestValue: math.MaxFloat64,
	}

	// 使用當前時間戳作為隨機數種子
	rand.Seed(time.Now().UnixNano())

	// 初始化粒子的位置和速度
	for i := 0; i < particles; i++ {
		// 隨機初始化粒子的位置
		pso.position[i] = xMin + (xMax - xMin) * rand.Float64()

		// 隨機初始化粒子的速度
		pso.velocity[i] = rand.Float64()

		// 初始時，局部最佳解等於粒子當前的位置
		pso.individualBestSolution[i] = pso.position[i]

		// 設置初始的局部最佳值為目標函數值
		pso.individualBestValue[i] = pso.f(pso.position[i])
	}

	return pso
}

func main() {
	// 初始化粒子群最佳化參數
	pso := initializePSO(f, 20, 300, 10, -10, 0.9, 0.4, 2, 2, 1e-6)

	// 使用粒子群最佳化求解
	updatePSO(pso)

	// 輸出結果
	fmt.Printf("粒子群最佳化求得的解: %v\n", pso.globalBestSolution)
	fmt.Printf("理論解: %v\n", 0)
}
