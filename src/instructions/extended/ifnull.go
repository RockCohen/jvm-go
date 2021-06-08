package extended

import (
	"instructions/base"
	"rtda"
)

// IFNULL Branch if reference is null
type IFNULL struct{ base.BranchInstruction }

func (self *IFNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		base.Branch(frame, self.Offset)
	}
}

// IFNONNULL Branch if reference not null
type IFNONNULL struct{ base.BranchInstruction }

func (self *IFNONNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref != nil {
		base.Branch(frame, self.Offset)
	}
}
