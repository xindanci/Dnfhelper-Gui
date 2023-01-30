package api

import (
	"fmt"
	"golang.org/x/sys/windows"
	"unsafe"
)

// NtQueryInformationThread information *ThreadBasicInformation
func NtQueryInformationThread(
	handle windows.Handle,
	informationClass uint32,
	information unsafe.Pointer,
	informationLength uint32,
	length *uint32,
) error {
	r0, _, err := procNtQueryInformationThread.Call(uintptr(handle),
		uintptr(informationClass),
		uintptr(information),
		uintptr(informationLength),
		uintptr(unsafe.Pointer(length)))
	if err != windows.ERROR_SUCCESS || windows.NTStatus(r0) != 0 {
		return fmt.Errorf("NtQueryInformationThread %v %v", r0, err)
	}
	return nil
}

func GetExitCodeThread(hThread windows.Handle) (uint32, error) {
	var exitCode uint32
	ret, _, err := procGetExitCodeThread.Call(uintptr(hThread), uintptr(unsafe.Pointer(&exitCode)), 0)
	if ret == 0 || (err != nil && err != windows.ERROR_SUCCESS) {
		panic(err)
	}
	return exitCode, nil
}

func TerminateThread(hThread windows.Handle, exitCode uint32) error {
	ret, _, err := procTerminateThread.Call(uintptr(hThread), uintptr(exitCode))
	if ret == 0 || (err != nil && err != windows.ERROR_SUCCESS) {
		panic(err)
	}
	return nil
}
