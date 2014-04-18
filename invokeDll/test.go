package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

func abort(funcname string, err int) {
	panic(funcname + " failed: " + string(err))
}

var (
	kernel32, _        = syscall.LoadLibrary("kernel32.dll")
	getModuleHandle, _ = syscall.GetProcAddress(kernel32, "GetModuleHandleW")

	user32, _     = syscall.LoadLibrary("user32.dll")
	messageBox, _ = syscall.GetProcAddress(user32, "MessageBoxW")
)

const (
	MB_OK                = 0x00000000
	MB_OKCANCEL          = 0x00000001
	MB_ABORTRETRYIGNORE  = 0x00000002
	MB_YESNOCANCEL       = 0x00000003
	MB_YESNO             = 0x00000004
	MB_RETRYCANCEL       = 0x00000005
	MB_CANCELTRYCONTINUE = 0x00000006
	MB_ICONHAND          = 0x00000010
	MB_ICONQUESTION      = 0x00000020
	MB_ICONEXCLAMATION   = 0x00000030
	MB_ICONASTERISK      = 0x00000040
	MB_USERICON          = 0x00000080
	MB_ICONWARNING       = MB_ICONEXCLAMATION
	MB_ICONERROR         = MB_ICONHAND
	MB_ICONINFORMATION   = MB_ICONASTERISK
	MB_ICONSTOP          = MB_ICONHAND

	MB_DEFBUTTON1 = 0x00000000
	MB_DEFBUTTON2 = 0x00000100
	MB_DEFBUTTON3 = 0x00000200
	MB_DEFBUTTON4 = 0x00000300
)

func MessageBox(caption, text string, style uintptr) (result int) {
	ret, _, callErr := syscall.Syscall9(uintptr(messageBox),
		0,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(text))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(caption))),
		style,
		0,
		0,
		0,
		0,
		0,
		0)
	if callErr != 0 {
		abort("Call MessageBox", int(callErr))
	}
	result = int(ret)
	return
}

func GetModuleHandle() (handle uintptr) {
	if ret, _, callErr := syscall.Syscall(uintptr(getModuleHandle), 0, 0, 0, 0); callErr != 0 {
		abort("Call GetModuleHandle", int(callErr))
	} else {
		handle = ret
	}
	return
}

func main() {
	defer syscall.FreeLibrary(kernel32)
	defer syscall.FreeLibrary(user32)

	//fmt.Printf("Retern: %d\n", MessageBox("Done Title", "This test is Done.", MB_YESNOCANCEL))
	num := MessageBox("Done Title", "This test is Done.", MB_YESNOCANCEL)
	if num == 6 {
		fmt.Printf("Return %d\n", num)
	} else {
		fmt.Printf("others")
	}
}

func init() {
	fmt.Print("Starting Up\n")
}
