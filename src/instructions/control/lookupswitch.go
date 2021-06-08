package control

import (
	"instructions/base"
	"rtda"
)

//LOOKUP_SWITCH
/*
matchOffsets有点像Map，它的key是case值，value是跳转偏移
量。Execute（）方法先从操作数栈中弹出一个int变量，然后用它查找
matchOffsets，看是否能找到匹配的key。如果能，则按照value给出的
偏移量跳转，否则按照defaultOffset跳转。
lookupswitch
<0-3 byte pad>
defaultbyte1
defaultbyte2
defaultbyte3
defaultbyte4
npairs1
npairs2
npairs3
npairs4
match-offset pairs...
*/
// Access jump table by key match and jump
type LOOKUP_SWITCH struct {
	defaultOffset int32
	npairs        int32
	matchOffsets  []int32 // 用数组存放key-value对，其中k(k%2==0)存放key值，k+1存放value
}

func (self *LOOKUP_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding()
	self.defaultOffset = reader.ReadInt32()
	self.npairs = reader.ReadInt32()
	self.matchOffsets = reader.ReadInt32s(self.npairs * 2)
}

func (self *LOOKUP_SWITCH) Execute(frame *rtda.Frame) {
	key := frame.OperandStack().PopInt()
	for i := int32(0); i < self.npairs*2; i += 2 {
		if self.matchOffsets[i] == key {
			offset := self.matchOffsets[i+1]
			base.Branch(frame, int(offset))
			return
		}
	}
	base.Branch(frame, int(self.defaultOffset))
}
