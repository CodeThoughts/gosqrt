package main

import (
	"fmt"
	"math"
	"math/big"
)

func Sqrt(x float64) float64 {
	if x == 0 {
		return x
	}
	return auxSqrt(x, 1, 1.0)
}

func auxSqrt(x float64, i int, z float64) float64 {
	if i == 10 {
		return z
	}
	return auxSqrt(x, i + 1, z - (math.Pow(z, 2) - x) / (2 * z))
}

func main() {
	i := 1.0
	for tranc8(Sqrt(i)) == tranc8(math.Sqrt(i)) {i++}
	fmt.Println(i)
}

func tranc8(x float64) string {
	return big.NewFloat(x).Text('f', 8)
}

