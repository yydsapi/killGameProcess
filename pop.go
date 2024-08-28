package main

import (
	"github.com/go-toast/toast"
	"log"
	"syscall"
	"unsafe"
)

func IntPtr(n int) uintptr {
	return uintptr(n)
}
func StrPtr(s string) uintptr {
	return uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(s)))
}

// ShowMessage windows下的另一种DLL方法调用
func ShowMessage(title, text string) {
	user32dll, _ := syscall.LoadLibrary("user32.dll")
	user32 := syscall.NewLazyDLL("user32.dll")
	MessageBoxW := user32.NewProc("MessageBoxW")
	MessageBoxW.Call(IntPtr(0), StrPtr(text), StrPtr(title), IntPtr(0))
	defer syscall.FreeLibrary(user32dll)
}

func PushNotification(x string, y string) {
	notification := toast.Notification{
		AppID:   "Microsoft.Windows.Shell.RunDialog",
		Title:   x,
		Message: y,
		//Icon:    "C:\\path\\to\\your\\logo.png", // 文件必须存在
		Actions: []toast.Action{
			{"protocol", "确定", "https://www.google.com/"},
		},
	}
	err := notification.Push()
	if err != nil {
		log.Fatalln(err)
	}
}
