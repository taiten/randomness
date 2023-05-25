package main

import (
	"github.com/Trisia/randomness"
	"log"
	"os"
	"path"
)

// Header_1E6 1_000_000 比特样本检测
const Header_1E6 = "源数据," +
	"[ 1] 单比特频数检测," +
	"[ 2] 块内频数检测 m=10000," +
	"[ 3] 扑克检测 m=4," +
	"[ 3] 扑克检测 m=8," +
	"[ 4] 重叠子序列检测 m=3 P1,重叠子序列检测 m=2 P2," +
	"[ 4] 重叠子序列检测 m=5 P1,重叠子序列检测 m=5 P2," +
	"[ 5] 游程总数检测," +
	"[ 6] 游程分布检测," +
	"[ 7] 块内最大游程检测 m=10000," +
	"[ 8] 二元推导检测 k=3," +
	"[ 8] 二元推导检测 k=7," +
	"[ 9] 自相关检测 d=1," +
	"[ 9] 自相关检测 d=2," +
	"[ 9] 自相关检测 d=8," +
	"[ 9] 自相关检测 d=16," +
	"[10] 矩阵秩检测," +
	"[11] 累加和检测," +
	"[12] 近似熵检测 m=2," +
	"[12] 近似熵检测 m=5," +
	"[13] 线性复杂度检测 m=500," +
	"[13] 线性复杂度检测 m=1000," +
	"[14] Maurer通用统计检测 L=7 Q=1280," +
	"[15] 离散傅里叶检测\n"

// 数据规模为 1 000 000 个比特的随机数列检测工作器
func worker_1E6(jobs <-chan string, out chan<- *R) {
	for filename := range jobs {
		buf, _ := os.ReadFile(filename)
		bits := randomness.B2bitArr(buf)
		buf = nil
		arr := make([]float64, 0, 64)

		log.Printf("[%s] 检测开始...\n", filename)

		// [1] 单比特频数检测
		p, _ := randomness.MonoBitFrequencyTest(bits)
		arr = append(arr, p)
		log.Printf("[%s] 单比特频数检测 P: %.5f", filename, p)

		// [2] 块内频数检测
		p, _ = randomness.FrequencyWithinBlockProto(bits, 10_000)
		arr = append(arr, p)
		log.Printf("[%s] 块内频数检测 m=100_000 P: %.5f", filename, p)

		// [3] 扑克检测
		p, _ = randomness.PokerProto(bits, 4)
		arr = append(arr, p)
		log.Printf("[%s] 扑克检测 m=4 P: %.5f", filename, p)
		p, _ = randomness.PokerProto(bits, 8)
		arr = append(arr, p)
		log.Printf("[%s] 扑克检测 m=8 P: %.5f", filename, p)

		// [4] 重叠子序列检测
		p1, p2, _, _ := randomness.OverlappingTemplateMatchingProto(bits, 3)
		arr = append(arr, p1, p2)
		log.Printf("[%s] 重叠子序列检测 m=3 P1: %.5f P2: %.5f", filename, p1, p2)
		p1, p2, _, _ = randomness.OverlappingTemplateMatchingProto(bits, 5)
		arr = append(arr, p1, p2)
		log.Printf("[%s] 重叠子序列检测 m=5 P1: %.5f P2: %.5f", filename, p1, p2)

		// [5] 游程总数检测
		p, _ = randomness.RunsTest(bits)
		arr = append(arr, p)
		log.Printf("[%s] 游程总数检测 P: %.5f", filename, p)

		// [6] 游程分布检测
		p, _ = randomness.RunsDistributionTest(bits)
		arr = append(arr, p)
		log.Printf("[%s] 游程分布检测 P: %.5f", filename, p)

		// [7] 块内最大游程检测
		p, _ = randomness.LongestRunOfOnesInABlockTest(bits, true)
		arr = append(arr, p)
		log.Printf("[%s] 块内最大游程检测 P: %.5f", filename, p)

		// [8] 二元推导检测
		p, _ = randomness.BinaryDerivativeProto(bits, 3)
		arr = append(arr, p)
		log.Printf("[%s] 二元推导检测 m=3 P: %.5f", filename, p)
		p, _ = randomness.BinaryDerivativeProto(bits, 7)
		arr = append(arr, p)
		log.Printf("[%s] 二元推导检测 m=7 P: %.5f", filename, p)

		// [9] 自相关检测
		p, _ = randomness.AutocorrelationProto(bits, 1)
		arr = append(arr, p)
		log.Printf("[%s] 自相关检测 m=1 P: %.5f", filename, p)
		p, _ = randomness.AutocorrelationProto(bits, 2)
		arr = append(arr, p)
		log.Printf("[%s] 自相关检测 m=2 P: %.5f", filename, p)
		p, _ = randomness.AutocorrelationProto(bits, 8)
		arr = append(arr, p)
		log.Printf("[%s] 自相关检测 m=8 P: %.5f", filename, p)
		p, _ = randomness.AutocorrelationProto(bits, 16)
		arr = append(arr, p)
		log.Printf("[%s] 自相关检测 m=16 P: %.5f", filename, p)

		// [10] 矩阵秩检测
		p, _ = randomness.MatrixRankTest(bits)
		arr = append(arr, p)
		log.Printf("[%s] 矩阵秩检测 P: %.5f", filename, p)

		// [11] 累加和检测
		p, _ = randomness.CumulativeTest(bits, true)
		arr = append(arr, p)
		log.Printf("[%s] 累加和检测 P: %.5f", filename, p)

		// [12] 近似熵检测
		p, _ = randomness.ApproximateEntropyProto(bits, 2)
		arr = append(arr, p)
		log.Printf("[%s] 近似熵检测 m=2 P: %.5f", filename, p)
		p, _ = randomness.ApproximateEntropyProto(bits, 5)
		arr = append(arr, p)
		log.Printf("[%s] 近似熵检测 m=5 P: %.5f", filename, p)

		// [13] 线性复杂度检测
		p, _ = randomness.LinearComplexityProto(bits, 500)
		arr = append(arr, p)
		log.Printf("[%s] 线性复杂度检测 m=500 P: %.5f", filename, p)
		p, _ = randomness.LinearComplexityProto(bits, 1000)
		arr = append(arr, p)
		log.Printf("[%s] 线性复杂度检测 m=1000 P: %.5f", filename, p)

		// [14] 通用统计检测
		p, _ = randomness.MaurerUniversalTest(bits)
		arr = append(arr, p)
		log.Printf("[%s] Maurer通用统计检测 P: %.5f", filename, p)

		// [15] 离散傅里叶变换检测
		p, _ = randomness.DiscreteFourierTransformTest(bits)
		arr = append(arr, p)
		log.Printf("[%s] 离散傅里叶变换检测 P: %.5f", filename, p)

		log.Printf("[%s] 检测结束\n", filename)

		go func(file string) {
			out <- &R{path.Base(file), arr}
		}(filename)
	}
}