package main

import (
	"dnfheler/gui"
	_ "embed"
	"fmt"
	"github.com/twgh/xcgui/xc"
	"os"
)

//go:embed resource/ui.dll
var dll []byte

func main() {
	dllDir := fmt.Sprintf("%s\\ui.dll", os.TempDir())
	_ = os.WriteFile(dllDir, dll, 0666)
	_ = xc.SetXcguiPath(dllDir)
	app := gui.NewGui()

	app.LoginView()

	app.Run()
	app.Exit()
}
