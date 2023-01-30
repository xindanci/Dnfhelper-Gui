package structure

import "golang.org/x/sys/windows"

type ThreadBasicInformation struct {
	ExitStatus     windows.NTStatus
	TebBaseAddress uint64 // 基础地址
	ClientId       struct {
		UniqueProcess uint64 // 进程ID
		UniqueThread  uint64 // 线程ID
	}
	AffinityMask uint64
	Priority     int32 // 优先权
	BasePriority int32 // 基本优先权
}
