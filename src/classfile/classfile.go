package classfile

import "fmt"

// ClassFile
/**
class文件结构体
构成class文件的基本数据单位是字节，可以把整个class文件当成一个字节流来处理。

为了描述class文件格式，Java虚拟机规范定义了u1、u2和u4三种数据类型来表示1、2和4字节无符号整数，
分别对应Go语言的uint8、uint16和uint32类型。

相同类型的多条数据一般按表（table）的形式存储在class文件中。
表由表头和表项（item）构成，表头是u2或u4整数。假设表头是n，后面就紧跟着n个表项数据。

java的class文件的结构如下：

_ClassFile{
	u4                	magic;
	u2                	minor_version;
	u2 					major_version;
	u2 					constant_pool_count;
	cp_info 			constant_pool[constant_pool_count-1];
	u2 					access_flags;
	u2 					this_class;
	u2 					super_class;
	u2 					interfaces_count;
	u2 					interfaces[interfaces_count];
	u2 					fields_count;
	field_info 			fields[fields_count];
	u2 					methods_count;
	method_info 		methods[methods_count];
	u2 					attributes_count;
	attribute_info 		attributes[attributes_count];
}


关于golang中的访问权限：
Go的访问控制非常简单：只有公开和私有两种。
所有首字母大写的类型、结构体、字段、变量、函数、方法等都是公开的，可供其他包使用。
首字母小写则是私有的，只能在包内部使用。
*/
type ClassFile struct {
	magic        uint32          //魔数
	minorVersion uint16          //次版本号
	majorVersion uint16          //主版本号
	constantPool ConstantPool    //常量池
	accessFlags  uint16          //访问权限
	thisClass    uint16          //自类
	superClass   uint16          //父类
	interfaces   []uint16        //接口
	fields       []*MemberInfo   //变量成员
	methods      []*MemberInfo   //方法成员
	attributes   []AttributeInfo //属性表
}

// Parse
// Parse（）函数把[]byte解析成ClassFile结构体
func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()
	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return
}

// read
// read 执行具体的解析过程
func (self *ClassFile) read(reader *ClassReader) {
	self.readAndCheckMagic(reader)               // 见 3.2.3
	self.readAndCheckVersion(reader)             // 见 3.2.4
	self.constantPool = readConstantPool(reader) // 见 3.3
	self.accessFlags = reader.readUint16()
	self.thisClass = reader.readUint16()
	self.superClass = reader.readUint16()
	self.interfaces = reader.readUint16s()
	self.fields = readMembers(reader, self.constantPool) // 见 3.2.8
	self.methods = readMembers(reader, self.constantPool)
	self.attributes = readAttributes(reader, self.constantPool) //见 3.4
}

// MajorVersion
// MajorVersion getter
func (self *ClassFile) MajorVersion() uint16 {
	return self.majorVersion
}

// ClassName
// ClassName 从常量池中获取类名
func (self *ClassFile) ClassName() string {
	return self.constantPool.getClassName(self.thisClass)
}

// SuperClassName
// SuperClassName从常量池中获取超类名
func (self *ClassFile) SuperClassName() string {
	if self.superClass > 0 {
		return self.constantPool.getClassName(self.superClass)
	}
	return "" // 只有java.lang.Object 没有超类
}

// InterfaceNames
// InterfaceNames 从常量池中获取接口名
func (self *ClassFile) InterfaceNames() []string {
	interfaceNames := make([]string, len(self.interfaces))
	for i, cpIndex := range self.interfaces {
		interfaceNames[i] = self.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}

// readAndCheckMagic
// readAndCheckMagic 负责读取魔数并检查魔数是否正确。class文件的魔数：0xCAFEBABE （咖啡宝贝）
func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
}

// readAndCheckVersion
// readAndCheckVersion 检查版本号是否支持
func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {
	self.minorVersion = reader.readUint16()
	self.majorVersion = reader.readUint16()
	switch self.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if self.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}
