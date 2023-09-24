package fyne

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// 初始化fyne的基本结构
func init() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Hello")

	hello := widget.NewLabel("Hello Fyne!")
	myWindow.SetContent(container.NewVBox(
		hello,
	))

	// 禁用关闭确认，以便点击关闭按钮可以立即关闭窗口
	// 关闭窗口时执行的回调
	myWindow.SetOnClosed(func() {
		myApp.Quit()
	})

	myWindow.ShowAndRun()
}
