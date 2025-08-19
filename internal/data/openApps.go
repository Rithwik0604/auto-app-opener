package data

import (
	"fmt"
	"os/exec"
	"rithwik/auto-app-opener/internal/models"
)

func openApp(appID string) error {
	const psCommand = "Start-process \"shell:AppsFolder\\"

	cmd := exec.Command("powershell.exe", "-Command", psCommand+appID+"\"")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to open app %s: %w", appID, err)
	}
	return nil
}

func OpenApps(selected *[]string, apps *[]models.App) error {

	for _, selectedApp := range *selected {
		name := selectedApp
		var appID string

		// get appID
		for _, app := range *apps {
			if app.Name == name {
				appID = app.AppID
				break
			}
		}
		if appID == "" {
			return fmt.Errorf("app ID not found for app: %s", name)
		}
		fmt.Println("Opening " + name + " ...")
		err := openApp(appID)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}

	return nil
}

func OpenGroup(groupName string, cfg *models.Config) error {
	group := cfg.Groups[groupName]
	if group == nil {
		return fmt.Errorf("no group %s", groupName)
	}

	for _, app := range group {
		fmt.Println("Opening App " + app.Name + " ...")
		err := openApp(app.AppID)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}

	return nil
}
