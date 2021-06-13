package base

import (
	"rtda"
	"rtda/heap"
)

func InvokeMethod(invokerFrame *rtda.Frame, method *heap.Method) {
	thread := invokerFrame.Thread()     //获取当前线程
	newFrame := thread.NewFrame(method) //新的方法栈帧
	argSlotSlot := int(method.ArgSlotCount())
	if argSlotSlot > 0 {
		for i := argSlotSlot - 1; i >= 0; i-- {
			slot := invokerFrame.OperandStack().PopSlot()
			newFrame.LocalVars().SetSlot(uint(i), slot)
		}
	}

}
