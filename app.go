package main

import (
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/widget/material"
)

type Context struct {
	Ops *op.Ops
}

func main() {
	go func() {
		w := app.NewWindow(app.Title("Console counter"))
		err := run(w)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func run(w *app.Window) error {

	theme := material.NewTheme()
	var op op.Ops
	for {
		e := <-w.Events()

		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			{

				graphicsCtx := layout.NewContext(&op, e)

				// Define an large label with an appropriate text:

				title := material.H6(theme, "Moses")

				// Change the color of the label.
				title.Color = color.NRGBA{R: 127, G: 0, B: 0, A: 255}

				// Change the position of the label.
				title.Alignment = text.Middle

				// Draw the label to the graphics context.
				title.Layout(graphicsCtx)

				// Pass the drawing operations to the GPU.
				e.Frame(graphicsCtx.Ops)

			}
		}
	}
}
