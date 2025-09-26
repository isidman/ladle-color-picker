package color

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// Palette function manages saved and recent colors
type Palette struct {
	RecentColors []string `json:"recent_colors"`
	SavedColors  []string `json:"saved_colors"`
	MaxRecent    int      `json:"-"`
	MaxSaved     int      `json:"-"`
}

// NewPalette creates a new palette
func NewPalette() *Palette {
	return &Palette{
		RecentColors: make([]string, 0),
		SavedColors:  make([]string, 0),
		MaxRecent:    8,
		MaxSaved:     16,
	}
}

// AddRecent adds a color to recent colors
func (p *Palette) AddRecent(hex string) {
	//Remove if the color already exists
	for i, color := range p.RecentColors {
		if color == hex {
			p.RecentColors = append(p.RecentColors[:i], p.RecentColors[i+1:]...)
			break
		}
	}

	//Add to the beginning
	p.RecentColors = append([]string{hex}, p.RecentColors...)

	//Limit
	if len(p.RecentColors) > p.MaxRecent {
		p.RecentColors = p.RecentColors[:p.MaxRecent]
	}
}

// AddSaved adds a color to saved colors
func (p *Palette) AddSaved(hex string) bool {
	//Checking for existence
	for _, color := range p.SavedColors {
		if color == hex {
			return false //Already saved
		}
	}

	//Add new color
	p.SavedColors = append(p.SavedColors, hex)

	//Limit
	if len(p.SavedColors) > p.MaxSaved {
		p.SavedColors = p.SavedColors[1:] //Remove oldest
	}

	return true
}

// RemoveSaved removes a color from saved colors
func (p *Palette) RemoveSaved(hex string) {
	for i, color := range p.SavedColors {
		if color == hex {
			p.SavedColors = append(p.SavedColors[:i], p.SavedColors[i+1:]...)
			break
		}
	}
}

// Save palette to file
func (p *Palette) Save() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	configDir := filepath.Join(homeDir, ".ladle-color-picker")
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return err
	}

	file := filepath.Join(configDir, "palette.json")
	data, err := json.MarshalIndent(p, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(file, data, 0644)
}

// Load palette from file
func (p *Palette) Load() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	file := filepath.Join(homeDir, ".ladle-color-picker", "palette.json")
	data, err := os.ReadFile(file)
	if err != nil {
		return err //File doesn't exist yet
	}

	return json.Unmarshal(data, p)
}
