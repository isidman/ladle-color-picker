package ui

import (
	"errors"
	"fmt"
	"os"

	"ladle-color-picker/internal/color"
	ladleTheme "ladle-color-picker/internal/theme"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
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
	isUpdating   bool

	themeToggleBtn *widget.Button
}

// NewColorPicker creates a new color picker application
func NewColorPicker() *ColorPicker {
	myApp := app.New()

	// Light mode
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
	// Load saved palette
	if err := app.palette.Load(); err != nil && !errors.Is(err, os.ErrNotExist) {
		fmt.Printf("could not load palette: %v\n", err)
	}

	// Setup UI
	app.setupUI()

	app.setupThemeEvents()

	app.setupExtendedEvents()

	// Setup event handlers
	app.setupAllEvents()
	app.setupPaletteEvents()

	// Called to refresh the UI initially
	app.updateUI()
	app.updateSavedColors()

	app.window.ShowAndRun()

	// Cleanup
	if app.animatedBg != nil {
		app.animatedBg.Stop()
	}

	return nil
}

// setupUI creates the user interface
func (app *ColorPicker) setupUI() {

	// UI Part creation
	app.createComponents()

	// Main content layout creation
	content := app.createMainLayout()

	// Creating animated BG with content
	app.animatedBg = ladleTheme.NewAnimatedBG(content)

	variant := fyneTheme.VariantLight
	card := ladleTheme.RoundedCard(app.animatedBg.GetContainer(), app.currentTheme, variant)
	// Explicit passing of current theme and variant
	app.window.SetContent(card)
	app.updateColorDisplay()
}

func (app *ColorPicker) createComponents() {
	app.themeToggleBtn = widget.NewButton("🌙 Dark Mode", nil)
}

func (app *ColorPicker) createMainLayout() fyne.CanvasObject {
	header := container.NewBorder(nil, nil, widget.NewLabel("Ladle: A color picker"), app.themeToggleBtn)

	return container.NewVBox(
		header,
		widget.NewSeparator(),
		app.components.CreateLayout(app.currentColor, app.palette),
	)
}

func (app *ColorPicker) updateColorDisplay() {
	app.components.UpdateColorDisplay(app.currentColor)
}

func (app *ColorPicker) toggleTheme() {
	// Switch theme mode
	if app.themeMode == 0 {
		app.themeMode = 1
		app.themeToggleBtn.SetText("☀️ Light Mode")
	} else {
		app.themeMode = 0
		app.themeToggleBtn.SetText("🌙 Dark Mode")
	}

	// Create new theme and apply it
	app.currentTheme = ladleTheme.NewTheme(app.themeMode)
	app.app.Settings().SetTheme(app.currentTheme)

	// Refresh the UI
	app.window.Content().Refresh()
}

func (app *ColorPicker) copyToClipboard(text string) {
	windows := app.app.Driver().AllWindows()
	if len(windows) == 0 {
		return
	}

	if clipboard := windows[0].Clipboard(); clipboard != nil {
		clipboard.SetContent(text)
		fmt.Printf("Copied: %s\n", text)
	}
}

func (app *ColorPicker) applyColorHex(hex string) {
	col, err := color.NewColorHex(hex)
	if err != nil {
		return
	}

	app.currentColor = col
	app.palette.AddRecent(hex)
	app.updateUI()
	app.savePalette()
	app.updateSavedColors()
}

func (app *ColorPicker) savePalette() {
	if err := app.palette.Save(); err != nil {
		fmt.Printf("could not save palette: %v\n", err)
	}
}
