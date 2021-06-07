package loads

import (
	"instructions/base"
	"rtda"
)

// ALOAD Load reference from local variable
/**
加载指令从局部变量表获取变量，然后推入操作数栈顶。加载
指令共33条，按照所操作变量的类型可以分为6类：aload系列指令
操作引用类型变量、dload系列操作double类型变量、fload系列操作
float变量、iload系列操作int变量、lload系列操作long变量、xaload操
作数组。
*/
type ALOAD struct{ base.Index8Instruction }

func (self *ALOAD) Execute(frame *rtda.Frame) {
	_aload(frame, self.Index)
}

type ALOAD_0 struct{ base.NoOperandsInstruction }

func (self *ALOAD_0) Execute(frame *rtda.Frame) {
	_aload(frame, 0)
}

type ALOAD_1 struct{ base.NoOperandsInstruction }

func (self *ALOAD_1) Execute(frame *rtda.Frame) {
	_aload(frame, 1)
}

type ALOAD_2 struct{ base.NoOperandsInstruction }

func (self *ALOAD_2) Execute(frame *rtda.Frame) {
	_aload(frame, 2)
}

type ALOAD_3 struct{ base.NoOperandsInstruction }

func (self *ALOAD_3) Execute(frame *rtda.Frame) {
	_aload(frame, 3)
}

func _aload(frame *rtda.Frame, index uint) {
	ref := frame.LocalVars().GetRef(index)
	frame.OperandStack().PushRef(ref)
}
