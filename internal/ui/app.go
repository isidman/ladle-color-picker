package ui

import (
	"fmt"

	"ladle-color-picker/internal/color"
	ladleTheme "ladle-color-picker/internal/theme"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	fyneTheme "fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// ColorPicker represents the main application
type ColorPicker struct {
	app          fyne.App
	window       fyne.Window
	currentColor *color.Color
	palette      *color.Palette
	components   *Components
	currentTheme *ladleTheme.LadleTheme
	themeMode    ladleTheme.ThemeMode
	animatedBg   *ladleTheme.AnimatedBG

	//UI Parts
	colorDisplay   *canvas.Rectangle
	hexLabel       *widget.Label
	rgbLabel       *widget.Label
	hslLabel       *widget.Label
	redSlider      *widget.Slider
	greenSlider    *widget.Slider
	blueSlider     *widget.Slider
	themeToggleBtn *widget.Button
}

// NewColorPicker creates a new color picker application
func NewColorPicker() *ColorPicker {
	myApp := app.New()

	//Light mode
	ladleTheme := ladleTheme.NewTheme(ladleTheme.Light)
	myApp.Settings().SetTheme(ladleTheme)

	myWindow := myApp.NewWindow("🥣 Ladle: A color picker")
	myWindow.Resize(fyne.NewSize(500, 700))
	myWindow.CenterOnScreen()

	return &ColorPicker{
		app:          myApp,
		window:       myWindow,
		currentTheme: ladleTheme,
		themeMode:    0,
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

	app.setupThemeEvents()

	app.setupExtendedEvents()

	//Setup event handlers
	app.setupAllEvents()

	//Called to refresh the UI initially
	app.updateUI()

	app.window.ShowAndRun()

	//Cleanup
	if app.animatedBg != nil {
		app.animatedBg.Stop()
	}

	return nil
}

// setupUI creates the user interface
func (app *ColorPicker) setupUI() {

	//UI Part creation
	app.createComponents()

	//Main content layout creation
	content := app.createMainLayout()

	//Creating animated BG with content
	app.animatedBg = ladleTheme.NewAnimatedBG(content)

	variant := fyneTheme.VariantLight
	card := ladleTheme.RoundedCard(app.animatedBg.GetContainer(), app.currentTheme, variant)
	//Explicit passing of current theme and variant
	app.window.SetContent(card)
	app.updateColorDisplay()
}

func (app *ColorPicker) createComponents() {
	//Color display
	app.colorDisplay = canvas.NewRectangle(app.currentColor.ToFyneColor())
	app.colorDisplay.Resize(fyne.NewSize(200, 100))

	//Labels
	app.hexLabel = widget.NewLabel("HEX: #ff0000")
	app.rgbLabel = widget.NewLabel("RGB: rgb(255, 0, 0)")
	app.hslLabel = widget.NewLabel("HSL: hsl(0, 100%, 50%)")

	//Sliders
	app.redSlider = widget.NewSlider(0, 255)
	app.redSlider.SetValue(255)
	app.greenSlider = widget.NewSlider(0, 255)
	app.blueSlider = widget.NewSlider(0, 255)

	//Theme toggle button
	app.themeToggleBtn = widget.NewButton("🌙 Dark Mode", nil)
}

func (app *ColorPicker) createMainLayout() fyne.CanvasObject {
	// Create rounded card for the main interface
	return container.NewVBox(
		//Header with theme toggle
		container.NewBorder(nil, nil, widget.NewLabel("Ladle: A color picker"), app.themeToggleBtn),
		widget.NewSeparator(),

		//Color display
		container.NewCenter(app.colorDisplay),

		//Color information
		app.hexLabel,
		app.rgbLabel,
		app.hslLabel,
		widget.NewSeparator(),

		//RGB sliders
		widget.NewLabel(" Red"),
		app.redSlider,
		widget.NewLabel(" Green"),
		app.greenSlider,
		widget.NewLabel(" Blue"),
		app.blueSlider,
		widget.NewSeparator(),

		//Action Buttons
		container.NewHBox(
			widget.NewButton("Copy HEX", func() {
				app.copyToClipboard(app.currentColor.ToHex())
			}),
			widget.NewButton("Copy RGB", func() {
				app.copyToClipboard(app.currentColor.ToRGB())
			}),
		),
	)
}

func (app *ColorPicker) updateColorDisplay() {
	app.colorDisplay.FillColor = app.currentColor.ToFyneColor()
	app.colorDisplay.Refresh()

	//Update labels
	app.hexLabel.SetText("HEX: " + app.currentColor.ToHex())
	app.rgbLabel.SetText("RGB: " + app.currentColor.ToRGB())
	app.hslLabel.SetText("HSL: " + app.currentColor.ToHSL())
}

func (app *ColorPicker) toggleTheme() {
	//Switch theme mode
	if app.themeMode == 0 {
		app.themeMode = 1
		app.themeToggleBtn.SetText("☀️ Light Mode")
	} else {
		app.themeMode = 0
		app.themeToggleBtn.SetText("🌙 Dark Mode")
	}

	//Create new theme and apply it
	app.currentTheme = ladleTheme.NewTheme(app.themeMode)
	app.app.Settings().SetTheme(app.currentTheme)

	//Refresh the UI
	app.window.Content().Refresh()
}

func (app *ColorPicker) copyToClipboard(text string) {
	if clipboard := app.app.Driver().AllWindows()[0].Clipboard(); clipboard != nil {
		clipboard.SetContent(text)
		fmt.Printf("Copied: %s\n", text)
	}
}
