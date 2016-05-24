package main

import (
	"math/rand"
	"time"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gtk"
)

type Actor interface {
	act(*Grid) bool
	getX() int
	setX(int)
	getY() int
	setY(int)
	draw(float64, *cairo.Context)
	setEnergy(int)
	getEnergy() int
}

func main() {
	gtk.Init(nil)

	// Handle constructing the plants and grid
	rand.Seed(time.Now().Unix())

	var pGrid Grid
	for i := 0; i < 10; i++ {
		var plant Plant
		plant.create()
		pGrid.append(&plant, plant.x, plant.y)
	}

	// gui boilerplate
	win, _ := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	da, _ := gtk.DrawingAreaNew()
	win.Add(da)
	win.SetTitle("Random Squares")
	win.Connect("destroy", gtk.MainQuit)
	win.SetPosition(gtk.WIN_POS_CENTER)
	win.SetDefaultSize(640, 480)
	win.ShowAll()

	// Data
	unitSize := 10.0

	// Event handlers
	da.Connect("draw", func(da *gtk.DrawingArea, cr *cairo.Context) {
		// Clear the background to one color
		cr.SetSourceRGB(.1, .1, .1)
		cr.Paint()

		// Plant Loop
		actors := pGrid.actors()
		for i := range actors {
			actors[i].act(&pGrid)
		}

		for _, v := range actors {
			v.draw(unitSize, cr)
		}
		win.QueueDraw()
	})
	gtk.Main()
}
