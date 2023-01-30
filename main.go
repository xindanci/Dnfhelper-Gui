package main

import (
	_ "embed"
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/xc"
	"os"
	"xctest/gui"
)

//go:embed resource/ui.dll
var dll []byte

func main() {
	_ = os.WriteFile("ui.dll", dll, 0666)
	_ = xc.SetXcguiPath("ui.dll")
	a := app.New(true)

	gui.LoginView(a)

	a.Run()
	a.Exit()
}
