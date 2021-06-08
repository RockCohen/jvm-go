package classfile

import (
	"math"
	"rtda/heap"
)

// 该文件主要定义基本数据类型的引用：
//1. Integer
//2. Float
//3. Long
//4. Double
// 基本结构：
/**
32bit
Integer_info{
	tag			u1
	bytes		u4
}
64bit
Double_info{
	tag			u1
	bytes		u8
}
*/

// ConstantIntegerInfo
// 32位整型数引用
type ConstantIntegerInfo struct {
	val int32
}

func (self *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	self.val = int32(bytes)
}

func (self *ConstantIntegerInfo) Value() heap.Constant {
	return self.val
}

// ConstantFloatInfo
// 32位浮点数引用
type ConstantFloatInfo struct {
	val float32
}

func (self *ConstantFloatInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	self.val = math.Float32frombits(bytes)
}
func (self *ConstantFloatInfo) Value() heap.Constant {
	return self.val
}

// ConstantLongInfo
// 64长整型引用
type ConstantLongInfo struct {
	val int64
}

func (self *ConstantLongInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	self.val = int64(bytes)
}

func (self *ConstantLongInfo) Value() heap.Constant {
	return self.val
}

// ConstantDoubleInfo
// 64双精度浮点数引用
type ConstantDoubleInfo struct {
	val float64
}

func (self *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	self.val = math.Float64frombits(bytes)
}
func (self *ConstantDoubleInfo) Value() heap.Constant {
	return self.val
}
