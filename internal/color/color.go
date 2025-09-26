package color

import (
	"fmt"
	"image/color"
	"math"
)

// Color represents an RGB color with conversion methods
type Color struct {
	R, G, B uint8
}

// NewColor creates a new color from RGB values
func NewColor(r, g, b uint8) *Color {
	return &Color{R: r, G: g, B: b}
}

// NewColorHex creates a color from a hex string like "#ff0000"
func NewColorHex(hex string) (*Color, error) {
	if len(hex) != 7 || hex[0] != '#' {
		return nil, fmt.Errorf("invalid hex format: %s", hex)
	}

	var r, g, b uint8
	if _, err := fmt.Sscanf(hex[1:], "%02x%02x%02x", &r, &g, &b); err != nil {
		return nil, err
	}

	return &Color{R: r, G: g, B: b}, nil
}

// ToHex returns the color as a hex string like "#ff0000"
func (c *Color) ToHex() string {
	return fmt.Sprintf("#%02x%02x%02x", c.R, c.G, c.B)
}

// ToRGB returns RGB string like "rgb(255, 0, 0)"
func (c *Color) ToRGB() string {
	return fmt.Sprintf("rgb(%d, %d, %d)", c.R, c.G, c.B)
}

// ToHSL return HSL string like "hsl(0, 100%, 50%)"
func (c *Color) ToHSL() string {
	h, s, l := c.rgbToHsl()
	return fmt.Sprintf("hsl(%.0f, %.0f%%, %.0f%%)", h, s*100, l*100)
}

// ToFyneColor converts to Fyne's color format
func (c *Color) ToFyneColor() color.RGBA {
	return color.RGBA{R: c.R, G: c.G, B: c.B, A: 255}
}

// rgbToHsl converts RGB to HSL values
func (c *Color) rgbToHsl() (float64, float64, float64) {
	r, g, b := float64(c.R)/255, float64(c.G)/255, float64(c.B)/255

	max := math.Max(math.Max(r, g), b)
	min := math.Min(math.Min(r, g), b)

	l := (max + min) / 2
	var h, s float64

	if max == min {
		h, s = 0, 0 //achromatic
	} else {
		d := max - min
		if l > 0.5 {
			s = d / (2 - max - min)
		} else {
			s = d / (max + min)
		}

		switch max {
		case r:
			h = (g - b) / d
			if g < b {
				h += 6
			}
		case g:
			h = (b-r)/d + 2
		case b:
			h = (r-g)/d + 4
		}
		h /= 6
	}

	return h * 360, s, l
}

// The GetPresetColors returns common preset colors
func GetPresetColors() []*Color {
	return []*Color{
		NewColor(255, 0, 0),     //Red
		NewColor(0, 255, 0),     //Green
		NewColor(0, 0, 255),     //Blue
		NewColor(255, 255, 0),   //Yellow
		NewColor(255, 0, 255),   //Magenta
		NewColor(0, 255, 255),   //Cyan
		NewColor(255, 255, 255), //White
		NewColor(0, 0, 0),       //Black
	}
}
