package rtda

type Thread struct {
	pc    int
	stack *Stack
}

// NewThread
// 创建线程中栈帧 的个数为1024
func NewThread() *Thread {
	return &Thread{
		stack: newStack(1024),
	}
}

func (self *Thread) PC() int {
	return self.pc
}

func (self *Thread) setPC(pc int) {
	self.pc = pc
}

func (self *Thread) PushFrame(frame *Frame) {
	self.stack.push(frame)
}

func (self *Thread) PopFrame() *Frame {
	return self.stack.pop()
}
func (self *Thread) CurrentFrame() *Frame {
	return self.stack.top()
}
