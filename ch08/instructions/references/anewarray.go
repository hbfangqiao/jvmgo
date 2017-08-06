package references

import "jvmgo/ch08/instructions/base"
import "jvmgo/ch08/rtda"
import "jvmgo/ch08/rtda/heap"

// Create new array of reference
/*anewarray指令也需要两个操作数。第一个操作数是uint16索
引，来自字节码。通过这个索引可以从当前类的运行时常量池中找
到一个类符号引用，解析这个符号引用就可以得到数组元素的类。
第二个操作数是数组长度，从操作数栈中弹出*/
type ANEW_ARRAY struct{ base.Index16Instruction }

func (self *ANEW_ARRAY) Execute(frame *rtda.Frame) {
	//获取运行帧所属的类的常量池
	cp := frame.Method().Class().ConstantPool()
	//从常量池中获取数组元素的类的ref
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	//解析这个ref，获取数组元素的类
	componentClass := classRef.ResolvedClass()

	// if componentClass.InitializationNotStarted() {
	// 	thread := frame.Thread()
	// 	frame.SetNextPC(thread.PC()) // undo anewarray
	// 	thread.InitClass(componentClass)
	// 	return
	// }

	stack := frame.OperandStack()
	count := stack.PopInt()
	if count < 0 {
		panic("java.lang.NegativeArraySizeException")
	}
	//根据数组元素的类创建数组类 Ljava.lang.String; -> [Ljava.lang.String;
	arrClass := componentClass.ArrayClass()
	arr := arrClass.NewArray(uint(count))
	stack.PushRef(arr)
}
