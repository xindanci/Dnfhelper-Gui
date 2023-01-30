package api

import (
	"golang.org/x/sys/windows"
)

var (
	kernel32             = windows.NewLazySystemDLL("kernel32.dll")
	procVirtualAllocEx   = kernel32.NewProc("VirtualAllocEx")
	getProcessIDOfThread = kernel32.NewProc("GetProcessIdOfThread")
	procTerminateThread  = kernel32.NewProc("TerminateThread")

	ntdll                        = windows.NewLazySystemDLL("ntdll.dll")
	procNtQueryInformationThread = ntdll.NewProc("NtQueryInformationThread")
	procGetExitCodeThread        = ntdll.NewProc("GetExitCodeThread")

	user32             = windows.NewLazySystemDLL("user32.dll")
	procMapVirtualKeyA = user32.NewProc("MapVirtualKeyA")
	procSendInput      = user32.NewProc("SendInput")
	procGetKeyState    = user32.NewProc("GetKeyState")
	procToUnicodeEx    = user32.NewProc("ToUnicodeEx")
)
