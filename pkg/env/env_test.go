package env

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetActiveFeatures(t *testing.T) {
	envs := [][]string{
		[]string{},
		[]string{"DEV_AUTO_ENV_FEATURES="},
		[]string{"DEV_AUTO_ENV_FEATURES=f1=v1"},
		[]string{"DEV_AUTO_ENV_FEATURES=f1=v1:f2=v2"},
	}
	features := []map[string]string{
		map[string]string{},
		map[string]string{},
		map[string]string{"f1": "v1"},
		map[string]string{"f1": "v1", "f2": "v2"},
	}

	for idx := range envs {
		env := New(envs[idx])
		require.Equal(t, features[idx], env.GetActiveFeatures())
	}
}

func TestSetFeatures(t *testing.T) {
	features := []map[string]string{
		map[string]string{},
		map[string]string{},
		map[string]string{"f1": "v1"},
		map[string]string{"f1": "v1", "f2": "v2"},
	}

	for idx := range features {
		env := New([]string{})
		env.setActiveFeatures(features[idx])
		require.Equal(t, features[idx], env.GetActiveFeatures())
	}
}

func TestChanged(t *testing.T) {
	env := New([]string{})
	require.Equal(t, []VariableChange{}, env.Changed())

	env.Set("K2", "1")
	require.Equal(t, []VariableChange{VariableChange{"K2", "1", false}}, env.Changed())

	env.Set("K2", "2")
	require.Equal(t, []VariableChange{VariableChange{"K2", "2", false}}, env.Changed())

	env = New([]string{"K1=1"})
	env.Unset("K1")
	require.Equal(t, []VariableChange{VariableChange{"K1", "", true}}, env.Changed())
}
