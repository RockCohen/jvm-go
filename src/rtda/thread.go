package rtda

// Thread
// 对线程级的封装
// Java虚拟机运行时数据区域可以分为线程私有与多线程共享
// 线程私有包括：程序计数器，Java虚拟机栈
/**
该结构体包括两个变量：
pc			: 程序计数器
stack		: 虚拟机栈
*/
type Thread struct {
	pc    int
	stack *Stack
}

// NewThread
// 创建线程中栈帧 的个数为1024
// 可以通过设置命令行参数来设置虚拟机栈的数量
func NewThread() *Thread {
	return &Thread{
		stack: newStack(1024),
	}
}

func (self *Thread) PC() int {
	return self.pc
}

func (self *Thread) SetPC(pc int) {
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

func (self *Thread) NewFrame(maxLocals, maxStack uint) *Frame {
	return newFrame(self, maxLocals, maxStack)
}
