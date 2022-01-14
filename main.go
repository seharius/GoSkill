package main

import (
    "fmt"

	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/app"
)

func main() {

    a := app.New()
    w := a.NewWindow("mój projekt")

    w.SetContent(widget.NewEntry())
    w.ShowAndRun()
    Exited()
}

func Exited(){
    fmt.Println("Exit")
}