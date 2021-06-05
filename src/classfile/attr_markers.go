package classfile

// DeprecatedAttribute
// Deprecated关键字用于标注不建议使用的类、方法等
type DeprecatedAttribute struct{ MarkerAttribute }

// SyntheticAttribute
// Synthetic属性用来标记源文件中不存在、由编译器生成的类成
//员，引入Synthetic属性主要是为了支持嵌套类和嵌套接口。
type SyntheticAttribute struct{ MarkerAttribute }

// MarkerAttribute
// DeprecatedAttribute 与 SyntheticAttribute 通过继承该类实现
type MarkerAttribute struct{}

func (self *MarkerAttribute) readInfo(reader *ClassReader) {
	// read nothing
	// 由于这两个属性没有具体的内容，所以该方法为空。
}
