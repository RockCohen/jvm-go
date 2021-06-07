package classfile

// LineNumberTableAttribute
/**
LineNumberTable_attribute {
	u2 						attribute_name_index;
	u4 						attribute_length;
	u2 						line_number_table_length;
	line_number_info		line_number_table[line_number_table_length];
}

LineNumberTable属性用于描述Java源码行号与字节码行号（字节码的偏移量）之间的对应关系。
它并不是运行时必需的属性，但默认会生成到Class文件之中，可以在Javac中使用-g：none或-g：lines
选项来取消或要求生成这项信息。如果选择不生成LineNumberTable属性，对程序运行产生的最主要影
响就是当抛出异常时，堆栈中将不会显示出错的行号，并且在调试程序的时候，也无法按照源码行来
设置断点。
*/
type LineNumberTableAttribute struct {
	lineNumberTable []*LineNumberInfo
}

func (self *LineNumberTableAttribute) readInfo(reader *ClassReader) {
	lineNumberTableLength := reader.readUint16()
	self.lineNumberTable = make([]*LineNumberInfo, lineNumberTableLength)
	for i := range self.lineNumberTable {
		self.lineNumberTable[i] = &LineNumberInfo{
			startPc:    reader.readUint16(),
			lineNumber: reader.readUint16(),
		}
	}
}

// LineNumberInfo
/**
字节码行号与Java源码行号对应关系：
line_number_info{
	u2 				start_pc;			字节码行号
	u2 				line_number;		Java源码行号
}
*/
type LineNumberInfo struct {
	startPc    uint16
	lineNumber uint16
}
