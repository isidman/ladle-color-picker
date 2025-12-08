package ui

import "fyne.io/fyne/v2/widget"

// The function below sets up all event handlers
func (app *ColorPicker) setupAllEvents() {
	// Slider events
	app.components.RedSlider.OnChanged = func(value float64) {
		if app.isUpdating {
			return
		}
		app.currentColor.R = uint8(value)
		app.afterColorChange()
	}

	app.components.GreenSlider.OnChanged = func(value float64) {
		if app.isUpdating {
			return
		}
		app.currentColor.G = uint8(value)
		app.afterColorChange()
	}

	app.components.BlueSlider.OnChanged = func(value float64) {
		if app.isUpdating {
			return
		}
		app.currentColor.B = uint8(value)
		app.afterColorChange()
	}
}

func (app *ColorPicker) setupPaletteEvents() {
	for _, btn := range app.components.PresetButtons {
		hex := btn.Text
		btn.OnTapped = func() {
			app.applyColorHex(hex)
		}
	}
}

func (app *ColorPicker) setupThemeEvents() {
	app.themeToggleBtn.OnTapped = func() {
		app.toggleTheme()
	}
}

func (app *ColorPicker) setupExtendedEvents() {
	// Save button event
	app.components.SaveBtn.OnTapped = func() {
		if app.palette.AddSaved(app.currentColor.ToHex()) {
			app.updateSavedColors()
			app.savePalette()
			app.showNotification("Color saved to palette!")
		} else {
			app.showNotification("Color already saved!")
		}
	}

	// Copy button events
	app.components.CopyHexBtn.OnTapped = func() {
		app.copyToClipboard(app.currentColor.ToHex())
	}

	app.components.CopyRGBBtn.OnTapped = func() {
		app.copyToClipboard(app.currentColor.ToRGB())
	}
}

// updateUI updates all UI elements
func (app *ColorPicker) updateUI() {
	app.components.UpdateColorDisplay(app.currentColor)
	app.updateSliders()
}

// updateSliders updates slider position without triggering events
func (app *ColorPicker) updateSliders() {
	app.isUpdating = true
	app.components.RedSlider.SetValue(float64(app.currentColor.R))
	app.components.GreenSlider.SetValue(float64(app.currentColor.G))
	app.components.BlueSlider.SetValue(float64(app.currentColor.B))
	app.isUpdating = false
}

// showNotifications shows a simple notification
func (app *ColorPicker) showNotification(message string) {
	// Console notification for now
	println("📣", message)
}

// updateSavedColors refreshes the saved colors UI
func (app *ColorPicker) updateSavedColors() {
	app.components.RecentBox.Objects = nil
	for _, hex := range app.palette.RecentColors {
		hex := hex
		btn := app.makeColorButton(hex)
		btn.OnTapped = func() { app.applyColorHex(hex) }
		app.components.RecentBox.Add(btn)
	}
	app.components.RecentBox.Refresh()

	app.components.SavedBox.Objects = nil
	for _, hex := range app.palette.SavedColors {
		hex := hex
		btn := app.makeColorButton(hex)
		btn.OnTapped = func() { app.applyColorHex(hex) }
		app.components.SavedBox.Add(btn)
	}
	app.components.SavedBox.Refresh()
}

func (app *ColorPicker) makeColorButton(hex string) *widget.Button {
	return widget.NewButton(hex, nil)
}

func (app *ColorPicker) afterColorChange() {
	app.updateColorDisplay()
	app.palette.AddRecent(app.currentColor.ToHex())
	app.savePalette()
	app.updateSavedColors()
}
