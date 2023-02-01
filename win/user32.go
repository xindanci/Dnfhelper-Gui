package win

import (
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

const (
	InputKeyboard = 1
)

func MapVirtualKey(uCode syscall.Handle) (handle syscall.Handle, err error) {
	r0, _, e1 := procMapVirtualKeyA.Call(uintptr(uCode), 0, 0)
	handle = syscall.Handle(r0)
	if handle == 0 || (e1 != nil && e1 != windows.ERROR_SUCCESS) {
		err = e1
	}
	return
}

// GetKeyState beep boop keyboard things
func GetKeyState(vKey int) uint16 {
	ret, _, _ := procGetKeyState.Call(uintptr(vKey))
	return uint16(ret)
}

// SendInput
// pInputs unsafe.Pointer 指向 KeyboardInput 或 HardwareInput 结构的切片
func SendInput(nInputs uint32, pInputs unsafe.Pointer, cbSize uint32) (result uint32, err error) {
	ret, _, e1 := procSendInput.Call(uintptr(nInputs), uintptr(pInputs), uintptr(cbSize))
	result = uint32(ret)
	if result == 0 || (e1 != nil && e1 != windows.ERROR_SUCCESS) {
		err = e1
	}
	return
}

func ToUnicodeEx(key syscall.Handle, scanCode syscall.Handle, keyState *uint16, pwszBuff *uint16) (handle syscall.Handle, err error) {
	r0, _, e1 := procToUnicodeEx.Call(
		uintptr(key),
		uintptr(scanCode),
		uintptr(unsafe.Pointer(keyState)),
		uintptr(unsafe.Pointer(pwszBuff)),
		1,
		0,
	)
	handle = syscall.Handle(r0)
	if handle == 0 || (e1 != nil && e1 != windows.ERROR_SUCCESS) {
		err = e1
	}
	return
}
