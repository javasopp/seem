package fyne

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	myApp "seem/app"
)

// Run 初始化fyne的基本结构
func Run() {
	app := app.New()
	myWindow := app.NewWindow("Hello")

	hello := widget.NewLabel("Hello Fyne!")
	content := container.NewHBox(hello)
	myWindow.SetContent(content)

	myWindow.Resize(fyne.NewSize(400, 200)) // 初始窗口大小

	go myApp.Run()

	// 禁用关闭确认，以便点击关闭按钮可以立即关闭窗口
	// 关闭窗口时执行的回调
	myWindow.SetOnClosed(func() {
		// 关闭回调用，处理程序关闭和窗口关闭
		// TODO 处理其他程序的调用

		// 当前窗口的关闭
		app.Quit()
	})
	myWindow.ShowAndRun()
}
