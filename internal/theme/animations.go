package theme

import (
	"image/color"
	"math"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

//The moving grid backgound creation
type AnimatedBG struct {
	container *container.Border
	gridSize float32
	offset float32
	animationTick *time.Ticker
	gridLines []*canvas.Line
}

//Nw backgound of animated grid was created
func NewABGR(content fyne.CanvasObject) *AnimatedBackground {
	bg: &AnimatedBG {
		gridSize: 50.0,
		offset: 0.0,
		animationTick: time.NewTicker(16 *time.Millisecond), // ~60 FPS 
	}

	//Create RGB rectangle
	BGR := canvas.NewRectangle(GridBackground)

	//Main container
	MC := container.NewBorder(nil, nil, nil, nil, content)

	//
	PC := container.NewPadding(
	container.NewBorder(nil, nil, nil, nil, MC),
	fyne.NewSize(20, 20), //Padding size
}

bg.container = container.NewWithoutLayout (background, paddedCard)

//Begin animation
go bg.animate()

return bg



//GetContainer returns the container with the animated BG
func (bg *AnimatedBackground) GetContainer() *container.Border {
	return bg.container
}

//Grid animation running
func (bg *AnimatedBackground) animate() {
	for range bg.animationTick.C {
		bg.offset += 0.5 //Speed
		if bg.offset >= bg.gridSize {
			bg.offset = 0
		}
		bg.updateGrid()
	}
}

//Grid line position updating
func (bg *AnimatedBackground) updateGrid() {
	if bg.container == nil {
		return
	}

	size := bg.container.Size()
	if size.Width == 0 || size.Height == 0 {
		return
	}

	//Clear existing grid lines
	for _, line := range bg.gridLines {
		bg.container.Objects = removeFromObjects(bg.container.Objects, line)
	}
	bg.gridLines = bg.gridLines[:0]

	//Vertical lines
	for x := -bg.gridSize + bg.offset; x < size.Width+bg.gridSize; x += bg.gridSize {
		line := canvas.NewLine(GridLines)
		line.StrokeWidth = 1
		line.Position1 = fyne.NewPos(x, 0)
		line.Position2 = fyne.NewPos(x, size.Height)
		bg.gridLines = append(bg.gridLines, line)
		bg.container.Objects = append(bg.container.Objects, line)
	}

	//Horizontal lines
	for y := -bg.gridSize + bg.offset; y < size.Height+bg.gridSize; y += bg.gridSize {
		line := canvas.NewLine(GridLines)
		line.StrokeWidth = 1
		line.Position1 = fyne.NewPos(0, y)
		line.Position2 = fyne.NewPos(size.Width, y)
		bg.gridLines = append(bg.gridLines, line)
		bg.container.Objects = append(bg.container.Objects, line)
	}

	bg.container.Refresh()
}

//Stopping the animation
func (bg *AnimatedBackground) Stop() {
	if bg.animationTick != nil {
		bg.animationTick.Stop()
	}
}

//Helper for object removal from slice
func removeFromObjects(objects []fyne.CanvasObject, target fyne.CanvasObject) []fyne.CanvasObject {
	for i, obj := range objects {
		if obj == target {
			return append(objects[:i], objects[i+1]...)
		}
	}
	return objects
}

//Rounded card container creation
func RoundedCard(content fyne.CanvasObject, theme *LadleTheme) *fyne.Container {
	//BG creation with rounded corners
	bg := canvas.NewRectangle(theme.Color(theme.ColorNameBackground, theme.VariantStandard))

	//Container with padding for rounded appearance
	card := container.NewBorder(
		nil, nil, nil, nil,
		container.NewPadding(content, fyne.NewSize(20, 20)),
	)

	return container.NewWithoutLayout(bg, card)
}