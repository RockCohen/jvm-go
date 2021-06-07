package classfile

// ConstantValueAttribute
/**
ConstantValue_attribute {
	u2 					attribute_name_index;    数量1
	u4 					attribute_length;		 数量1
	u2 					constantvalue_index;	 数量1
}

ConstantValue是一个定长属性。
attribute_name_index：
attribute_length：固定为2
constantvalue_index：数据项代表了常量池中一个字面量常量的引用
	long								CONSTANT_Long_info,
	float								CONSTANT_Float_info,
	double								CONSTANT_Double_info,
	boolean,byte,char,short,int			CONSTANT_Integer_info,
	java.lang.String					CONSTANT_String_info.


ConstantValue属性的作用是通知虚拟机自动为静态变量赋值。
只有被static关键字修饰的变量（类变量）才可以使用这项属性。

而对于类变量，则有两种方式可以选择：
1. 在类构造器<clinit>()方法中
2. 使用ConstantValue属性。

目前的甲骨文的做法是：
对于被final与static同时修饰的基本数据类型或者java.lang.String，用ConstantValue属性进行赋值；
而对于其他的静态变量使用类构造器<clinit>()进行赋值。

由于ConstantValue属性的属性值只是常量池的一个索引，由于Class文件格式的常量类型中只有与基本属性和字符串相对应的字面量，所以ConstantValue只支持基本数据类型的常量。
*/
type ConstantValueAttribute struct {
	constantValueIndex uint16
}

func (self *ConstantValueAttribute) readInfo(reader *ClassReader) {
	self.constantValueIndex = reader.readUint16()
}
func (self *ConstantValueAttribute) ConstantValueIndex() uint16 {
	return self.constantValueIndex
}
