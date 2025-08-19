package main

import (
	"fmt"
	"log"
	"os"
	"rithwik/auto-app-opener/internal/data"
	"rithwik/auto-app-opener/internal/models"
	"rithwik/auto-app-opener/internal/storage"
	"rithwik/auto-app-opener/internal/ui"

	"github.com/charmbracelet/huh"
)

var (
	isInit bool
	config models.Config
)

var (
	firstFormValue       models.FirstFormOptionEnum
	selectedGroup        string
	manageGroupFormValue models.ManageFormOptionEnum
	form                 *huh.Form
)

func main() {

	ui.MakeAsciiArt()

	ui.InitSpinner.Action(initApp).Run()

	runFetchForm()

	noApps := len(config.Apps)
	noGroups := len(config.Groups)

	groupsWord := "group"

	if noGroups != 1 {
		groupsWord += "s"
	}

	fmt.Printf("%d apps found\nYou have %d %s\n", noApps, noGroups, groupsWord)

	runFirstForm()
}

func initApp() {
	isInit = storage.InitialiseStorage()
}

func fetch() {
	err := storage.ReadStorage(&config)
	if err != nil {
		panic(err)
	}
	if isInit || len(config.Apps) == 0 {
		if err := data.RetrieveAppsPowershell(&config); err != nil {
			log.Fatalln(err)
		}
		storage.WriteStorage(&config)
	}
}

func runFetchForm() {
	ui.FetchSpinner.Action(fetch).Run()
}

func runFirstForm() {
	firstForm := ui.MakeFirstForm(&firstFormValue)
	err := firstForm.Run()
	if err != nil {
		end("Till next time...", 0)
	}
	firstFormSelection()
}

func firstFormSelection() {
	switch firstFormValue {
	case models.OpenApp:
		var selectedApps []string
		openAppForm := ui.MakeOpenAppMultiSelectForm(&selectedApps, config.Apps)
		err := openAppForm.Run()
		if err != nil {
			end("", 0)
		}
		err = data.OpenApps(&selectedApps, &config.Apps)
		if err != nil {
			fmt.Println(fmt.Errorf("something went wrong\n%e", err))
			return
		}

	case models.OpenGroupEnum:
		openGroupForm := ui.MakeOpenGroupsForm(&selectedGroup, &config)
		if openGroupForm == nil {
			fmt.Println("No groups. Create one first in 'Manage Groups'")
			runFirstForm()
			return
		} else {
			err := openGroupForm.Run()
			if err != nil {
				end("", 0)
			}
		}
		data.OpenGroup(selectedGroup, &config)
		end("Quitting...", 0)

	case models.ManageGroupsEnum:
		manageGroupsForm := ui.MakeManageGroupsForm(&manageGroupFormValue)
		err := manageGroupsForm.Run()
		if err != nil {
			end("", 0)
		}
		manageGroupFormSelection()

	case models.RefetchAppsEnum:
		runFetchForm()
		noApps := len(config.Apps)
		fmt.Printf("Re-fetched all apps. %d apps found\n", noApps)
		runFirstForm()

	case models.QuitEnum:
		end("Till next time...", 0)
	}

	end("", 0)
}

func manageGroupFormSelection() {
	switch manageGroupFormValue {
	case models.CreateGroupEnum:
		var groupName string
		var selectedApps []string

		form = ui.MakeCreateGroupForm(&groupName, &config.Apps, &selectedApps)
		err := form.Run()
		if err != nil {
			end("", 0)
		}

		config.Groups[groupName] = data.ModifyGroupApps(&selectedApps, &config)
		storage.WriteStorage(&config)

		fmt.Printf("%s group created with %d apps\n", groupName, len(selectedApps))

	case models.EditGroupEnum:
		var groupName string
		groups := data.GetAllGroupNames(&config)

		form = ui.MakeEditGroupForm(&groupName, groups)
		if form == nil {
			fmt.Println("No groups. Create one first in 'Manage Groups -> Create Group'")
			break
		}
		err := form.Run()
		if err != nil {
			end("", 0)
		}

		oldName := groupName

		selectedApps := data.GetAllAppNamesInGroup(&config, groupName)

		newSelected := []string{}

		form = ui.MakeEditSpecificGroupForm(&groupName, &config, &selectedApps, &newSelected)
		err = form.Run()
		if err != nil {
			end("", 0)
		}

		if groupName != oldName {
			delete(config.Groups, oldName)
		}
		config.Groups[groupName] = data.ModifyGroupApps(&newSelected, &config)

		fmt.Printf("%s group updated. %d apps\n", groupName, len(newSelected))

		storage.WriteStorage(&config)

	case models.DeleteGroupEnum:
		var groupNames []string
		groups := data.GetAllGroupNames(&config)

		form = ui.MakeDeleteGroupForm(&groupNames, &groups)
		if err := form.Run(); err != nil {
			end("", 1)
		}

		data.DeleteGroups(groupNames, &config)

		storage.WriteStorage(&config)

		for _, g := range groupNames {
			fmt.Printf("Group '%s' deleted\n", g)
		}

	}

	runFirstForm()

}

func end(message string, exitStatus int) {
	fmt.Println(message)
	os.Exit(exitStatus)
}
