package control

import (
	"instructions/base"
	"rtda"
)

// TABLE_SWITCH
/*
Java语言中的switch-case语句有两种实现方式：如果case值可以
编码成一个索引表，则实现成tableswitch指令；否则实现成
lookupswitch指令。

defaultOffset对应默认情况下执行跳转所需的字节码偏移量；
low和high记录case的取值范围；jumpOffsets是一个索引表，里面存
放high-low+1个int值，对应各种case情况下，执行跳转所需的字节
码偏移量。
tableswitch
	<0-3 byte pad>

	defaultbyte1
	defaultbyte2
	defaultbyte3
	defaultbyte4

	lowbyte1
	lowbyte2
	lowbyte3
	lowbyte4

	highbyte1
	highbyte2
	highbyte3
	highbyte4

	jump offsets...
*/
// Access jump table by index and jump
type TABLE_SWITCH struct {
	defaultOffset int32   // 默认的跳转地址(匹配源代码中的default关键字)
	low           int32   // 索引表中的下限
	high          int32   // 索引表中的上限
	jumpOffsets   []int32 // 跳转指令索引表
}

func (self *TABLE_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	// tableswitch指令操作码的后面有0~3字节的padding，
	// 以保证 defaultOffset 在字节码中的地址是4的倍数
	reader.SkipPadding()
	self.defaultOffset = reader.ReadInt32()
	self.low = reader.ReadInt32()
	self.high = reader.ReadInt32()
	jumpOffsetsCount := self.high - self.low + 1
	self.jumpOffsets = reader.ReadInt32s(jumpOffsetsCount)
}

func (self *TABLE_SWITCH) Execute(frame *rtda.Frame) {
	index := frame.OperandStack().PopInt()
	var offset int
	if index >= self.low && index <= self.high {
		offset = int(self.jumpOffsets[index-self.low])
	} else {
		offset = int(self.defaultOffset)
	}
	base.Branch(frame, offset)
}
