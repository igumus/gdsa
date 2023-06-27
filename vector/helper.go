package vector

import "math"

// Checks inifinity with all sings (zero, positive, negative)
func isInfinity(factor float64) bool {
	return math.IsInf(factor, 0) || math.IsInf(factor, 1) || math.IsInf(factor, -1)
}

// Rounds value's decimal part to given precision
func round(val float64, precision int) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

// Rounds value's decimal part to 11 digits.
func round11(val float64) float64 {
	return round(val, 11)
}

// Calculates round11 of given values square root.
func sqrt(val float64) float64 {
	return round11(math.Sqrt(val))
}
