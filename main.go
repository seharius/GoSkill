package main

import (
	"fmt"
	"goskill/symulation"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("ULTORON")
	w.Resize(fyne.NewSize(500, 400))

	btn := widget.NewButton("", func() {
		symulation.Main().Show()
	})

	w.SetContent(btn)
	w.ShowAndRun()
	Exited()
}

func Exited() {
	fmt.Println("Exit")
}
