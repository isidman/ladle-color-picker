package main

import (
	"log"

	"ladle-color-picker/internal/ui"
)

func main() {
	app := ui.NewColorPickerApp()
	if err := app.Run(); err != nil {
		log.Fatal("Failed to start app:", err)
	}
}
