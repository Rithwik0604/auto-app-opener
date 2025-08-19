package storage

import (
	"encoding/json"
	"os"
	"path"
	"rithwik/auto-app-opener/internal/models"
)

var storagePath string

// Initialises the storage json file in the user's home directory
// Returns true if the file didn't exist,
// False otherwise
func InitialiseStorage() bool {
	userFolder := os.Getenv("USERPROFILE")
	storagePath = path.Join(userFolder, "autoappopener_storage.json")
	file, err := os.OpenFile(storagePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		panic(err)
	}
	if info.Size() == 0 {
		// make new map and put in file
		cfg := models.Config{
			Apps:   []models.App{},
			Groups: map[string][]models.App{},
		}

		json.NewEncoder(file).Encode(cfg)
		return true
	}

	return false
}

func ReadStorage(cfg *models.Config) error {

	file, err := os.OpenFile(storagePath, os.O_RDONLY, 0644)
	if err != nil {
		return err
	}

	json.NewDecoder(file).Decode(&cfg)

	return nil
}

func WriteStorage(cfg *models.Config) error {
	file, err := os.OpenFile(storagePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(cfg); err != nil {
		return err
	}
	return nil
}
