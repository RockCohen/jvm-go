package stack

import (
	"instructions/base"
	"rtda"
)

// POP Pop the top operand stack value
type POP struct{ base.NoOperandsInstruction }

// Execute
/*
bottom -> top
[...][c][b][a]
            |
            V
[...][c][b]
*/
func (self *POP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
}

// POP2 Pop the top one or two operand stack values
type POP2 struct{ base.NoOperandsInstruction }

// Execute
/*
bottom -> top
[...][c][b][a]
         |  |
         V  V
[...][c]
*/
func (self *POP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}
