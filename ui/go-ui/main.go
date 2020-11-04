package main

/*
@Author: by LH
@date:  2020/9/11
@function:
*/
import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"time"
)

func main() {
	a := app.New()
	w := a.NewWindow("网易云")

	hello := widget.NewLabel("Hello Fyne!")
	w.SetContent(widget.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome :)")
			a.SendNotification(&fyne.Notification{
				Title:   "网易云",
				Content: fmt.Sprintf("现在时间%s", time.Now().String()),
			})
		}),
	))

	w.ShowAndRun()
}
