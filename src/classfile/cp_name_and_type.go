package classfile

// ConstantNameAndTypeInfo
// 该字段的结构如下：
/**
ConstantNameAndType_info{
		tag				u1
		index			u2		该字段的方法名称常量索引
		index			u2      该字段的描述符号常量索引
}

1）类型描述符。
	①基本类型byte、short、char、int、long、float和double的描述符
	是单个字母，分别对应B、S、C、I、J、F和D。注意，long的描述符是J
	而不是L。
	②引用类型的描述符是L＋类的完全限定名＋分号。
	③数组类型的描述符是[＋数组元素类型描述符。
2）字段描述符就是字段类型的描述符。
3）方法描述符是（分号分隔的参数类型描述符）+返回值类型描述符，其中void返回值由单个字母V表示。

标识字符			含义
B				基本类型byte
C				基本类型char
D				基本类型double
F				基本类型float
I				基本类型int
J				基本类型long
S				基本类型short
Z				基本类型boolean
V				特殊类型void
L				对象类型，如Ljava/lang/Object;
[				数组类型，多个维度则有多个[

*/
type ConstantNameAndTypeInfo struct {
	nameIndex       uint16
	descriptorIndex uint16
}

func (self *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
	self.descriptorIndex = reader.readUint16()
}
