package ball_test

import (
	"github.com/VSS-SDK/VSS-SampleGo/ball"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestZeroBall(t *testing.T) {
	ball := ball.ZeroBall()

	assert.NotNil(t, ball)
	assert.Zero(t, ball.X)
	assert.Zero(t, ball.Y)
	assert.Zero(t, ball.SpeedX)
	assert.Zero(t, ball.SpeedY)
}

func TestRandBall(t *testing.T) {
	ball := ball.RandBall()

	assert.NotNil(t, ball)
	assert.NotZero(t, ball.X)
	assert.NotZero(t, ball.Y)
	assert.NotZero(t, ball.SpeedX)
	assert.NotZero(t, ball.SpeedY)
}

func TestNewBall(t *testing.T) {
	rand := ball.RandBall()
	ball := ball.NewBall(rand.X, rand.Y, rand.SpeedX, rand.SpeedY)

	assert.NotNil(t, ball)
	assert.Equal(t, ball.X, rand.X)
	assert.Equal(t, ball.Y, rand.Y)
	assert.Equal(t, ball.SpeedX, rand.SpeedX)
	assert.Equal(t, ball.SpeedY, rand.SpeedY)
}
