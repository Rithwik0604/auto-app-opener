package ui

import (
	"rithwik/auto-app-opener/internal/models"
	"slices"

	"github.com/charmbracelet/huh"
)

func MakeEditGroupForm(selected *string, groups []string) *huh.Form {
	if len(groups) == 0 {
		return nil
	}
	return huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Select group to edit").
				OptionsFunc(func() []huh.Option[string] {
					opts := []huh.Option[string]{}

					for _, group := range groups {
						opts = append(opts, huh.NewOption(group, group))
					}

					return opts
				}, nil).Value(selected),
		),
	)
}

func MakeEditSpecificGroupForm(groupName *string, cfg *models.Config, oldSelected *[]string, newSelected *[]string) *huh.Form {
	apps := &cfg.Apps
	return huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Group Name").Value(groupName),
		),
		huh.NewGroup(

			huh.NewMultiSelect[string]().
				Title("Select Apps to add").
				OptionsFunc(func() []huh.Option[string] {
					opts := make([]huh.Option[string], 0, len(*apps))

					for _, app := range *apps {
						isSelected := slices.Contains(*oldSelected, app.Name)

						option := huh.Option[string]{
							Key:   app.Name,
							Value: app.Name,
						}.Selected(isSelected)

						opts = append(opts, option)
					}
					return opts
				}, nil).Value(newSelected),
		),
	)
}
