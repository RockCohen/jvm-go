package comparisons

import (
	"instructions/base"
	"rtda"
)

// FCMPG Compare float
/**
。由于浮点数计算有可能产生NaN（Not a
Number）值，所以比较两个浮点数时，除了大于、等于、小于之外，
还有第4种结果：无法比较。

当两个float变量中至少有一个是NaN时，用fcmpg指
令比较的结果是1，而用fcmpl指令比较的结果是-1。
*/
type FCMPG struct{ base.NoOperandsInstruction }

func (self *FCMPG) Execute(frame *rtda.Frame) {
	_fcmp(frame, true)
}

type FCMPL struct{ base.NoOperandsInstruction }

func (self *FCMPL) Execute(frame *rtda.Frame) {
	_fcmp(frame, false)
}

func _fcmp(frame *rtda.Frame, gFlag bool) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else if v1 < v2 {
		stack.PushInt(-1)
	} else if gFlag {
		stack.PushInt(1)
	} else {
		stack.PushInt(-1)
	}
}
