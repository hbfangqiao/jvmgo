package lang

import "jvmgo/ch10/native"
import "jvmgo/ch10/rtda"
import "jvmgo/ch10/rtda/heap"

func init() {
	native.Register("java/lang/Class", "getPrimitiveClass", "(Ljava/lang/String;)Ljava/lang/Class;", getPrimitiveClass)
	native.Register("java/lang/Class", "getName0", "()Ljava/lang/String;", getName0)
	native.Register("java/lang/Class", "desiredAssertionStatus0","(Ljava/lang/Class;)Z", desiredAssertionStatus0)
}
/*将帧的局部变量表的第一个引用的类的类对象推入操作数栈顶*/
func getPrimitiveClass(frame *rtda.Frame) {
	//先从局部变量表中拿到类名
	nameObj := frame.LocalVars().GetRef(0)
	//转换为Go字符串
	name := heap.GoString(nameObj)
	//调用类加载器的LoadClass()方法获取基本类型的类对象
	loader := frame.Method().Class().Loader()
	class := loader.LoadClass(name).JClass()
	frame.OperandStack().PushRef(class)
}
/*将frame对应的Class的类名推入操作数栈顶*/
func getName0(frame *rtda.Frame) {
	//从局部变量表中拿到this引用,是一个类对象引用
	this := frame.LocalVars().GetThis()
	//获得与之对应的Class结构体指针
	class := this.Extra().(*heap.Class)
	//获取类名
	name := class.JavaName()
	nameObj := heap.JString(class.Loader(), name)
	frame.OperandStack().PushRef(nameObj)
}

/*把false推入操作数栈顶*/
func desiredAssertionStatus0(frame *rtda.Frame) {
	frame.OperandStack().PushBoolean(false)
}