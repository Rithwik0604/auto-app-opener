package ui

import (
	"rithwik/auto-app-opener/internal/models"

	"github.com/charmbracelet/huh"
)

func MakeFirstForm(value *models.FirstFormOptionEnum) huh.Form {
	return *huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[models.FirstFormOptionEnum]().Title("").
				Options(
					huh.NewOption(models.OpenApp.String(), models.OpenApp),
					huh.NewOption(models.OpenGroupEnum.String(), models.OpenGroupEnum),
					huh.NewOption(models.ManageGroupsEnum.String(), models.ManageGroupsEnum),
					huh.NewOption(models.RefetchAppsEnum.String(), models.RefetchAppsEnum),
					huh.NewOption(models.QuitEnum.String(), models.QuitEnum),
				).
				Value(value),
		),
	)
}
