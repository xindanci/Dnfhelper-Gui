package gui

import (
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

func LoginView(a *app.App) {
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
		Dashboard()
		return 0
	})

	w.ShowWindow(xcc.SW_SHOW)
}
