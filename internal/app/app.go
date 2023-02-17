package app

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"github.com/VashUber/golang-todo/infrastructure"
)

func AddTODO(input *widget.Entry, root *widget.Box) func() {
	return func() {
		lbl := widget.NewLabel(input.Text)
		input.SetText("")

		root.Append(lbl)
	}
}

func Run(cfg *infrastructure.Config) {
	a := app.New()
	w := a.NewWindow("Todo")
	w.Resize(fyne.NewSize(cfg.Size.W, cfg.Size.H))

	input := widget.NewEntry()

	content := &widget.Box{
		Children: []fyne.CanvasObject{
			input,
		},
	}

	addTodo := AddTODO(input, content)
	btn := widget.NewButton("Click", addTodo)
	content.Append(btn)

	w.SetContent(content)
	w.ShowAndRun()
}
