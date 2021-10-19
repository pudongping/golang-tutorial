package simplemath

import "math"

// 定义平方根算法
func Sqrt(i int) int {
	v := math.Sqrt(float64(i))
	return int(v)
}
