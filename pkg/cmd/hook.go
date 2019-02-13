package cmd

import (
	"github.com/devbuddy/devbuddy/pkg/config"
	"github.com/devbuddy/devbuddy/pkg/env"
	"github.com/devbuddy/devbuddy/pkg/hook"
	"github.com/devbuddy/devbuddy/pkg/project"
	"github.com/devbuddy/devbuddy/pkg/termui"
)

func runHook() {
	// In the shell hook, the stdout is evaluated by the shell
	// stderr is used to display messages to the user

	// Also, we can't annoy the user here, so we always just quit silently

	cfg, err := config.Load()
	if err != nil {
		return
	}

	ui := termui.NewHook(cfg)

	env := env.NewFromOS()

	proj, err := project.FindCurrent()
	if err != nil && err != project.ErrProjectNotFound { // The hook code must run whether the project is found or not
		return
	}

	err = hook.Run(cfg, ui, env, proj, hook.NewCommands())
	if err != nil {
		ui.Debug("%s", err)
	}
}
