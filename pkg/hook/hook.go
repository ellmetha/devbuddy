package hook

import (
	"github.com/devbuddy/devbuddy/pkg/config"
	"github.com/devbuddy/devbuddy/pkg/env"
	"github.com/devbuddy/devbuddy/pkg/features"
	"github.com/devbuddy/devbuddy/pkg/project"
	"github.com/devbuddy/devbuddy/pkg/tasks/taskapi"
	"github.com/devbuddy/devbuddy/pkg/termui"
)

// Run is responsible for emitting the commands needed to mutate the shell to the desired state
func Run(cfg *config.Config, ui *termui.UI, env *env.Env, proj *project.Project, commands *Commands) error {
	ui.Debug("project: %+v", proj)

	allFeatures, err := getFeaturesFromProject(proj)
	if err != nil {
		return err
	}
	ui.Debug("features: %+v", allFeatures)

	features.Sync(cfg, proj, ui, env, allFeatures)

	for _, change := range env.Changed() {
		ui.Debug("Env change: %+v", change)

		if change.Deleted {
			commands.UnsetEnv(change.Name)
		} else {
			commands.SetEnv(change.Name, change.Value)
		}
	}

	return nil
}

func getFeaturesFromProject(proj *project.Project) (features.FeatureSet, error) {
	if proj == nil {
		// When no project was found, we must deactivate all potentially active features
		// So we continue with an empty feature map
		return features.NewFeatureSet(), nil
	}
	allTasks, err := taskapi.GetTasksFromProject(proj)
	if err != nil {
		return nil, err
	}
	return taskapi.GetFeaturesFromTasks(allTasks), nil
}
