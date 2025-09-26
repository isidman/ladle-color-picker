package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"ladle-color-picker/internal/color"
)

// The Components struct holds all the ui components
type Components struct {
	ColorDisplay *widget.Card
	HexLabel     *widget.Label
	RGBLabel     *widget.Label
	HSLLabel     *widget.Label
	RedSlider    *widget.Slider
	GreenSlider  *widget.Slider
	BlueSlider   *widget.Slider
	CopyHexBtn   *widget.Button
	CopyRGBBtn   *widget.Button
	SaveBtn      *widget.Button
	RecentBox    *fyne.Container
	SavedBox     *fyne.Container
}

// New UI Components are created below
func NewComponents() *Components {
	return &Components{
		ColorDisplay: widget.NewCard("Current Color", "", widget.NewLabel("	")),
		HexLabel:     widget.NewLabel("HEX: #ff0000"),
		RGBLabel:     widget.NewLabel("RGB: rgb(255, 0, 0)"),
		HSLLabel:     widget.NewLabel("HSL: hsl(0, 100%, 50%)"),
		RedSlider:    widget.NewSlider(0, 255),
		GreenSlider:  widget.NewSlider(0, 255),
		BlueSlider:   widget.NewSlider(0, 255),
		CopyHexBtn:   widget.NewButton("Copy HEX", nil),
		CopyRGBBtn:   widget.NewButton("Copy RGB", nil),
		SaveBtn:      widget.NewButton("Save Color", nil),
		RecentBox:    container.NewHBox(),
		SavedBox:     container.NewHBox(),
	}
}

// CreateLayout function creates the main application layout
func (c *Components) CreateLayout(currentColor *color.Color, palette *color.Palette) fyne.CanvasObject {
	// Set initial values
	c.RedSlider.SetValue(float64(currentColor.R))
	c.GreenSlider.SetValue(float64(currentColor.G))
	c.BlueSlider.SetValue(float64(currentColor.B))

	return container.NewVBox(
		c.ColorDisplay,
		widget.NewSeparator(),
		c.HexLabel,
		c.RGBLabel,
		c.HSLLabel,
		widget.NewSeparator(),
		widget.NewLabel("🔴 Red:"),
		c.RedSlider,
		widget.NewLabel("🟢 Green:"),
		c.GreenSlider,
		widget.NewLabel("🔵 Blue:"),
		c.BlueSlider,
		widget.NewSeparator(),
		container.NewHBox(c.CopyHexBtn, c.CopyRGBBtn, c.SaveBtn),
		widget.NewSeparator(),
		widget.NewLabel(" Preset Colors:"),
		c.createPresetColors(),
		widget.NewSeparator(),
		widget.NewLabel(" Recent Colors:"),
		c.RecentBox,
		widget.NewSeparator(),
		widget.NewLabel(" Saved Colors:"),
		c.SavedBox,
	)
}

// createPresetColors creates preset color buttons
func (c *Components) createPresetColors() *fyne.Container {
	presetColors := color.GetPresetColors()
	buttons := make([]fyne.CanvasObject, len(presetColors))

	for i, col := range presetColors {
		//Create button for each preset color
		buttons[i] = widget.NewButton(col.ToHex(), nil)
	}

	return container.NewHBox(buttons...)
}

// All color-related UI elements are getting updated here
func (c *Components) UpdateColorDisplay(col *color.Color) {
	c.HexLabel.SetText("HEX: " + col.ToHex())
	c.RGBLabel.SetText("RGB: " + col.ToRGB())
	c.HSLLabel.SetText("HSL: " + col.ToHSL())
	c.ColorDisplay.SetTitle("Current Color: " + col.ToHex())
}
