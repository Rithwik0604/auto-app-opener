package ui

import (
	"rithwik/auto-app-opener/internal/models"

	"github.com/charmbracelet/huh"
)

func MakeManageGroupsForm(selected *models.ManageFormOptionEnum) *huh.Form {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[models.ManageFormOptionEnum]().
				Title("").
				Options(
					huh.NewOption(models.CreateGroupEnum.String(), models.CreateGroupEnum),
					huh.NewOption(models.EditGroupEnum.String(), models.EditGroupEnum),
					huh.NewOption(models.DeleteGroupEnum.String(), models.DeleteGroupEnum),
				).Value(selected),
		),
	)

	return form
}
