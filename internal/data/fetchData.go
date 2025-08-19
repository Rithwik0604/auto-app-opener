package data

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"rithwik/auto-app-opener/internal/models"
)

// retrieves all apps from start programs shortcuts
func RetrieveAppsPowershell(cfg *models.Config) error {
	if cfg.Apps == nil {
		cfg.Apps = []models.App{}
	}

	cmd := exec.Command("powershell.exe", "-Command",
		"Get-StartApps | ConvertTo-Json -Depth 2 | Set-Content out.json")

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("powershell failed: %w", err)
	}

	file, err := os.Open("out.json")
	if err != nil {
		return err
	}

	var allApps []models.App

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&allApps)
	if err != nil {
		return err
	}

	cfg.Apps = allApps

	file.Close()

	if err = os.Remove("out.json"); err != nil {
		return err
	}

	return nil
}
