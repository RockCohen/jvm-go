package classfile

// MemberInfo
/**
实例成员的信息类
class文件中关于实例成员与方法成员的结构如下：

field_info {
	u2 					access_flags;
	u2 					name_index;
	u2 					descriptor_index;
	u2 					attributes_count;
	attribute_info 		attributes[attributes_count];
}
*/
type MemberInfo struct {
	cp              ConstantPool //cp字段保存常量池指针，
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributes      []AttributeInfo
}

// readMembers
// readMembers（）读取字段表或方法表
func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	/**
	u2 					fields_count;
	field_info 			fields[fields_count];
	u2 					methods_count;
	method_info 		methods[methods_count];
	*/
	memberCount := reader.readUint16()
	members := make([]*MemberInfo, memberCount)
	for i := range members {
		members[i] = readMember(reader, cp)
	}
	return members
}

// readMember
// readMember（）函数读取字段或方法数据
func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		cp:              cp,
		accessFlags:     reader.readUint16(),
		nameIndex:       reader.readUint16(),
		descriptorIndex: reader.readUint16(),
		attributes:      readAttributes(reader, cp), // 见 3.4
	}
}

// Name
// Name（）从常量 池查找字段或方法名
func (self *MemberInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}

// Descriptor
// Descriptor（）从常量池查找字段或方法描述符
func (self *MemberInfo) Descriptor() string {
	return self.cp.getUtf8(self.descriptorIndex)
}

func (self *MemberInfo) CodeAttribute() *CodeAttribute {
	for _, attrInfo := range self.attributes {
		switch attrInfo.(type) {
		case *CodeAttribute:
			return attrInfo.(*CodeAttribute)
		}
	}
	return nil
}

func (self *MemberInfo) AccessFlags() uint16 {
	return self.accessFlags

}
func (self *MemberInfo) ConstantValueAttribute() *ConstantValueAttribute {
	for _, attrInfo := range self.attributes {
		switch attrInfo.(type) {
		case *ConstantValueAttribute:
			return attrInfo.(*ConstantValueAttribute)
		}
	}
	return nil
}
