package rtda

// Stack
// 对Java虚拟机栈的定义
/**
 Java虚拟机栈应该包括三个元素：
maxSize			栈的最大容量
size			栈当前的容量
_top			栈帧
*/
type Stack struct {
	maxSize uint   //栈的最大容量
	size    uint   //当前栈的大小
	_top    *Frame //帧
}

func newStack(maxSize uint) *Stack {
	return &Stack{
		maxSize: maxSize,
	}
}
func (self *Stack) push(frame *Frame) {
	if self.size >= self.maxSize {
		panic("java.lang.StackOverFlowError")
	}
	if self._top != nil {
		frame.lower = self._top
	}
	self._top = frame
	self.size++
}

func (self *Stack) pop() *Frame {
	if self._top == nil {
		panic("jvm stack is empty!")
	}
	top := self._top
	self._top = top.lower
	top.lower = nil
	self.size--
	return top
}

func (self *Stack) top() *Frame {
	if self._top == nil {
		panic("jvm stack is empty!")
	}
	return self._top
}
