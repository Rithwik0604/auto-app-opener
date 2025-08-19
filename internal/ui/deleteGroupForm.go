package ui

import "github.com/charmbracelet/huh"

func MakeDeleteGroupForm(groupName *[]string, groups *[]string) *huh.Form {

	return huh.NewForm(
		huh.NewGroup(
			huh.NewMultiSelect[string]().OptionsFunc(func() []huh.Option[string] {
				opts := []huh.Option[string]{}

				for _, g := range *groups {
					s := huh.NewOption[string](g, g)
					opts = append(opts, s)
				}

				return opts

			}, nil).Value(groupName),
		).Title("Select Group To delete"),
	)
}
