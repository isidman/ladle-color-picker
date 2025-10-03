package theme

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

// Light/Dark mode
type ThemeMode int

const (
	Light ThemeMode = iota //0
	Dark                   //1
)

// Theme colors
var (
	//Light mode
	LightBGR = color.RGBA{245, 224, 172, 255} // #F5E0AC
	LightPR  = color.RGBA{123, 93, 233, 255}  // #7b5de9

	//Dark mode
	DarkBGR = color.RGBA{123, 93, 233, 255}  // #7b5de9
	DarkPR  = color.RGBA{245, 224, 172, 255} // #F5E0AC

	//Grid animation colors (stay the same)
	GridBGR = color.RGBA{245, 224, 172, 255} // #F5E0AC
	GridLN  = color.RGBA{123, 93, 233, 255}  // #7B5DE9
)

// LadleTheme implements fyne's theme interface
type LadleTheme struct {
	mode ThemeMode
}

// New theme
func NewTheme(mode ThemeMode) *LadleTheme {
	return &LadleTheme{mode: mode}
}

// Returning theme colors
func (t *LadleTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch t.mode {
	case Light:
		return t.lightColor(name, variant)
	case Dark:
		return t.darkColor(name, variant)
	default:
		return theme.DefaultTheme().Color(name, variant)
	}
}

func (t *LadleTheme) lightColor(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return LightBGR
	case theme.ColorNameForeground:
		return LightPR
	case theme.ColorNamePrimary:
		return LightPR
	case theme.ColorNameButton, theme.ColorNameInputBackground:
		return color.RGBA{255, 255, 255, 200} //Semi-transparent white
	case theme.ColorNameSeparator:
		return LightPR
	default:
		return theme.DefaultTheme().Color(name, variant)
	}
}

func (t *LadleTheme) darkColor(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return DarkBGR
	case theme.ColorNameForeground:
		return DarkPR
	case theme.ColorNamePrimary:
		return DarkPR
	case theme.ColorNameButton, theme.ColorNameInputBackground:
		return color.RGBA{0, 0, 0, 100} //Semi-transparent black
	case theme.ColorNameSeparator:
		return DarkPR
	default:
		//Invert the default colors for dark theme
		defaultColor := theme.DefaultTheme().Color(name, variant)
		return invertColor(defaultColor)
	}
}

// Theme fonts return
func (t *LadleTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

// Theme sizes return
func (t *LadleTheme) Size(name fyne.ThemeSizeName) float32 {
	switch name {
	case theme.SizeNamePadding:
		return 8
	case theme.SizeNameScrollBar:
		return 12
	case theme.SizeNameSeparatorThickness:
		return 2
	default:
		return theme.DefaultTheme().Size(name)
	}
}

// Icon returns theme icons (Added missing method)
func (t *LadleTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

// Helper for color invertion
func invertColor(c color.Color) color.Color {
	r, g, b, a := c.RGBA()
	return color.RGBA{
		R: uint8(255 - (r >> 8)),
		G: uint8(255 - (g >> 8)),
		B: uint8(255 - (b >> 8)),
		A: uint8(a >> 8),
	}
}
