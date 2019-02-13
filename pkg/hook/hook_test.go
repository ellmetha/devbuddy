package hook

import (
	"bytes"
	"testing"

	"github.com/Flaque/filet"
	"github.com/stretchr/testify/require"

	"github.com/devbuddy/devbuddy/pkg/config"
	"github.com/devbuddy/devbuddy/pkg/env"
	"github.com/devbuddy/devbuddy/pkg/project"
	"github.com/devbuddy/devbuddy/pkg/termui"
	"github.com/devbuddy/devbuddy/pkg/test"
)

func buildContext(projectPath string) (*config.Config, *termui.UI, *project.Project, *Commands, *bytes.Buffer) {
	_, ui := termui.NewTesting(false)
	proj := project.NewFromPath(projectPath)
	output, commands := NewTestCommands()
	return config.NewTestConfig(), ui, proj, commands, output
}

func TestRun(t *testing.T) {
	tmpdir := filet.TmpDir(t, "")
	defer filet.CleanUp(t)

	test.Project(tmpdir).Manifest().WriteString(t, "")

	cfg, ui, proj, commands, output := buildContext(tmpdir)
	env := env.New([]string{})

	Run(cfg, ui, env, proj, commands)

	require.Equal(t, "", output.String())
}
