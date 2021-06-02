package classfile

// ConstantClassInfo
// class_info
/**
Class_info{
	tag          u1
    index        u2			指向全限定名常量的索引
}
*/
type ConstantClassInfo struct {
	cp        ConstantPool // 常量池句柄
	nameIndex uint16       //指向全限定名常量的索引
}

func (self *ConstantClassInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
}

// Name
// 获取类名字符串
func (self *ConstantClassInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}
