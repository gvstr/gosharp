package math

import "golang.org/x/exp/constraints"

// Returns the larger of two ordered numbers (numbers that supports <,<=,>=,> operators).
func Max[T constraints.Ordered](val1, val2 T) T {
	if val1 > val2 {
		return val1
	}
	return val2
}
