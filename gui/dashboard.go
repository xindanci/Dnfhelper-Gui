package gui

import (
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
	"time"
	"xctest/helper"
)

func Dashboard() {
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
