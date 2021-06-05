package constants

import (
	"instructions/base"
	"rtda"
)

// NOP
// 该指令啥也不做
type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Execute(frame *rtda.Frame) {
	// nothing to do
}
