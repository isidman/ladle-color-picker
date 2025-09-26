package storage

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const (
	configDirName   = ".ladle-color-picker"
	paletteFileName = "palette.json"
)

// PaletteData holds the recent and saved colors.
type PaletteData struct {
	Recent []string `json:"recent_colors"`
	Saved  []string `json:"saved_colors"`
}

// Save writes the palette data to the user's config directory.
func Save(data *PaletteData) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	configDir := filepath.Join(homeDir, configDirName)
	if err := os.MkdirAll(configDir, 0o755); err != nil {
		return err
	}
	filePath := filepath.Join(configDir, paletteFileName)
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")
	return encoder.Encode(data)
}

// The below function reads the palette data from user's config directory.
// If the file does not exist, it returns an empty PaletteData.
func Load() (*PaletteData, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	filePath := filepath.Join(homeDir, configDirName, paletteFileName)
	file, err := os.Open(filePath)
	if os.IsNotExist(err) {
		//Return empty data, since there's no file yet.
		return &PaletteData{}, nil
	}
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var data PaletteData
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		return nil, err
	}
	return &data, nil
}
