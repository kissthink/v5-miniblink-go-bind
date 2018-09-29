package main

import (
	"syscall"
	"unsafe"
)

func IntPtr(n int) uintptr {
	return uintptr(n)
}

func StrPtr(s string) uintptr {
	return uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(s)))
}

//func main() {
//	user32:=syscall.NewLazyDLL("user32.dll")
//	MesssageBox:=user32.NewProc("MessageBox")
//	MesssageBox.Call(IntPtr(0), StrPtr("我是消息内容"), StrPtr("提示"), IntPtr(0))
//}

//func main() {
//	user32:=syscall.NewLazyDLL("node.dll")
//	wkeLoadURL:=user32.NewProc("wkeLoadURL")
//	wkeLoadURL.Call(StrPtr("http://www.baidu.com"))
//}

func main() {
	user32,_ := syscall.LoadDLL("node.dll")
	//MessageBoxW,_ := user32.FindProc("wkeLoadURL")
	wkeInitialize,_ := user32.FindProc("wkeInitialize")
	wkeCreateWebWindow,_ := user32.FindProc("wkeCreateWebWindow")
	wkeInitialize.Call()
	wkeCreateWebWindow.Call(StrPtr("http://www.baidu.com"),IntPtr(800),IntPtr(600))
	//MessageBoxW.Call(StrPtr("http://www.baidu.com"))
}
