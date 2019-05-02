package pose_test

import (
	"github.com/VSS-SDK/VSS-SampleGo/pose"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestZeroPoint(t *testing.T) {
	pose := pose.ZeroPose()

	assert.NotNil(t, pose)
	assert.Zero(t, pose.X)
	assert.Zero(t, pose.Y)
	assert.Zero(t, pose.Angle)
}

func TestRandPose(t *testing.T) {
	pose := pose.RandPose()

	assert.NotNil(t, pose)
	assert.NotZero(t, pose.X)
	assert.NotZero(t, pose.Y)
	assert.NotZero(t, pose.Angle)
}

func TestNewPose(t *testing.T) {
	rand := pose.RandPose()
	pose := pose.NewPose(rand.X, rand.Y, rand.Angle)

	assert.NotNil(t, pose)
	assert.Equal(t, rand.X, pose.X)
	assert.Equal(t, rand.Y, pose.Y)
	assert.Equal(t, rand.Angle, pose.Angle)
}
