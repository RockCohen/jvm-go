package base

import "rtda"

// Instruction
/**
字节码指令的通用接口
包括两个方法：
1. FetchOperands : 获取字节码指令中的操作数
2. Execute : 执行字节码指令
*/
type Instruction interface {
	FetchOperands(reader *BytecodeReader)
	Execute(frame *rtda.Frame)
}

// NoOperandsInstruction
// 表示什么都不用作的字节码指令，其实现的接口方法体为空
type NoOperandsInstruction struct{}

func (self *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {
	// nothing to do
}

// BranchInstruction
// 跳转指令
type BranchInstruction struct {
	Offset int
}

func (self *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	self.Offset = int(reader.ReadInt16())
}

// Index8Instruction
/**
存储和加载类指令需要根据索引存取局部变量表，索引由单字节操作数给出。
把这类指令抽象成Index8Instruction结构体，
用Index字段表示局部变量表索引。
*/
type Index8Instruction struct {
	Index uint
}

func (self *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint8())
}

// Index16Instruction
/**
有一些指令需要访问运行时常量池，常量池索引由两字节操
作数给出。把这类指令抽象成Index16Instruction结构体，用Index字
段表示常量池索引。
*/
type Index16Instruction struct {
	Index uint
}

func (self *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint16())
}
