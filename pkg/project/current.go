package project

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/devbuddy/devbuddy/pkg/manifest"
)

var ErrProjectNotFound = errors.New("project not found")

func FindCurrent() (*Project, error) {
	path, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("error while searching for project: %s", err)
	}
	return findByPath(path)
}

func findByPath(path string) (*Project, error) {
	for {
		if manifest.ExistsIn(path) {
			return NewFromPath(path), nil
		}

		// Continue searching in top directory
		path = filepath.Dir(path)
		if path == "/" {
			break
		}
	}

	return nil, ErrProjectNotFound
}
