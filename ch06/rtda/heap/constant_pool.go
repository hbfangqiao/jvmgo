package heap

import "fmt"
import "jvmgo/ch06/classfile"

type Constant interface{}
/*运行时常量池：存放字面量和符号引用（类符号引用，字段符号引用，方法符号引用，接口方法符号引用）*/
type ConstantPool struct {
	class  *Class
	consts []Constant
}
/*把class文件中的常量池转换成运行时常量池*/
func newConstantPool(class *Class, cfCp classfile.ConstantPool) *ConstantPool{

}

/*根据索引返回常量*/
func (self *ConstantPool) GetConstant(index uint) Constant {
	if c := self.consts[index]; c != nil {
		return c
	}
	panic(fmt.Sprintf("No constants at index %d",index))
}
