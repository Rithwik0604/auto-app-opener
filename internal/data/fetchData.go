package data

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"rithwik/auto-app-opener/internal/models"
)

// retrieves all apps from start programs shortcuts
func RetrieveAppsPowershell(cfg *models.Config) error {
	if cfg.Apps == nil {
		cfg.Apps = []models.App{}
	}

	cmd := exec.Command("powershell.exe", "-Command",
		"Get-StartApps | ConvertTo-Json")

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("powershell failed: %s", err.Error())
	}

	var allApps []models.App

	decoder := json.NewDecoder(stdout)

	err = decoder.Decode(&allApps)
	if err != nil {
		return err
	}

	cfg.Apps = allApps

	return nil
}
