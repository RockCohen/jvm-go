package rtda

import "rtda/heap"

// Frame stack frame
// 对栈帧的定义
/**
通过链表（数据结构）实现栈帧的连接
变量解释：
	lower 				指向下一帧的指针
	localVars   		局部变量表指针
	operandStack		操作栈指针
	thread				栈帧所属线程指针
	nextPC				下一条指令

至此为止：
Thread Stack Frame 的关系如下：

	Thread					Stack                 Frame
-------------          ---------------        ---------------
|   *stack  |--------> |   *_top     |------> |    *lower   |--------> ......
|	pc	    |		   | size/maxSize|        |     .....   |
-------------          ---------------        ---------------
*/

// Frame stack frame
type Frame struct {
	lower        *Frame // stack is implemented as linked list
	localVars    LocalVars
	operandStack *OperandStack
	thread       *Thread
	method       *heap.Method
	nextPC       int // the next instruction after the call
}

/**
执行方法所需的局部变量表大小和操作数栈深度是由编译器
预先计算好的，存储在class文件method_info结构的Code属性中
*/
func newFrame(thread *Thread, method *heap.Method) *Frame {
	return &Frame{
		thread:       thread,
		method:       method,
		localVars:    newLocalVars(method.MaxLocals()),
		operandStack: newOperandStack(method.MaxStack()),
	}
}

// getters & setters
func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}
func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}
func (self *Frame) Thread() *Thread {
	return self.thread
}
func (self *Frame) Method() *heap.Method {
	return self.method
}
func (self *Frame) NextPC() int {
	return self.nextPC
}
func (self *Frame) SetNextPC(nextPC int) {
	self.nextPC = nextPC
}
