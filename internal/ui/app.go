package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"

	"ladle-color-picker/internal/color"
)

// ColorPicker represents the main application
type ColorPicker struct {
	app          fyne.App
	window       fyne.Window
	currentColor *color.Color
	palette      *color.Palette
	components   *Components
}

// NewColorPicker creates a new color picker application
func NewColorPicker() *ColorPicker {
	myApp := app.New()
	myApp.SetMetadata(&fyne.AppMetadata{
		ID:   "com.example.ladle-color-picker",
		Name: "Ladle Color Picker",
	})

	myWindow := myApp.NewWindow("🥣 Ladle")
	myWindow.Resize(fyne.NewSize(500, 700))
	myWindow.CenterOnScreen()

	return &ColorPicker{
		app:          myApp,
		window:       myWindow,
		currentColor: color.NewColor(255, 0, 0),
		palette:      color.NewPalette(),
		components:   NewComponents(),
	}
}

// Run starts the application
func (app *ColorPicker) Run() error {
	//Load saved palette
	app.palette.Load()

	//Setup UI
	app.setupUI()

	//Setup event handlers
	app.setupEvents()

	//Show and run
	app.window.ShowAndRun()

	//Save palette on exit
	return app.palette.Save()
}

// setupUI creates the user interface
func (app *ColorPicker) setupUI() {
	content := app.components.CreateLayout(app.currentColor, app.palette)
	app.window.SetContent(content)
}
