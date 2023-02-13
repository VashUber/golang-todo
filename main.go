package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello World")

	w.Resize(fyne.NewSize(450, 600))

	w.SetContent(widget.NewLabel("Hello World!"))
	w.ShowAndRun()
}
