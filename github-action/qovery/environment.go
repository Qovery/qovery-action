package qovery

import (
	"fmt"
	"github-action/pkg"
)

func GetEnvironmentIdByName(qoveryAPIClient pkg.QoveryAPIClient, projectId string, name string) (string, error) {
	environments, err := qoveryAPIClient.ListEnvironments(projectId)
	if err != nil {
		return "", err
	}

	for _, env := range environments {
		if env.Name == name {
			return env.ID, nil
		}
	}

	return "", fmt.Errorf("can't find environment with name %v! (it's case sensitive)", name)
}
