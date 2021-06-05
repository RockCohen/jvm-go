package classfile

// CodeAttribute
/**
Code_attribute {
	u2 					attribute_name_index; 			CONSTANT_Utf8_info型常量索引，定值为"Code"
	u4 					attribute_length;				length
	u2 					max_stack;						操作数栈的最大深度
	u2 					max_locals;						局部变量的存储空间（单位:Slot,槽，uint32）
	u4 					code_length;					字节码指令长度
	u1 					code[code_length];				java源程序编译后生成的字节码指令
	u2 					exception_table_length;
	u1					exception_table[exception_table_length];
	u2 					attributes_count;
	attribute_info 		attributes[attributes_count];
}

对上述的一些变量进行详细的说明：
max_stack	:		在方法执行的任意时刻，操作数栈都不会超过这个深度。虚拟机运行的时候需要根据这个值来分配栈帧（Stack Frame）中的操作栈深度。
max_locals	:		方法参数（包括实例方法中的隐藏参数“this”）、显式异常处理程序的参数（Exception Handler Parameter，就是try-catch语句中catch块中所定义的异常）、方法体中
定义的局部变量都需要依赖局部变量表来存放,该空间的使用方法是：重用。javac会根据变量的作用域来重用该控件，详细一点就是：如果槽中的变量已经不在当前的作用域中，那么便可以进行覆盖。
*/
type CodeAttribute struct {
	cp             ConstantPool
	maxStack       uint16
	maxLocals      uint16
	code           []byte
	exceptionTable []*ExceptionTableEntry
	attributes     []AttributeInfo
}

// ExceptionTableEntry
/**
exception_table{
	u2 				start_pc;
	u2 				end_pc;
	u2 				handler_pc;
	u2 				catch_type;
}
异常表：
如果当字节码从第start_pc行[1]到第end_pc行之间（不含第end_pc行）出现了类型为catch_type或者其子类的异常
（catch_type为指向一个CONSTANT_Class_info型常量的索引），则转到第handler_pc行继续处理。
当catch_type的值为0时，代表任意异常情况都需要转到handler_pc处进行处理。

异常表实际上是Java代码的一部分，尽管字节码中有最初为处理异常而设计的跳转指令，但《Java
虚拟机规范》中明确要求Java语言的编译器应当选择使用异常表而不是通过跳转指令来实现Java异常及
finally处理机制
*/
type ExceptionTableEntry struct {
	startPc   uint16
	endPc     uint16
	handlerPc uint16
	catchType uint16
}

func (self *CodeAttribute) readInfo(reader *ClassReader) {
	self.maxStack = reader.readUint16()
	self.maxLocals = reader.readUint16()
	codeLength := reader.readUint32()
	self.code = reader.readBytes(codeLength)
	self.exceptionTable = readExceptionTable(reader)
	self.attributes = readAttributes(reader, self.cp)
}

func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
	exceptionLength := reader.readUint16()
	exceptions := make([]*ExceptionTableEntry, exceptionLength)
	for i := range exceptions {
		exceptions[i] = &ExceptionTableEntry{
			startPc:   reader.readUint16(),
			endPc:     reader.readUint16(),
			handlerPc: reader.readUint16(),
			catchType: reader.readUint16(),
		}
	}
	return exceptions
}
func (self *CodeAttribute) MaxStack() uint {
	return uint(self.maxStack)
}
func (self *CodeAttribute) MaxLocals() uint {
	return uint(self.maxLocals)
}
func (self *CodeAttribute) Code() []byte {
	return self.code
}
func (self *CodeAttribute) ExceptionTable() []*ExceptionTableEntry {
	return self.exceptionTable
}

func (self *ExceptionTableEntry) StartPc() uint16 {
	return self.startPc
}
func (self *ExceptionTableEntry) EndPc() uint16 {
	return self.endPc
}
func (self *ExceptionTableEntry) HandlerPc() uint16 {
	return self.handlerPc
}
func (self *ExceptionTableEntry) CatchType() uint16 {
	return self.catchType
}
