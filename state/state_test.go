package state_test

import (
	"github.com/VSS-SDK/VSS-SampleGo/state"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestZeroCommand(t *testing.T) {
	state := state.ZeroState()

	assert.NotNil(t, state)
	assert.Empty(t, state.YellowRobots)
	assert.Empty(t, state.BlueRobots)
	assert.NotNil(t, state.Ball)
}

func TestRandState(t *testing.T) {
	state := state.RandState()

	assert.NotNil(t, state)
	assert.NotEmpty(t, state.YellowRobots)
	assert.NotEmpty(t, state.BlueRobots)
	assert.NotNil(t, state.Ball)

	for _, blueRobot := range state.BlueRobots {
		assert.NotZero(t, blueRobot)
	}

	for _, yellowRobot := range state.YellowRobots {
		assert.NotZero(t, yellowRobot)
	}
}

func TestNewState(t *testing.T) {
	rand := state.RandState()
	state := state.NewState(rand.YellowRobots, rand.BlueRobots, rand.Ball)

	assert.NotNil(t, state)
	assert.ElementsMatch(t, state.BlueRobots, rand.BlueRobots)
	assert.ElementsMatch(t, state.YellowRobots, rand.YellowRobots)
	assert.Equal(t, state.Ball, rand.Ball)
}
