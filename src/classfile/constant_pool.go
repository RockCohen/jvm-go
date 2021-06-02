package classfile

type ConstantInfo interface {
	readInfo(reader *ClassReader)
}

// ConstantPool
/**
常量池占据了class文件很大一部分数据，里面存放着各式各样
的常量信息，包括数字和字符串常量、类和接口名、字段和方法
名等等.
于常量池中常量的数量是不固定的，所以在常量池的入口需要放置一项u2类型的数据，代表常
量池容量计数值（constant_pool_count）。

常量池中每一项常量都是一个表.

常量表类型     						标志值			描述
CONSTANT_Utf8 						1 				UTF-8编码的Unicode字符串
CONSTANT_Integer 					3 				int类型的字面值
CONSTANT_Float 						4 				float类型的字面值
CONSTANT_Long 						5 				long类型的字面值
CONSTANT_Double 					6 				double类型的字面值
CONSTANT_Class 						7 				对一个类或者是接口的符号引用
CONSTANT_String 					8 				String类型的字面值的引用
CONSTANT_Fieldref 					9 				对一个字段的符号
CONSTANT_Methodref 					10 				对一个类中方法的符号应用
CONSTANT_InterfaceMethodref 		11 				对一个接口中方法的符号引用
CONSTANT_NameAndType 				12 				对一个字段或方法的部分符号引用
CONSTANT_MethodHandle               15              对方法句柄的引用
CONSTANT_MethodType                 16              对方法类型的引用
CONSTANT_Dynamic					17              对动态常量的引用
CONSTANT_InvokeDynamic              18              动态调用引用
...

可以把常量池中的常量分为两类：字面量（literal）和符号引用
（symbolic reference）。字面量包括数字常量和字符串常量，符号引
用包括类和接口名、字段和方法信息等。除了字面量，其他常量都
是通过索引直接或间接指向CONSTANT_Utf8_info常量.
*/
type ConstantPool []ConstantInfo

func readConstantPool(reader *ClassReader) ConstantPool {
	// 首先获取常量池数据的数量，数据类型为：u2
	// u2 					constant_pool_count;
	// cp_info 			    constant_pool[constant_pool_count-1];
	cpCount := int(reader.readUint16())
	cp := make([]ConstantInfo, cpCount) // 申请内存
	for i := 1; i < cpCount; i++ {      // 注意索引从 1开始
		cp[i] = readConstantInfo(reader, cp) //第0个表示的常量池的数据的数量
		switch cp[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++ // 占两个位置
		}
	}
	return cp
}

// readConstantInfo
/**
常量池中的表结构起始的第一位是个u1类型的标志位
该标志位表示该表的数据类型
*/
func readConstantInfo(reader *ClassReader, cp ConstantPool) ConstantInfo {
	tag := reader.readUint8()
	c := newConstantInfo(tag, cp)
	c.readInfo(reader)
	return c
}

func newConstantInfo(tag uint8, cp ConstantPool) ConstantInfo {
	switch tag {
	case CONSTANT_Integer:
		return &ConstantIntegerInfo{}
	case CONSTANT_Float:
		return &ConstantFloatInfo{}
	case CONSTANT_Long:
		return &ConstantLongInfo{}
	case CONSTANT_Double:
		return &ConstantDoubleInfo{}
	case CONSTANT_Utf8:
		return &ConstantUtf8Info{}
	case CONSTANT_String:
		return &ConstantStringInfo{cp: cp}
	case CONSTANT_Class:
		return &ConstantClassInfo{cp: cp}
	case CONSTANT_Fieldref:
		return &ConstantFieldrefInfo{ConstantMemberrefInfo{cp: cp}}
	case CONSTANT_Methodref:
		return &ConstantMethodrefInfo{ConstantMemberrefInfo{cp: cp}}
	case CONSTANT_InterfaceMethodref:
		return &ConstantInterfaceMethodrefInfo{ConstantMemberrefInfo{cp: cp}}
	case CONSTANT_NameAndType:
		return &ConstantNameAndTypeInfo{}
	case CONSTANT_MethodType:
		return &ConstantMethodTypeInfo{}
	case CONSTANT_MethodHandle:
		return &ConstantMethodHandleInfo{}
	case CONSTANT_InvokeDynamic:
		return &ConstantInvokeDynamicInfo{}
	default:
		panic("java.lang.ClassFormatError: constant pool tag!")
	}
}

func (self ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if cpInfo := self[index]; cpInfo != nil {
		return cpInfo
	}
	panic("Invalid constant pool index!")
}

func (self ConstantPool) getNameAndType(index uint16) (string, string) {
	//首先获取NameAndType_info字段的索引
	//根据NameAndType_info字段的信息进行检索
	ntInfo := self.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := self.getUtf8(ntInfo.nameIndex)
	_type := self.getUtf8(ntInfo.descriptorIndex)
	return name, _type
}

func (self ConstantPool) getClassName(index uint16) string {
	classInfo := self.getConstantInfo(index).(*ConstantClassInfo)
	return self.getUtf8(classInfo.nameIndex)
}

func (self ConstantPool) getUtf8(index uint16) string {
	utf8Info := self.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}
