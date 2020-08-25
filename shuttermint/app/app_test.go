package app

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewShutterApp(t *testing.T) {
	app := NewShutterApp()
	require.Equal(t, len(app.Configs), 1)
	require.Equal(t, app.Configs[0], &BatchConfig{})
}

func TestGetBatch(t *testing.T) {
	app := NewShutterApp()

	err := app.addConfig(BatchConfig{StartBatchIndex: 100, Threshhold: 1})
	require.Nil(t, err)

	err = app.addConfig(BatchConfig{StartBatchIndex: 200, Threshhold: 2})
	require.Nil(t, err)

	err = app.addConfig(BatchConfig{StartBatchIndex: 300, Threshhold: 3})
	require.Nil(t, err)

	require.Equal(t, app.getBatch(0).Config.Threshhold, uint32(0))
	require.Equal(t, app.getBatch(99).Config.Threshhold, uint32(0))
	require.Equal(t, app.getBatch(100).Config.Threshhold, uint32(1))
	require.Equal(t, app.getBatch(101).Config.Threshhold, uint32(1))
	require.Equal(t, app.getBatch(199).Config.Threshhold, uint32(1))
	require.Equal(t, app.getBatch(200).Config.Threshhold, uint32(2))
	require.Equal(t, app.getBatch(1000).Config.Threshhold, uint32(3))
}
