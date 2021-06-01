package classfile

import "encoding/binary"

// ClassReader
/**
ClassReader 用于读取class文件作格式化处理
*/
type ClassReader struct {
	data []byte // 存放class文件的字节数组
}

// 用于读取u1类型的数据
// 获取当前data数组中的第一个字节，并且将该字节从数组中移除
func (self *ClassReader) readUint8() uint8 {
	val := self.data[0]
	self.data = self.data[1:]
	return val
}

// 用于读取u2类型的数据
// 获取当前data数组的前两个字节，并将其从中移除。
// Go标准库encoding/binary包中定义了一个变量BigEndian，正好
// 可以从[]byte中解码多字节数据。
func (self *ClassReader) readUint16() uint16 {
	val := binary.BigEndian.Uint16(self.data)
	self.data = self.data[2:]
	return val
}

// 用于读取u4类型的数据
func (self *ClassReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(self.data)
	self.data = self.data[4:]
	return val
}

// 用于读取u8类型的数据
func (self *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(self.data)
	self.data = self.data[8:]
	return val
}

// 用于获取u2类型的表
func (self *ClassReader) readUint16s() []uint16 {
	length := self.readUint16()
	s := make([]uint16, length) //内存分配函数
	for i := range s {
		s[i] = self.readUint16()
	}
	return s
}

// 用于读取指定数量的字节
func (self *ClassReader) readBytes(n uint32) []byte {
	bytes := self.data[:n]
	self.data = self.data[n:]
	return bytes
}
