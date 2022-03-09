package matrix

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func Przyciski() []fyne.CanvasObject {

	var X float32 = 500
	var Y float32 = 400

	var A string
	entry := widget.NewEntry()
	entry.SetPlaceHolder("")
	entry.Resize(fyne.NewSize(X, 50))

	btn0 := widget.NewButton("0", func() {
		A += "0"
		entry.SetText(A)
	})
	btn0.Resize(fyne.NewSize(75, 75))
	btn0.Move(fyne.NewPos(0, Y-10))

	btn1 := widget.NewButton("1", func() {
		A += "1"
		entry.SetText(A)
	})
	btn1.Resize(fyne.NewSize(75, 75))
	btn1.Move(fyne.NewPos(100, 50))

	btn2 := widget.NewButton("2", func() {
		A += "2"
		entry.SetText(A)
	})
	btn2.Resize(fyne.NewSize(75, 75))
	btn2.Move(fyne.NewPos(200, 200))

	btn3 := widget.NewButton("3", func() {
		A += "3"
		entry.SetText(A)
	})
	btn3.Resize(fyne.NewSize(75, 75))
	btn3.Move(fyne.NewPos(200, 200))

	btn4 := widget.NewButton("4", func() {
		A += "4"
		entry.SetText(A)
	})
	btn4.Resize(fyne.NewSize(75, 75))
	btn4.Move(fyne.NewPos(200, 200))

	btn5 := widget.NewButton("5", func() {
		A += "5"
		entry.SetText(A)
	})
	btn5.Resize(fyne.NewSize(75, 75))
	btn5.Move(fyne.NewPos(200, 200))

	btn6 := widget.NewButton("6", func() {
		A += "6"
		entry.SetText(A)
	})
	btn6.Resize(fyne.NewSize(75, 75))
	btn6.Move(fyne.NewPos(200, 200))

	btn7 := widget.NewButton("7", func() {
		A += "7"
		entry.SetText(A)
	})
	btn7.Resize(fyne.NewSize(75, 75))
	btn7.Move(fyne.NewPos(200, 200))

	btn8 := widget.NewButton("8", func() {
		A += "8"
		entry.SetText(A)
	})
	btn8.Resize(fyne.NewSize(75, 75))
	btn8.Move(fyne.NewPos(200, 200))

	btn9 := widget.NewButton("9", func() {
		A += "9"
		entry.SetText(A)
	})
	btn9.Resize(fyne.NewSize(75, 75))
	btn9.Move(fyne.NewPos(200, 200))

	var Tbtn []fyne.CanvasObject
	Tbtn = append(Tbtn,
		entry,
		btn0,
		btn1,
		btn2,
		btn3,
		btn4,
		btn5,
		btn6,
		btn7,
		btn8,
		btn9)

	return Tbtn

}
