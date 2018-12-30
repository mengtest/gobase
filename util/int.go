package util

// Int64Max 求int64的最大值
func Int64Max(a1, a2 int64) int64 {
	if a1 > a2 {
		return a1
	}
	return a2
}

// Int64Min 求int64的最小值
func Int64Min(a1, a2 int64) int64 {
	if a1 < a2 {
		return a1
	}
	return a2
}
