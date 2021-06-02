package classfile

// ConstantStringInfo
// 字符串字面量常量
// 结构与Class_info一致
// 因此，方法如出一辙
type ConstantStringInfo struct {
	cp          ConstantPool
	stringIndex uint16
}

func (self *ConstantStringInfo) readInfo(reader *ClassReader) {
	self.stringIndex = reader.readUint16()
}

func (self *ConstantStringInfo) String() string {
	return self.cp.getUtf8(self.stringIndex)
}
