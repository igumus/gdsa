package vector

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVectorCreation(t *testing.T) {
	assert.True(t, IsZeroVector(CreateWithPoints(0.0, 0.0)))
	assert.False(t, IsZeroVector(CreateUnit(0.0)))
	assert.False(t, IsZeroVector(CreateWithAngleAndLength(0.01, 10)))

	v := CreateWithPoints(0.0, 0.0)
	assert.True(t, IsZeroVector(v))

	length := 2.0
	x := sqrt(3)
	assert.Equal(t, 1.73205080757, x)
	y := 1.0
	angle := round11(math.Pi / 6.0)

	v0 := CreateWithPoints(x, y)
	assert.Equal(t, x, v0.x)
	assert.Equal(t, y, v0.y)
	assert.Equal(t, angle, v0.Angle())
	assert.Equal(t, length, v0.Length())

	v1 := CreateWithAngleAndLength(angle, 2)

	assert.Equal(t, v0.x, v1.x)
	assert.Equal(t, v0.y, v1.y)
	assert.Equal(t, v0.length, v1.length)
	assert.Equal(t, v0.Angle(), v1.Angle())
}

// Test for vector scaling using scale method in vector.
// The test also covers Divide and Multiply methods.
func TestVectorScale(t *testing.T) {
	testcases := []struct {
		name            string
		v               *Vector
		factor          float64
		lengthResettted bool
	}{
		{
			name:            "scaleByZero",
			v:               CreateWithPoints(10.0, 20.0),
			factor:          0.0,
			lengthResettted: false,
		},
		{
			name:            "scaleByOne",
			v:               CreateWithPoints(10.0, 20.0),
			factor:          1.0,
			lengthResettted: false,
		},
		{
			name:            "scaleByNegativeOne",
			v:               CreateWithPoints(10.0, 20.0),
			factor:          -1.0,
			lengthResettted: false,
		},
		{
			name:            "scaleByPositiveInfinity",
			v:               CreateWithPoints(10.0, 20.0),
			factor:          math.Inf(1),
			lengthResettted: false,
		},
		{
			name:            "scaleByZeroInfinity",
			v:               CreateWithPoints(10.0, 20.0),
			factor:          math.Inf(0),
			lengthResettted: false,
		},
		{
			name:            "scaleByNegativeInfinity",
			v:               CreateWithPoints(10.0, 20.0),
			factor:          math.Inf(-1),
			lengthResettted: false,
		},
		{
			name:            "scaleByPositiveScalar",
			v:               CreateWithPoints(10.0, 20.0),
			factor:          3.0,
			lengthResettted: true,
		},
		{
			name:            "scaleByPositiveFriction",
			v:               CreateWithPoints(10.0, 20.0),
			factor:          3.50,
			lengthResettted: true,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			factor := tc.factor
			var (
				dx, dy float64
			)
			tc.v.Length()
			tc.v.Angle()
			if !math.IsInf(factor, 0) && !math.IsInf(factor, 1) && !math.IsInf(factor, 0) && !math.IsInf(factor, -1) {
				dx = factor * tc.v.x
				dy = factor * tc.v.y
			} else {
				dx = tc.v.x
				dy = tc.v.y
			}
			tc.v.scale(tc.factor)
			assert.Equal(t, dx, tc.v.x)
			assert.Equal(t, dy, tc.v.y)
			assert.Equal(t, tc.lengthResettted, math.IsInf(tc.v.length, 1))
		})
	}
}

func TestVectorRotate(t *testing.T) {
	expectedX := -20.0
	expectedY := 10.0
	angle := round11(math.Pi / 2.0)
	v0 := CreateWithPoints(10.0, 20.0)
	v1 := v0.Clone()
	v1.Rotate(angle)

	assert.Equal(t, expectedX, v1.x)
	assert.Equal(t, expectedY, v1.y)
	assert.Equal(t, v0.length, v1.length)

	v2 := Rotate(v0, angle)
	assert.Equal(t, expectedX, v2.x)
	assert.Equal(t, expectedY, v2.y)
	assert.Equal(t, v0.length, v2.length)
}

func TestVectorDotProduct(t *testing.T) {
	expected := 2200.0
	v1 := CreateWithPoints(10, 20)
	v2 := CreateWithPoints(60, 80)
	assert.Equal(t, expected, v1.DotProduct(v2))
	assert.Equal(t, expected, DotProduct(v1, v2))
}

func TestVectorLerp(t *testing.T) {
	expected := CreateWithPoints(50.0, 50.0)
	amount := 0.5
	target := CreateWithPoints(100.0, 100.0)

	v0 := Create()
	v0.Lerp(target, amount)
	assert.Equal(t, expected.x, v0.x)
	assert.Equal(t, expected.y, v0.y)

	v0 = Create()
	v2 := Lerp(v0, target, amount)
	assert.Equal(t, expected.x, v2.x)
	assert.Equal(t, expected.y, v2.y)

}

func TestVectorDistance(t *testing.T) {
	v0 := CreateWithPoints(10.0, 20.0)
	v1 := CreateWithPoints(60.0, 80.0)
	expectedDistance := 78.10249675907
	assert.Equal(t, expectedDistance, v0.Distance(v1))
	assert.Equal(t, expectedDistance, Distance(v0, v1))
}

func TestVectorUnit(t *testing.T) {
	v := CreateUnit(0.01)
	expectedX := round11(0.9999500004166653)
	expectedY := round11(0.009999833334166664)
	assert.Equal(t, v.x, expectedX)
	assert.Equal(t, v.y, expectedY)
}

func TestVectorAngleCalculation(t *testing.T) {
	v := CreateWithPoints(10.0, 20.0)
	assert.Equal(t, 1.10714871779, v.Angle())
}

func BenchmarkAngleCalculation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f := float64(i)
		v := CreateWithPoints(f, f)
		v.Angle()
	}
}
