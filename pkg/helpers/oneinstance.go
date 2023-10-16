package helpers

import (
	"fmt"
	"syscall"
	"unsafe"
)

var (
	modkernel32      = syscall.NewLazyDLL("kernel32.dll")
	pCreateSemaphore = modkernel32.NewProc("CreateSemaphoreA")
	pGetLastError    = modkernel32.NewProc("GetLastError")
)

func GetLastError() uint32 {
	ret, _, _ := pGetLastError.Call()
	return uint32(ret)
}

func GetInstance(ident string) error {
	semaphoreName := uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(`Global\` + ident)))
	_, _, _ =
		pCreateSemaphore.Call(
			0,
			uintptr(1),
			uintptr(1),
			semaphoreName,
		)

	fmt.Println(GetLastError())
	fmt.Println(syscall.GetLastError())
	if GetLastError() != 0 {
		return syscall.GetLastError()
	}
	return nil
}
