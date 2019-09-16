package utils

import (
	"os"
	"os/signal"
	"syscall"
)
//ShutDownHook 服务端关闭勾子事件
func ShutDownHook(f func())  {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)
	<-quit
	//log.Println("ShutDownHook ....")
	f()
}
