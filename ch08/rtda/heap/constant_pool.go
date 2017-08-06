package heap

import "fmt"
import "jvmgo/ch08/classfile"

type Constant interface{}
/*运行时常量池：存放字面量和符号引用（类符号引用，字段符号引用，方法符号引用，接口方法符号引用）*/
type ConstantPool struct {
	class  *Class
	consts []Constant
}
/*把class文件中的常量池转换成运行时常量池*/
func newConstantPool(class *Class, cfCp classfile.ConstantPool) *ConstantPool {
	cpCount := len(cfCp)//获取classfile中ConstantPool的长度
	consts := make([]Constant, cpCount)//创建运行时常量池所包含的常量切片
	rtCp := &ConstantPool{class, consts}//创建运行时常量池，指向自身所属类
	for i := 1; i < cpCount; i++ {
		//根据cf ConstantPool的实际类型做相应的处理
		cpInfo := cfCp[i]
		switch cpInfo.(type) {
		case *classfile.ConstantIntegerInfo ://int常量
			intInfo := cpInfo.(*classfile.ConstantIntegerInfo)//转为ConstantIntegerInfo
			consts[i] = intInfo.Value() //int32
		case *classfile.ConstantFloatInfo ://float常量
			floatInfo := cpInfo.(*classfile.ConstantFloatInfo)
			consts[i] = floatInfo
		case *classfile.ConstantLongInfo://long常量
			//ConstantLongInfo和ConstantDoubleInfo占classfile.ConstantPool的两个索引，故i需要额外自增一次
			longInfo := cpInfo.(*classfile.ConstantLongInfo)
			consts[i] = longInfo.Value()//int64
			i++
		case *classfile.ConstantDoubleInfo://double常量
			doubleInfo := cpInfo.(*classfile.ConstantDoubleInfo)
			consts[i] = doubleInfo.Value()//float64
			i++
		case *classfile.ConstantStringInfo://string常量
			stringInfo := cpInfo.(*classfile.ConstantStringInfo)
			consts[i] = stringInfo.String()//string
		case *classfile.ConstantClassInfo:
			classInfo := cpInfo.(*classfile.ConstantClassInfo)
			consts[i] = newClassRef(rtCp, classInfo)
		case *classfile.ConstantFieldrefInfo:
			fieldrefInfo := cpInfo.(*classfile.ConstantFieldrefInfo)
			consts[i] = newFieldRef(rtCp, fieldrefInfo)
		case *classfile.ConstantMethodrefInfo:
			methodrefInfo := cpInfo.(*classfile.ConstantMethodrefInfo)
			consts[i] = newMethodRef(rtCp, methodrefInfo)
		case *classfile.ConstantInterfaceMethodrefInfo:
			methodrefInfo := cpInfo.(*classfile.ConstantInterfaceMethodrefInfo)
			consts[i] = newInterfaceMethodRef(rtCp, methodrefInfo)
		}
	}
	return rtCp
}

/*根据索引返回常量*/
func (self *ConstantPool) GetConstant(index uint) Constant {
	if c := self.consts[index]; c != nil {
		return c
	}
	panic(fmt.Sprintf("No constants at index %d", index))
}
