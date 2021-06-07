package classfile

// ExceptionsAttribute
/**
Exceptions_attribute {
	u2 					attribute_name_index;
	u4 					attribute_length;
	u2 					number_of_exceptions;
	u2 					exception_index_table[number_of_exceptions];
}


Exceptions属性的作用是列举出方法中可能抛出的受查异常（Checked Exceptions），也
就是方法描述时在throws关键字后面列举的异常。
此属性中的number_of_exceptions项表示方法可能抛出number_of_exceptions种受查异常，
每一种受查异常使用一个exception_index_table项表示；
exception_index_table是一个指向常量池中CONSTANT_Class_info型常量的索引，代表了该受查异常的类型。
*/
type ExceptionsAttribute struct {
	exceptionIndexTable []uint16
}

func (self ExceptionsAttribute) readInfo(reader *ClassReader) {
	self.exceptionIndexTable = reader.readUint16s()
}
func (self *ExceptionsAttribute) ExceptionIndexTable() []uint16 {
	return self.exceptionIndexTable
}
