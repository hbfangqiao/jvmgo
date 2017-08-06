package lang

import "jvmgo/ch09/native"
import "jvmgo/ch09/rtda"

func init() {
	native.Register("java/lang/Object", "getClass", "()Ljava/lang/Class;", getClass)
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