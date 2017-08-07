package lang

import "unsafe"
import "jvmgo/ch10/native"
import "jvmgo/ch10/rtda"

func init() {
	native.Register("java/lang/Object", "getClass", "()Ljava/lang/Class;", getClass)
	native.Register("java/lang/Object", "hashCode", "()I", hashCode)
	native.Register("java/lang/Object", "clone", "()Ljava/lang/Object;", clone)
}
/*本地方法 把帧所属的类的Class对象推入操作数栈顶*/
func getClass(frame *rtda.Frame) {
	//从局部变量表中拿到this引用
	this := frame.LocalVars().GetThis()
	//通过Class()方法拿到它的Class结构体指针
	class := this.Class().JClass()
	//把类对象推 入操作数栈顶
	frame.OperandStack().PushRef(class)
}
/*Object hashCode 本地方法的实现*/
func hashCode(frame *rtda.Frame) {
	//从局部变量表中拿到this引用Object结构体指针
	this := frame.LocalVars().GetThis()
	//计算hash值
	hash := int32(uintptr(unsafe.Pointer(this)))
	//推入操作数栈顶
	frame.OperandStack().PushInt(hash)
}

func clone(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()
	cloneable := this.Class().Loader().LoadClass("java/lang/Cloneable")
	//如果类没有实现Cloneable接口,则抛出CloneNotSupportedException异常
	if !this.Class().IsImplements(cloneable) {
		panic("java.lang.CloneNotSupportedException")
	}
	//调用Object结构体的 Clone()方法克隆对象,把对象副本引用推入操作数栈顶。
	frame.OperandStack().PushRef(this.Clone())
}