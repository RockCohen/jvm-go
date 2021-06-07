package stores

import (
	"instructions/base"
	"rtda"
)

// ASTORE Store reference into local variable
/**
和加载指令刚好相反，存储指令把变量从操作数栈顶弹出，然
后存入局部变量表。和加载指令一样，存储指令也可以分为6类。
*/
type ASTORE struct{ base.Index8Instruction }

func (self *ASTORE) Execute(frame *rtda.Frame) {
	_astore(frame, uint(self.Index))
}

type ASTORE_0 struct{ base.NoOperandsInstruction }

func (self *ASTORE_0) Execute(frame *rtda.Frame) {
	_astore(frame, 0)
}

type ASTORE_1 struct{ base.NoOperandsInstruction }

func (self *ASTORE_1) Execute(frame *rtda.Frame) {
	_astore(frame, 1)
}

type ASTORE_2 struct{ base.NoOperandsInstruction }

func (self *ASTORE_2) Execute(frame *rtda.Frame) {
	_astore(frame, 2)
}

type ASTORE_3 struct{ base.NoOperandsInstruction }

func (self *ASTORE_3) Execute(frame *rtda.Frame) {
	_astore(frame, 3)
}

func _astore(frame *rtda.Frame, index uint) {
	ref := frame.OperandStack().PopRef()
	frame.LocalVars().SetRef(index, ref)
}
