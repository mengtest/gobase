package util

import "math"

// Decimal2 保留两位小数点
func Decimal2(value float64) float64 {
	return math.Trunc(value*1e2+0.5) * 1e-2
}

// Float64Max 用于求float64的最大值
func Float64Max(a1, a2 float64) float64 {
	if a1 > a2 {
		return a1
	}
	return a2
}

// Float64Min 用于求float64的最小值
func Float64Min(a1, a2 float64) float64 {
	if a1 < a2 {
		return a1
	}
	return a2
}
