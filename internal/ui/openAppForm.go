package ui

import (
	"sort"

	"rithwik/auto-app-opener/internal/models"

	"github.com/charmbracelet/huh"
)

// Pass a POINTER to the slice so selections flow back to the caller.
func MakeOpenAppMultiSelectForm(selected *[]string, apps []models.App) huh.Form {
	// Optional: sort by named
	sort.Slice(apps, func(i, j int) bool { return apps[i].Name < apps[j].Name })

	multi := huh.NewMultiSelect[string]().
		Title("Select Apps to open").
		OptionsFunc(func() []huh.Option[string] {
			opts := make([]huh.Option[string], 0, len(apps))
			for _, app := range apps {
				opts = append(opts, huh.Option[string]{
					Key:   app.Name,
					Value: app.Name,
				})
			}
			return opts
		}, nil).Value(selected)

	form := huh.NewForm(huh.NewGroup(multi))
	return *form
}
