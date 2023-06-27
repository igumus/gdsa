package vector

import (
	"math"
)

var (
	zeroVector = Create()
)

type Vector struct {
	x float64
	y float64

	// cached value
	length float64
	angle  float64
}

func Create() *Vector {
	return &Vector{x: 0, y: 0, length: 0, angle: 0}
}

func (v *Vector) Equal(o *Vector) bool {
	if v == nil || o == nil {
		return false
	}
	return v.x == o.x && v.y == o.y && v.Angle() == o.Angle()
}

// Creates vector with given x and y coordinates.
func CreateWithPoints(x, y float64) *Vector {
	v := Create()
	v.x = x
	v.y = y

	if x != 0.0 && y != 0.0 {
		v.resetAngle()
		v.resetLength()
	}
	return v
}

// Creates unit vector (https://en.wikipedia.org/wiki/Unit_vector)
// with given angle
func CreateUnit(angle float64) *Vector {
	return CreateWithAngleAndLength(angle, 1.0)
}

// Creates vector according to given angle and length
func CreateWithAngleAndLength(angle, length float64) *Vector {
	v := &Vector{}
	v.length = round11(length)
	v.angle = round11(angle)
	v.calculateCoordinates()
	return v
}

func (v *Vector) X() float64 {
	return v.x
}

func (v *Vector) Y() float64 {
	return v.y
}

// Angle calculates angle between x and y values as radian.
func (v *Vector) Angle() float64 {
	if math.IsInf(v.angle, 1) {
		v.angle = round11(math.Atan2(v.y, v.x))
	}
	return v.angle
}

// Calculates (if not already cached) length of the vector and returns it.
func (v *Vector) Length() float64 {
	if math.IsInf(v.length, 1) {
		v.length = sqrt(v.x*v.x + v.y*v.y)
	}
	return v.length
}

func (v *Vector) Clone() *Vector {
	ret := &Vector{}
	ret.x = v.x
	ret.y = v.y
	ret.length = v.Length()
	ret.angle = v.Angle()
	return ret
}

// Adds other vector to this vector
func (v *Vector) Add(other *Vector) {
	v.x += other.x
	v.y += other.y

	if !zeroVector.Equal(v) {
		v.resetLength()
	}
}

// Substracts other vector to this vector
func (v *Vector) Sub(other *Vector) {
	v.x -= other.x
	v.y -= other.y

	if !zeroVector.Equal(v) {
		v.resetLength()
	}
}

// Multiplies the vector with given scalar factor.
func (v *Vector) Multiply(factor float64) {
	v.scale(factor)
}

// Divides the vector with given scalar factor.
func (v *Vector) Divide(factor float64) {
	v.scale(factor)
}

// Rotate  vector by an angle.
func (v *Vector) Rotate(theta float64) {
	theta = round11(theta)
	v.angle = theta
	cosTheta := round11(math.Cos(theta))
	sinTheta := round11(math.Sin(theta))
	temp := v.x
	v.x = round11(v.x*cosTheta - v.y*sinTheta)
	v.y = round11(temp*sinTheta + v.y*cosTheta)
}

// Normalize the vector to length 1 (make it a unit vector).
func (v *Vector) Normalize() {
	v.length = 1.0
	v.calculateCoordinates()
}

// Calculates the Euclidean distance this vector and other.
func (v *Vector) Distance(other *Vector) float64 {
	x := math.Pow(v.x-other.x, 2)
	y := math.Pow(v.y-other.y, 2)
	return sqrt(x + y)
}

// Calculates dot product of this and other vector.
func (v *Vector) DotProduct(other *Vector) float64 {
	return v.x*other.x + v.y*other.y
}

// Calculates linear interpolation from the vector to other vector.
func (v *Vector) Lerp(other *Vector, amount float64) {
	amount = round11(amount)
	v.x = v.x + (other.x-v.x)*amount
	v.y = v.y + (other.y-v.y)*amount
	v.resetAngle()
	v.resetLength()
}

// Scales the vector with given factor.
// Invalidates length cache if factor is not equals to 1.0, -1.0, 0.0
func (v *Vector) scale(factor float64) {
	if !math.IsInf(factor, 0) && !math.IsInf(factor, 1) && !math.IsInf(factor, 0) && !math.IsInf(factor, -1) {
		factor = round11(factor)
		v.x *= factor
		v.y *= factor

		if factor == 0.0 {
			v.length = 0
			v.angle = 0
			return
		}
		if factor != 1.0 && factor != -1.0 {
			v.resetLength()
		}
	}
}

// Calculates x and y coordinates based on length and angle values
func (v *Vector) calculateCoordinates() {
	angle := v.Angle()
	v.x = round11(math.Cos(angle) * v.length)
	v.y = round11(math.Sin(angle) * v.length)
}

// Resets cached angle value
func (v *Vector) resetAngle() {
	v.angle = math.Inf(1)
}

// Resets cached length value
func (v *Vector) resetLength() {
	v.length = math.Inf(1)
}
