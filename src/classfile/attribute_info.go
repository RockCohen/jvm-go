package classfile

// 定义Attribute_info结构
/**
attribute_info {
	u2 			attribute_name_index;				指向UTF-8_info的一个索引
	u4 			attribute_length;					属性表的个数
	u1 			info[attribute_length];             属性内容
}
*/

type AttributeInfo interface {
	readInfo(reader *ClassReader)
}

// readAttributeInfo
// 读取属性内容
func readAttributes(reader *ClassReader, cp ConstantPool) []AttributeInfo {
	attributeCount := reader.readUint16()
	attributes := make([]AttributeInfo, attributeCount)
	for i := range attributes {
		attributes[i] = readAttribute(reader, cp)
	}
	return attributes
}

// 每一个属性表都包括如下的几个部分
// 属性表特征包括：
// 1. 属性名索引
// 2. 属性表长度
// 3. 属性信息数组
func readAttribute(reader *ClassReader, cp ConstantPool) AttributeInfo {
	attributeNameIndex := reader.readUint16()
	attributeName := cp.getUtf8(attributeNameIndex)
	attributeLen := reader.readUint32()
	attributeInfo := newAttributeInfo(attributeName, attributeLen, cp)

	attributeInfo.readInfo(reader)
	return attributeInfo
}

// newAttributeInfo
/**
按照用途，23种预定义属性可以分为三组。
第一组属性是实现Java虚拟机所必需的，共有5种；
第二组属性是Java类库所必需的，共有12种；
第三组属性主要提供给工具使用，共有6种。
第三组属性是可选的，也就是说可以不出现在class文件中。
如果class文件中存在第三组属性，Java虚拟机实现或者Java类库也是可以利用它们
的，比如使用LineNumberTable属性在异常堆栈中显示行号。
*/
func newAttributeInfo(attributeName string, attributeLen uint32, cp ConstantPool) AttributeInfo {
	switch attributeName {
	case "Code":
		return &CodeAttribute{cp: cp}
	case "ConstantValue":
		return &ConstantValueAttribute{}
	case "Exceptions":
		return &ExceptionsAttribute{}
	case "Deprecated":
		return &DeprecatedAttribute{}
	case "LineNumberTable":
		return &LineNumberTableAttribute{}
	case "LocalVariableTable":
		return &LocalVariableTableAttribute{}
	case "SourceFile":
		return &SourceFileAttribute{cp: cp}
	case "Synthetic":
		return &SyntheticAttribute{}
	default:
		return &UnparsedAttribute{attributeName, attributeLen, nil}
	}
}
