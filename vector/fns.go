package vector

import "math"

func round11(val float64) float64 {
	precision := 11
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func sqrt(val float64) float64 {
	return round11(math.Sqrt(val))
}

// Calculates the Euclidean distance between two vectors.
func Distance(v0, v1 *Vector) float64 {
	return v0.Distance(v1)
}

// Rotate  vector by an angle.
func Rotate(v0 *Vector, theta float64) *Vector {
	ret := v0.Clone()
	ret.Rotate(theta)
	return ret
}

// Calculates the dot product of two vectors.
func DotProduct(v0, v1 *Vector) float64 {
	return v0.DotProduct(v1)
}

// Calculates linear interpolation from one vector to another vector.
func Lerp(v0, v1 *Vector, amount float64) *Vector {
	ret := v0.Clone()
	ret.Lerp(v1, amount)
	return ret
}

// Multiply v0 vector by a factor. Returns new vector
func Multiply(v0 *Vector, factor float64) *Vector {
	ret := v0.Clone()
	ret.Multiply(factor)
	return ret
}

func Divide(v0 *Vector, factor float64) *Vector {
	ret := v0.Clone()
	ret.Divide(factor)
	return ret
}

func Add(v0, v1 *Vector) *Vector {
	ret := v0.Clone()
	ret.Add(v1)
	return ret
}

func Sub(v0, v1 *Vector) *Vector {
	ret := v0.Clone()
	ret.Sub(v1)
	return ret
}
