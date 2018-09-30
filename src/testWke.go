package main

/*
#cgo CFLAGS: -g -std=c99
#include "wkedefine.h"
*/
import "C"
import "unsafe"

type WebView struct {
	v *C.wkeWebView
}

func (w *WebView) LoadURL(url string) {
	s := C.CString(url)
	C.wkeLoadURL(w.v, (*C.utf8)(s))
	C.free(unsafe.Pointer(s))
}

func Initialize() {
	C.wkeInitialize()
}

func main() {
	Initialize();

	//v := C.wkeCreateWebWindow(C.WKE_WINDOW_TYPE_POPUP,C.NULL,)
	//w := &WebView{v: v}
	//w.LoadURL("http://www.baidu.com")
}
