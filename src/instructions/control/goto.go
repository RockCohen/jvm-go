package control

import (
	"instructions/base"
	"rtda"
)

// GOTO Branch always
// goto指令进行无条件跳转
type GOTO struct{ base.BranchInstruction }

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}
