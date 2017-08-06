package references

import "jvmgo/ch09/instructions/base"
import "jvmgo/ch09/rtda"
import "jvmgo/ch09/rtda/heap"

/**
instanceof指令判断对象是否是某个类的实例（或者对象的类是 否实现了某个接口），并把结果推入操作数栈
两个操作数
1.类符号引用 从字节码中获取
2.对象引用，从操作数栈中弹出
*/
type INSTANCE_OF struct {
	base.Index16Instruction
}

func (self *INSTANCE_OF) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	ref := stack.PopRef()
	if ref == nil { // 如果是null，则把0推入操作数栈。
		stack.PushInt(0)
		return
	}
	//如果对象引用不是null，则解析类符号引用，判断对象是否是类的实例
	cp :=  frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()

	if ref.IsInstanceOf(class) {
		stack.PushInt(1)
	} else {
		stack.PushInt(0)
	}
}