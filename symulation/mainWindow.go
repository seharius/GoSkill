package symulation

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func convert() {

}

func Main() fyne.Window {
	w2 := fyne.CurrentApp().NewWindow("Symulator")
	w2.FullScreen()

	var a float64

	money := widget.NewProgressBar()
	money.SetValue(a)
	addMoney := widget.NewButton("addMoney", func() {
		a++
	})
	addMoney.Move(fyne.NewPos(100, 0))

	w2.SetContent(fyne.NewContainer(money))
	return w2
}
