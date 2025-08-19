package ui

import (
	"rithwik/auto-app-opener/internal/models"

	"github.com/charmbracelet/huh"
)

func MakeOpenGroupsForm(selected *string, cfg *models.Config) *huh.Form {
	if len(cfg.Groups) == 0 {
		return nil
	}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Select Group to open").
				Description("Opens up a group of apps").
				OptionsFunc(func() []huh.Option[string] {
					opts := make([]huh.Option[string], 0, len(cfg.Groups))
					for group := range cfg.Groups {
						opts = append(opts, huh.Option[string]{
							Key:   group,
							Value: group,
						})
					}
					return opts
				}, nil).Value(selected),
		),
	)
	return form
}
