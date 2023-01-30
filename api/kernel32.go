package api

import (
	"golang.org/x/sys/windows"
	"log"
	"runtime/debug"
	"strconv"
	"unsafe"
)

func OpenProcess(desiredAccess uint32, inheritHandle bool, processId uint32) (uintptr, error) {
	process, err := windows.OpenProcess(desiredAccess, inheritHandle, processId)
	if err != nil && err != windows.ERROR_SUCCESS {
		panic(err)
	}

	return uintptr(process), nil
}

func CloseHandle(object uintptr) bool {
	err := windows.CloseHandle(windows.Handle(object))
	if err != nil && err != windows.ERROR_SUCCESS {
		panic(err)
	}
	return true
}

// WriteProcessMemory 写进程内存
func WriteProcessMemory(hProcess uintptr, lpBaseAddress uint64, data []byte, size uint) error {
	err := windows.WriteProcessMemory(windows.Handle(hProcess), uintptr(lpBaseAddress), &data[0], uintptr(size), nil)
	if err == windows.ERROR_PARTIAL_COPY {
		err = windows.ERROR_SUCCESS
	}
	if err != nil && err != windows.ERROR_SUCCESS {
		log.Printf("写进程内存错误,地址:%s 值:%+v 错误：%+v 调用栈：%s \n", strconv.FormatUint(lpBaseAddress, 16), data, err, debug.Stack())
	}
	return nil
}

// ReadProcessMemory 读进程内存
func ReadProcessMemory(hProcess uintptr, lpBaseAddress uint64, size uint32) ([]byte, error) {
	data := make([]byte, size)
	err := windows.ReadProcessMemory(windows.Handle(hProcess), uintptr(lpBaseAddress), &data[0], uintptr(size), nil)
	if err == windows.ERROR_PARTIAL_COPY {
		err = windows.ERROR_SUCCESS
	}
	if err != nil && err != windows.ERROR_SUCCESS {
		log.Printf("读进程内存错误,地址:%s,错误：%+v 调用栈：%s \n", strconv.FormatUint(lpBaseAddress, 16), err, debug.Stack())
	}
	return data, nil
}

// DeviceIoControl Ico驱动通讯
func DeviceIoControl(handle uintptr, ioControlCode uint32, inBuffer *byte, inBufferSize uint32, outBuffer *byte, outBufferSize uint32, bytesReturned *uint32, overlapped *windows.Overlapped) error {
	err := windows.DeviceIoControl(
		windows.Handle(handle),
		ioControlCode,
		inBuffer,
		inBufferSize,
		outBuffer,
		outBufferSize,
		bytesReturned,
		overlapped,
	)
	if err != nil && err != windows.ERROR_SUCCESS {
		panic(err)
	}
	return nil
}

func CreateFile(fileName *uint16, desiredAccess uint32, shareMode uint32, securityAttributes *windows.SecurityAttributes, creationDisposition uint32, flagsAndAttributes uint32, templateFile uintptr) (uintptr, error) {
	file, err := windows.CreateFile(fileName, desiredAccess, shareMode, securityAttributes, creationDisposition, flagsAndAttributes, windows.Handle(templateFile))
	if err != nil && err != windows.ERROR_SUCCESS {
		panic(err)
	}
	return uintptr(file), nil
}

func DeleteFile(fileName string) error {
	err := windows.DeleteFile(windows.StringToUTF16Ptr(fileName))
	if err != nil && err != windows.ERROR_SUCCESS {
		panic(err)
	}

	return nil
}

// VirtualAllocEx 申请内存
func VirtualAllocEx(hProcess uintptr, lpBaseAddress uint64, size uint, flAllocationType, flProtect uint32) (uintptr, error) {
	res, _, err := procVirtualAllocEx.Call(hProcess, uintptr(lpBaseAddress), uintptr(size), uintptr(flAllocationType), uintptr(flProtect))
	if res == 0 || (err != nil && err != windows.ERROR_SUCCESS) {
		panic(err)
	}
	return res, nil
}

// EnableDebugPrivilege 提权
func EnableDebugPrivilege(fEnable bool) bool {
	var hToken windows.Token

	var tp windows.Tokenprivileges
	var result = false
	var luid windows.LUID

	err := windows.OpenProcessToken(windows.CurrentProcess(), windows.TOKEN_ADJUST_PRIVILEGES|windows.TOKEN_QUERY, &hToken)
	if err != nil {
		return false
	}

	tp.PrivilegeCount = 1
	err = windows.LookupPrivilegeValue(nil, windows.StringToUTF16Ptr("SeDebugPrivilege"), &luid)
	if err != nil {
		return false
	}

	tp.Privileges[0].Luid = luid
	if fEnable {
		tp.Privileges[0].Attributes = windows.SE_PRIVILEGE_ENABLED
	} else {
		tp.Privileges[0].Attributes = 0
	}

	err = windows.AdjustTokenPrivileges(hToken, false, &tp, uint32(unsafe.Sizeof(tp)), nil, nil)
	if err != nil {
		return false
	}

	result = windows.GetLastError() == windows.ERROR_SUCCESS

	return result
}

func GetPIDFromThread(handle windows.Handle) (uint32, error) {
	ret, _, err := getProcessIDOfThread.Call(uintptr(handle))
	if ret == 0 || (err != nil && err != windows.ERROR_SUCCESS) {
		panic(err)
	}
	return uint32(ret), nil
}
