package utils

import "strconv"

// Abs - Просто потому что я могу так сделать
func Abs(n int64) int64 {
	y := n >> 63
	return (n ^ y) - y
}

// LimitMaxValue - Я просто линивый
func LimitMaxValue(n int64, max int64) int64 {
	if n > max {
		return max
	}
	return n
}

// LimitMinValue - Крайне ленивый
func LimitMinValue(n int64, min int64) int64 {
	if n < min {
		return min
	}
	return n
}

// Uint32 - Преобразуем в uint32
func Uint32(value interface{}) uint32 {
	switch v := value.(type) {
	case string:
		if r, err := strconv.ParseUint(v, 10, 32); err == nil {
			return uint32(r)
		}
	case int64:
		return uint32(v)
	default:
		return 0
	}
	return 0
}

// Uint16 - Преобразуем в uint16
func Uint16(value interface{}) uint16 {
	switch v := value.(type) {
	case string:
		if r, err := strconv.ParseUint(v, 10, 16); err == nil {
			return uint16(r)
		}
	case int64:
		return uint16(v)
	default:
		return 0
	}
	return 0
}

// Find returns the smallest index i at which x == a[i],
// or len(a) if there is no such index.
func Find(a []string, x string) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}
	return len(a)
}
