package vector

// Checks given vector is zero vector.
// Zero vector is a vector which all fields (x,y,angle,length) is zero
func IsZeroVector(v *Vector) bool {
	return zeroVector.Equal(v)
}

// Calculates the Euclidean distance between two vectors.
func Distance(v0, v1 *Vector) float64 {
	return v0.Distance(v1)
}

// Calculates the dot product of two vectors.
func DotProduct(v0, v1 *Vector) float64 {
	return v0.DotProduct(v1)
}

// Rotates given vector by theta angle without mutating original vector.
func Rotate(v0 *Vector, theta float64) *Vector {
	ret := v0.Clone()
	ret.Rotate(theta)
	return ret
}

// Calculates linear interpolation (lerp) from one vector to another vector
// without mutating vectors.
func Lerp(v0, v1 *Vector, amount float64) *Vector {
	ret := v0.Clone()
	ret.Lerp(v1, amount)
	return ret
}

// Multiplies v0 vector by factor without mutating original vector
func Multiply(v0 *Vector, factor float64) *Vector {
	ret := v0.Clone()
	ret.Multiply(factor)
	return ret
}

// Divides v0 vector by factor without mutating original vector
func Divide(v0 *Vector, factor float64) *Vector {
	ret := v0.Clone()
	ret.Divide(factor)
	return ret
}

// Adds v1 onto v0 vector without mutating originial vectors.
func Add(v0, v1 *Vector) *Vector {
	ret := v0.Clone()
	ret.Add(v1)
	return ret
}

// Substracts v1 from v0 vector without mutating original vectors.
func Sub(v0, v1 *Vector) *Vector {
	ret := v0.Clone()
	ret.Sub(v1)
	return ret
}
