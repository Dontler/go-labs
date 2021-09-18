package labs

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"time"
)

type Config struct {
	N int
	A int
	B int
}

type FirstLab struct{
	Config Config
}

func (fl *FirstLab) ThirdTask() {
	x := make([]float32, fl.Config.N)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < fl.Config.N; i++ {
		x[i] = rand.Float32()
	}
	f := func (x float32) float32 {
		return x * x
	}

	fmt.Println(calculateFunction(fl.Config.A, fl.Config.B, x, f))
}

func (fl *FirstLab) SeventhTask() {
	x := make([]int, fl.Config.N)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(x); i++ {
		x[i] = rand.Int() % 10
	}

	fmt.Println(x)
	median := getMedian(x)
	fmt.Println(x)
	fmt.Println(fmt.Sprintf("Median: %d\n", median))
}

func (fl *FirstLab) TenthTask() {
	var x = make([]int, fl.Config.N)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(x); i++ {
		x[i] = rand.Int() % 16 - 5
	}

	fmt.Println(x)
	products := calculatePairProducts(x)
	fmt.Println(products)
}

func (fl *FirstLab) EleventhTask() {
	x := make([]int, fl.Config.N)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(x); i++ {
		x[i] = rand.Int() % 16 - 5
	}

	maxSum, subsequence := calculateMaxPositiveSubsequenceSum(x)
	fmt.Println("Initial array", x)
	fmt.Println("Max sum:", maxSum, "Subsequence:", subsequence)
}

func calculateFunction(a int, b int, x []float32, f func(float32) float32) float32 {
	var sum float32 = 0.0
	for i := a; i < b; i++ {
		sum += f(x[i])
	}

	return sum
}

func getMedian(x []int) int {
	tmp := make([]int, len(x))
	copy(tmp, x)
	sort.Ints(tmp)
	return tmp[len(tmp) / 2]
}

func calculatePairProducts(x []int) []int {
	var products = make([]int, len(x) - 1)

	for i := 0; i < len(x) - 1; i++ {
		products[i] = x[i] * x[i + 1]
	}

	return products
}

func calculateMaxPositiveSubsequenceSum(x []int) (int, string) {
	sum := 0
	i := 0
	subsequence := ""
	for i < len(x) {
		curSum := 0
		curSubsequence := ""
		for i < len(x) && x[i] >= 0 {
			curSum += x[i]
			curSubsequence += strconv.Itoa(x[i]) + " "
			i++
		}

		if curSum > sum {
			sum = curSum
			subsequence = curSubsequence
		}
		i++
	}

	return sum, subsequence
}