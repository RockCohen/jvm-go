package heap

import (
	"classfile"
	"fmt"
)

type Constant interface{}
type ConstantPool struct {
	class  *Class
	consts []Constant
}

func newConstantPool(class *Class, cfCp classfile.ConstantPool) *ConstantPool {
	cpCount := len(cfCp)
	consts := make([]Constant, cpCount)
	rtCp := &ConstantPool{class, consts}
	for i := 1; i < cpCount; i++ {
		cpInfo := cfCp[i]
		switch cpInfo.(type) {
		case *classfile.ConstantIntegerInfo:
			intInfo := cpInfo.(*classfile.ConstantIntegerInfo)
			consts[i] = intInfo.Value() // int32
		case *classfile.ConstantFloatInfo:
			floatInfo := cpInfo.(*classfile.ConstantFloatInfo)
			consts[i] = floatInfo.Value() // float32
		case *classfile.ConstantLongInfo:
			longInfo := cpInfo.(*classfile.ConstantLongInfo)
			consts[i] = longInfo.Value() // int64
			i++
		case *classfile.ConstantDoubleInfo:
			doubleInfo := cpInfo.(*classfile.ConstantDoubleInfo)
			consts[i] = doubleInfo.Value() // float64
			i++
		case *classfile.ConstantStringInfo:
			stringInfo := cpInfo.(*classfile.ConstantStringInfo)
			consts[i] = stringInfo.String() // string
		case *classfile.ConstantClassInfo:
			classInfo := cpInfo.(*classfile.ConstantClassInfo)
			consts[i] = newClassRef(rtCp, classInfo) // 见 6.2.1小节
		case *classfile.ConstantFieldrefInfo:
			fieldrefInfo := cpInfo.(*classfile.ConstantFieldrefInfo)
			consts[i] = newFieldRef(rtCp, fieldrefInfo) // 见 6.2.2小节
		case *classfile.ConstantMethodrefInfo:
			methodrefInfo := cpInfo.(*classfile.ConstantMethodrefInfo)
			consts[i] = newMethodRef(rtCp, methodrefInfo) // 见 6.2.3小节
		case *classfile.ConstantInterfaceMethodrefInfo:
			methodrefInfo := cpInfo.(*classfile.ConstantInterfaceMethodrefInfo)
			consts[i] = newInterfaceMethodRef(rtCp, methodrefInfo) // 见 6.2.4小节

		}
	}
	return rtCp
}
func (self *ConstantPool) GetConstant(index uint) Constant {
	if c := self.consts[index]; c != nil {
		return c
	}
	panic(fmt.Sprintf("No constants at index %d", index))
}
