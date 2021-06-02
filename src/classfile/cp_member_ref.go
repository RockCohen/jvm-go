package classfile

// ConstantMemberrefInfo
// 包括方法成员和实例成员，接口
// 结构如下：
/**
_CONSTANT_Fieldref/Methodref_info{
	tag				u1
	index			u2			指向成员所在类或者接口的Class_info索引
	index			u2			指向成员的NameAndType索引
}
*/
type ConstantMemberrefInfo struct {
	cp               ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (self *ConstantMemberrefInfo) readInfo(reader *ClassReader) {
	self.classIndex = reader.readUint16()
	self.nameAndTypeIndex = reader.readUint16()
}
func (self *ConstantMemberrefInfo) ClassName() string {
	// getClassName 方法调用get_utf8获取UTF-8字符表中的类型信息
	return self.cp.getClassName(self.classIndex)
}
func (self *ConstantMemberrefInfo) NameAndDescriptor() (string, string) {
	return self.cp.getNameAndType(self.nameAndTypeIndex)
}

// 继承MemberInfo来分别定义：
// 1. ConstantFieldrefInfo
// 2. ConstantMethodrefInfo
// 3. ConstantInterfaceMethodrefInfo

type ConstantFieldrefInfo struct{ ConstantMemberrefInfo }

type ConstantMethodrefInfo struct{ ConstantMemberrefInfo }

type ConstantInterfaceMethodrefInfo struct{ ConstantMemberrefInfo }
