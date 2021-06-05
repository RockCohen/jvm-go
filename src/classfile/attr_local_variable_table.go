package classfile

// LocalVariableTableAttribute
/**
local_variable_table{
	u2 						attribute_name_index;
	u4 						attribute_length;
	u2						local_variable_table_length
    local_variable_info		local_variable_info[local_variable_table_length]
}
LocalVariableTable属性用于描述栈帧中局部变量表的变量与Java源码中定义的变量之间的关系，它
也不是运行时必需的属性，但默认会生成到Class文件之中，可以在Javac中使用-g：none或-g：vars选项
来取消或要求生成这项信息。如果没有生成这项属性，最大的影响就是当其他人引用这个方法时，所
有的参数名称都将会丢失，譬如IDE将会使用诸如arg0、arg1之类的占位符代替原有的参数名，这对程
序运行没有影响，但是会对代码编写带来较大不便，而且在调试期间无法根据参数名称从上下文中获
得参数值。
*/
type LocalVariableTableAttribute struct {
	localVariableTable []*LocalVariableInfo
}

func (self *LocalVariableTableAttribute) readInfo(reader *ClassReader) {
	localVariableTableLength := reader.readUint16()
	self.localVariableTable = make([]*LocalVariableInfo, localVariableTableLength)
	for i := range self.localVariableTable {
		self.localVariableTable[i] = &LocalVariableInfo{
			startPc:    reader.readUint16(),
			length:     reader.readUint16(),
			nameIndex:  reader.readUint16(),
			descriptor: reader.readUint16(),
			index:      reader.readUint16(),
		}
	}
}

// LocalVariableInfo
/**
local_variable_info{
	u2			start_pc
	u2			length
	u2			name_index
	u2			descriptor
	u2			index
}
start_pc和length属性分别代表了这个局部变量的生命周期开始的字节码偏移量及其作用范围覆盖的长度，
两者结合起来就是这个局部变量在字节码之中的作用域范围。
name_index和descriptor_index都是指向常量池中CONSTANT_Utf8_info型常量的索引，
分别代表了局部变量的名称以及这个局部变量的描述符。
index是这个局部变量在栈帧的局部变量表中变量槽的位置。
当这个变量数据类型是64位类型时（double和long），它占用的变量槽为index和index+1两个。

*/
type LocalVariableInfo struct {
	startPc    uint16
	length     uint16
	nameIndex  uint16
	descriptor uint16
	index      uint16
}
