package ui

import (
	"rithwik/auto-app-opener/internal/data"
	"rithwik/auto-app-opener/internal/models"
	"sort"

	"github.com/charmbracelet/huh"
)

func MakeCreateGroupForm(groupName *string, apps *[]models.App, selected *[]string) *huh.Form {
	sort.Slice(*apps, func(i, j int) bool { return (*apps)[i].Name < (*apps)[j].Name })

	return huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Group Name").Value(groupName).Validate(data.ValidateGroupName),
		),
		huh.NewGroup(

			huh.NewMultiSelect[string]().
				Title("Select Apps to add").
				OptionsFunc(func() []huh.Option[string] {
					opts := make([]huh.Option[string], 0, len(*apps))
					for _, app := range *apps {
						opts = append(opts, huh.Option[string]{
							Key:   app.Name,
							Value: app.Name,
						})
					}
					return opts
				}, nil).Value(selected),
		),
	)
}
