package base

import "rtda"

// Branch 指令跳转方法
// 获取当前线程的程序计数器
// 设置下一条指令的地址 基地址+offset
func Branch(frame *rtda.Frame, offset int) {
	pc := frame.Thread().PC()
	nextPC := pc + offset
	frame.SetNextPC(nextPC)
}
