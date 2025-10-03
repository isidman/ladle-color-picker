package ui

// The function below sets up all event handlers
func (app *ColorPicker) setupAllEvents() {
	//Slider events
	app.components.RedSlider.OnChanged = func(value float64) {
		app.currentColor.R = uint8(value)
		app.updateColorDisplay()
		app.palette.AddRecent(app.currentColor.ToHex())
	}

	app.components.GreenSlider.OnChanged = func(value float64) {
		app.currentColor.G = uint8(value)
		app.updateColorDisplay()
		app.palette.AddRecent(app.currentColor.ToHex())
	}

	app.components.BlueSlider.OnChanged = func(value float64) {
		app.currentColor.B = uint8(value)
		app.updateColorDisplay()
		app.palette.AddRecent(app.currentColor.ToHex())
	}
}

func (app *ColorPicker) setupThemeEvents() {
	app.themeToggleBtn.OnTapped = func() {
		app.toggleTheme()
	}
}

func (app *ColorPicker) setupExtendedEvents() {
	//Save button event
	app.components.SaveBtn.OnTapped = func() {
		if app.palette.AddSaved(app.currentColor.ToHex()) {
			app.updateSavedColors()
			app.showNotification("Color saved to palette!")
		} else {
			app.showNotification("Color already saved!")
		}
	}

	//Copy button events
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
	app.components.RedSlider.SetValue(float64(app.currentColor.R))
	app.components.GreenSlider.SetValue(float64(app.currentColor.G))
	app.components.BlueSlider.SetValue(float64(app.currentColor.B))
}

// showNotifications shows a simple notification
func (app *ColorPicker) showNotification(message string) {
	//Console notification for now
	println("📣", message)
}

// updateSavedColors refreshes the saved colors UI
func (app *ColorPicker) updateSavedColors() {
	//Rebuild saved colors container
}
