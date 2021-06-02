package classfile

// ConstantUtf8Info
/**
UTF-8编码字符串
UTF-8_info结构：
UTF-8_info{
      tag       u1
      length    u2
      bytes    u1
}
用来存放常量池中用utf-8进行编码的字符，比如class类名，通过index索引指向该字符数组中的某个值
*/
type ConstantUtf8Info struct {
	str string
}

// readInfo
// readConstantInfo 方法中已经对tag标志位进行读取了，这里无需再进行读取。
// 相当于把所有的常量池中的表的tag标志位抽取出来在 readConstantInfo 中实现读取。
func (self *ConstantUtf8Info) readInfo(reader *ClassReader) {
	length := uint32(reader.readUint16())
	bytes := reader.readBytes(length)
	self.str = decodeMUTF8(bytes)
}

// 简化版，完整版参考
func decodeMUTF8(bytes []byte) string {
	return string(bytes)
}
