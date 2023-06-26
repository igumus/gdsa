package vector

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVectorCreation(t *testing.T) {
	v := CreateWithPoints(0.0, 0.0)
	assert.Equal(t, v, ZeroVector)

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

	v0 := ZeroVector.Clone()
	v0.Lerp(target, amount)
	assert.Equal(t, expected.x, v0.x)
	assert.Equal(t, expected.y, v0.y)

	v0 = ZeroVector.Clone()
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

func TestVectorAngle(t *testing.T) {
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
