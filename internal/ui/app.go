package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"ladle-color-picker/internal/color"
	"ladle-color-picker/internal/theme"
)

// ColorPicker represents the main application
type ColorPicker struct {
	app          fyne.App
	window       fyne.Window
	currentColor *color.Color
	palette      *color.Palette
	components   *Components
	currentTheme *theme.LadleTheme
	themeMode    theme.ThemeMode
	animatedBg   *theme.AnimatedBG

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
	ladleTheme := theme.NewTheme(theme.Light)
	myApp.Settings().SetTheme(ladleTheme)

	myWindow := myApp.NewWindow("🥣 Ladle: A color picker")
	myWindow.Resize(fyne.NewSize(500, 700))
	myWindow.CenterOnScreen()

	return &ColorPicker{
		app:          myApp,
		window:       myWindow,
		currentTheme: ladleTheme,
		themeMode:    theme.Light,
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
	app.setupAllEvents()

	//Show and run
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
	app.animatedBg = theme.NewAnimatedBG(content)

	//Setting the animated bg as window content
	app.window.SetContent(app.animatedBg.GetContainer())
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
	mainContent := container.NewVBox(
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
			widget.NewButton(" Copy HEX", func() {
				app.copyToClipboard(app.currentColor.ToHex())
			}),
			widget.NewButton(" Copy RGB", func() {
				app.copyToClipboard(app.currentColor.ToRGB())
			}),
		),
	)

	// Return the content in a rounded card
	return theme.RoundedCard(mainContent, app.currentTheme)
}

func (app *ColorPicker) setupEvents() {
	//Slider events
	app.redSlider.OnChanged = func(value float64) {
		app.currentColor.R = uint8(value)
		app.updateColorDisplay()
	}

	app.greenSlider.OnChanged = func(value float64) {
		app.currentColor.G = uint8(value)
		app.updateColorDisplay()
	}

	app.blueSlider.OnChanged = func(value float64) {
		app.currentColor.B = uint8(value)
		app.updateColorDisplay()
	}

	//Theme toggle event
	app.themeToggleBtn.OnTapped = func() {
		app.toggleTheme()
	}
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
	if app.themeMode == theme.Light {
		app.themeMode == theme.Dark
		app.themeToggleBtn.SetText("☀️ Light Mode")
	} else {
		app.themeMode = theme.Light
		app.themeToggleBtn.SetText("🌙 Dark Mode")
	}

	//Create new theme and apply it
	app.currentTheme = theme.NewTheme(app.themeMode)
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
