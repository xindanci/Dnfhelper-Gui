package main

import (
	_ "embed"
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
	"os"
	"time"
	"xctest/helper"
)

//go:embed ui.dll
var dll []byte

var svgIcon = `<svg t="1674984352573" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="10315" width="16" height="16"><path d="M901.957085 349.126786c-60.072164-87.975001-153.76426-100.09183-187.170868-101.49977-79.698013-8.106329-155.59885 46.931378-196.0025 46.931377-40.40365 0-102.779718-45.822091-168.86763-44.627473-86.908379 1.279947-166.990375 50.515229-211.788509 128.421315-90.32157 156.665472-23.12437 388.762468 64.850631 515.818508 43.048873 62.248073 94.332069 132.133161 161.6146 129.615933 64.850631-2.559893 89.425607-41.982251 167.673013-41.982251 78.418066 0 100.433149 41.982251 169.03829 40.702304 69.799758-1.279947 114.000583-63.400025 156.665473-125.818758 49.405941-72.188992 69.714429-141.98875 70.909045-145.572601-1.578601-0.725303-135.973001-52.221824-137.380942-207.095371-1.279947-129.573268 105.68093-191.778676 110.502062-194.893213zM715.852839 0.042665c-51.496521 2.133244-113.829924 34.302571-150.820382 77.479438-33.107954 38.3984-58.706887 99.622516-50.899213 158.414733 57.51227 4.479813 112.720637-29.182784 148.473814-72.530311 35.710512-43.176868 59.816174-103.377026 53.245781-163.36386z" fill="#1afa29" opacity=".65" p-id="10316"></path></svg>`

var appVersion = "v1.0.0"

func main() {
	_ = os.WriteFile("ui.dll", dll, 0666)
	_ = xc.SetXcguiPath("ui.dll")
	a := app.New(true)

	loginView(a)

	a.Run()
	a.Exit()
}

func loginView(a *app.App) {
	w := window.New(0, 0, 350, 200, "gopher助手", 0, xcc.Window_Style_Modal)
	// 设置窗口图标
	w.SetIcon(xc.XImage_LoadSvgStringW(svgIcon))
	// 禁止改变大小
	w.EnableDragBorder(false)
	// 设置边框
	w.SetBorderSize(0, 30, 0, 0)

	widget.NewShapeText(40, 70, 100, 30, "卡号:", w.Handle)
	card := widget.NewEdit(80, 70, 200, 30, w.Handle)
	card.SetText("admin")

	loginBut := widget.NewButton(135, 122, 86, 26, "登录", w.Handle)
	loginBut.Event_BnClick(func(pbHandled *bool) int {
		if card.GetTextEx() != "admin" {
			a.Alert("信息:", "卡密错误")
			return 0
		}
		w.CloseWindow()
		mainView()
		return 0
	})

	w.ShowWindow(xcc.SW_SHOW)
}

func mainView() {
	w := window.New(0, 0, 510, 300, "当前时间："+time.Now().Format("2006-01-02 15:04:05"), 0, xcc.Window_Style_Modal)
	// 设置窗口图标
	w.SetIcon(xc.XImage_LoadSvgStringW(svgIcon))
	// 禁止改变大小
	w.EnableDragBorder(false)
	// 设置边框
	w.SetBorderSize(0, 30, 0, 0)
	// 动态设置软件标题
	go func() {
		for {
			time.Sleep(time.Second)
			w.SetTitle("当前时间：" + time.Now().Format("2006-01-02 15:04:05"))
			w.Redraw(false)
		}
	}()

	widget.NewShapeText(208, 30, 60, 30, "软件日志", w.Handle)

	widget.NewShapeText(1, 266, 60, 30, "运行时间:", w.Handle)
	runTime := widget.NewShapeText(56, 266, 60, 30, "00:00:00", w.Handle)
	go func() {
		for {
			time.Sleep(time.Second)
			runTime.SetText(helper.GetAppRunTime())
		}
	}()

	widget.NewShapeText(420, 266, 60, 30, "版本号:", w.Handle)
	widget.NewShapeText(464, 266, 60, 30, appVersion, w.Handle)

	editContent := widget.NewEdit(1, 54, 509, 210, w.Handle)
	editContent.EnableMultiLine(true)
	editContent.AutoScroll()
	editContent.EnableReadOnly(true)

	go func() {
		for {
			time.Sleep(time.Second)
			editContent.InsertText(0, 0, time.Now().Format("2006-01-02 15:04:05")+"\n")
			editContent.Redraw(false)
		}
	}()

	w.ShowWindow(xcc.SW_SHOW)
}
