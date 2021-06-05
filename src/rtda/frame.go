package rtda

type Frame struct {
	lower        *Frame
	localVars    LocalVars
	operandStack *OperandStack
}

func NewFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}
func (self *Frame) GetLocalVars() LocalVars {
	return self.localVars
}
func (self *Frame) GetOperandStack() *OperandStack {
	return self.operandStack
}
