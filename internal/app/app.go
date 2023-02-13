package app

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"github.com/VashUber/golang-todo/infrastructure"
)

func Run(cfg *infrastructure.Config) {
	a := app.New()
	w := a.NewWindow("Hello World")

	w.Resize(fyne.NewSize(cfg.Size.W, cfg.Size.H))

	w.SetContent(widget.NewLabel("Hello World!"))
	w.ShowAndRun()
}
