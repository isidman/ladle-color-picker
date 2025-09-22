# ladle-color-picker
A desktop color picker built with Go and Fyne


## Features
-  Interactive RGB sliders
-  Preset colors
-  Save favorite colors
-  Recent colors history
-  Copy colors in HEX, RGB, HSL formats
-  Persistent storage

## Installation

1. Install Go (1.21 or later)
2. Clone this repository
3. Run: `go mod tidy`
4. Run: `go run cmd/main.go`

## Building
`go build -o ladle-color-picker cmd/main.go`

## Usage

- Use the RGB sliders to pick colors
- Click preset colors for quick selection
- Save colors you like with the "Save Color" button
- Recent colors appear automatically
- Click any color to copy to clipboard
