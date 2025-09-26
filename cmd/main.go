package main

import (
	"ladle-color-picker/internal/ui"
	"log"
)

func main() {
	app := ui.NewColorPicker()
	if err := app.Run(); err != nil {
		log.Fatal("Failed to start app:", err)
	}
}
