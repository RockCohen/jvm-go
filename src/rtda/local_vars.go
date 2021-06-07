package rtda

import (
	"math"
	"rtda/heap"
)

// LocalVars
/**
局部变量表（Local Variables Table）是一组变量值的存储空间，用于存放方法参数和方法内部定义的局部变量。
在Java程序被编译为Class文件时，就在方法的Code属性的max_locals数据项中确定了该方法所需分配的局部变量表的最大容量。
由于局部变量表是建立在线程堆栈中的，属于线程私有的数据，无论读写两个连续的变量槽是否为原子操作，都不会引起数据竞争和线程安全问题。
*/
type LocalVars []Slot

func newLocalVars(maxLocals uint) LocalVars {
	if maxLocals > 0 {
		return make([]Slot, maxLocals)
	}
	return nil
}

// getter and setter
// 对于64bit的数据，需要对Slot数组执行两次取数据的操作

func (self LocalVars) SetInt(index uint, val int32) {
	self[index].num = val
}
func (self LocalVars) GetInt(index uint) int32 {
	return self[index].num
}
func (self LocalVars) SetFloat(index uint, val float32) {
	bits := math.Float32bits(val)
	self[index].num = int32(bits)
}
func (self LocalVars) GetFloat(index uint) float32 {
	bits := uint32(self[index].num)
	return math.Float32frombits(bits)
}
func (self LocalVars) SetLong(index uint, val int64) {
	self[index].num = int32(val)
	self[index+1].num = int32(val >> 32)
}
func (self LocalVars) GetLong(index uint) int64 {
	low := uint32(self[index].num)
	high := uint32(self[index+1].num)
	return int64(high)<<32 | int64(low)
}
func (self LocalVars) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	self.SetLong(index, int64(bits))
}
func (self LocalVars) GetDouble(index uint) float64 {
	bits := uint64(self.GetLong(index))
	return math.Float64frombits(bits)
}
func (self LocalVars) SetRef(index uint, ref *heap.Object) {
	self[index].ref = ref
}
func (self LocalVars) GetRef(index uint) *heap.Object {
	return self[index].ref
}
