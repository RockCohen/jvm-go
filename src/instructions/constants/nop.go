package constants

import (
	"instructions/base"
	"rtda"
)

// NOP
// 该指令啥也不做
// 隐含操作数的常量值零可以直接继承 base.NoOperandsInstruction

type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Execute(frame *rtda.Frame) {
	// nothing to do
}
