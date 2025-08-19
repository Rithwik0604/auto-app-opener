package data

import (
	"fmt"
	"rithwik/auto-app-opener/internal/models"
	"slices"
)

// Retrieves and returns the name of all the groups
func GetAllGroupNames(cfg *models.Config) []string {
	groupNames := []string{}
	for group := range cfg.Groups {
		groupNames = append(groupNames, group)
	}

	return groupNames
}

// Returns all the names of the apps in a group
func GetAllAppNamesInGroup(cfg *models.Config, groupName string) []string {
	apps := []string{}

	for _, app := range cfg.Groups[groupName] {
		apps = append(apps, app.Name)
	}

	return apps
}

func ModifyGroupApps(selected *[]string, cfg *models.Config) []models.App {
	apps := []models.App{}
	for _, app := range cfg.Apps {
		if slices.Contains(*selected, app.Name) {
			apps = append(apps, app)
		}
	}
	return apps
}

func DeleteGroups(groupNames []string, cfg *models.Config) {
	for _, name := range groupNames {
		delete(cfg.Groups, name)
	}
}

func ValidateGroupName(groupName string) error {
	if groupName == "" {
		return fmt.Errorf("group name must be at least 1 character")
	}
	return nil
}
