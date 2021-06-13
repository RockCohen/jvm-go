package rtda

import (
	"math"
	"rtda/heap"
)

// OperandStack
// 对操作数栈的定义
/**
由于编译器在编译Java程序时已经将操作数栈的最大容量记录在class文件的Code属性中。
所以直接利用Slot数组进行操作数栈的模拟。

	size	栈顶指针（指向栈顶元素的下一个位置）
	slots	Slot数组，存放具体的操作数
*/
type OperandStack struct {
	size  uint
	slots []Slot
}

func newOperandStack(maxStack uint) *OperandStack {
	if maxStack > 0 {
		return &OperandStack{
			slots: make([]Slot, maxStack),
		}
	}
	return nil
}
func (self *OperandStack) PushInt(val int32) {
	self.slots[self.size].num = val
	self.size++
}
func (self *OperandStack) PopInt() int32 {
	self.size--
	return self.slots[self.size].num
}

func (self *OperandStack) PushFloat(val float32) {
	bits := math.Float32bits(val)
	self.slots[self.size].num = int32(bits)
	self.size++
}
func (self *OperandStack) PopFloat() float32 {
	self.size--
	bits := uint32(self.slots[self.size].num)
	return math.Float32frombits(bits)
}

func (self *OperandStack) PushLong(val int64) {
	self.slots[self.size].num = int32(val)
	self.slots[self.size+1].num = int32(val >> 32)
	self.size += 2
}
func (self *OperandStack) PopLong() int64 {
	self.size -= 2
	low := uint32(self.slots[self.size].num)
	high := uint32(self.slots[self.size+1].num)
	return int64(high)<<32 | int64(low)
}
func (self *OperandStack) PushDouble(val float64) {
	bits := math.Float64bits(val)
	self.PushLong(int64(bits))
}
func (self *OperandStack) PopDouble() float64 {
	bits := uint64(self.PopLong())
	return math.Float64frombits(bits)
}
func (self *OperandStack) PushRef(ref *heap.Object) {
	self.slots[self.size].ref = ref
	self.size++
}
func (self *OperandStack) PopRef() *heap.Object {
	self.size--
	ref := self.slots[self.size].ref
	self.slots[self.size].ref = nil // 使用go的垃圾回收技术进行垃圾回收
	return ref
}

/**
下面两个方法用于栈指令的操作，由于栈相关的指令不关心数据的类型，所以直接操作即可。
*/

func (self *OperandStack) PushSlot(slot Slot) {
	self.slots[self.size] = slot
	self.size++
}
func (self *OperandStack) PopSlot() Slot {
	self.size--
	return self.slots[self.size]
}

func (self *OperandStack) GetRefFromTop(n uint) *heap.Object {
	return self.slots[self.size-1-n].ref
}
