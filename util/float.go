package util

import "math"

// Decimal2 保留两位小数点
func Decimal2(value float64) float64 {
	return math.Trunc(value*1e2+0.5) * 1e-2
}
