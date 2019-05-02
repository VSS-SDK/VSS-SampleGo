package point_test

import (
	"github.com/VSS-SDK/VSS-SampleGo/point"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestZeroPoint(t *testing.T) {
	point := point.ZeroPoint()

	assert.NotNil(t, point)
	assert.Zero(t, point.X)
	assert.Zero(t, point.Y)
}

func TestRandPoint(t *testing.T) {
	point := point.RandPoint()

	assert.NotNil(t, point)
	assert.NotZero(t, point.X)
	assert.NotZero(t, point.Y)
}

func TestNewPoint(t *testing.T) {
	rand := point.RandPoint()
	point := point.NewPoint(rand.X, rand.Y)

	assert.NotNil(t, point)
	assert.Equal(t, rand.X, point.X)
	assert.Equal(t, rand.Y, point.Y)
}
