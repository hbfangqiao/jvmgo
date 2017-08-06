package references

import "jvmgo/ch07/instructions/base"
import "jvmgo/ch07/rtda"
import "jvmgo/ch07/rtda/heap"
/*
NEW 指令：
1.从当前类的运行时常量池中找到一个类符号引用
2.解析这个类符号引用，拿到类数据，创建对象
3.把对象引用推入栈顶
 */
type NEW struct {
	//操作数是一个uint16索引，来自字节码。作用是从当前类的运行时常量池中找到一个类符号引用
	base.Index16Instruction
}

func (self *NEW) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)//根据自己的操作数，拿到一个符号引用
	class := classRef.ResolvedClass()
	//类还没初始化，则先初始化类
	if !class.InitStarted() {
		frame.RevertNextPC()//将指令重新指向当前指令
		base.InitClass(frame.Thread(), class)//初始化类
		return//终止执行当前指令
	}

	if class.IsInterface() || class.IsAbstract() {
		//接口和抽象类不能实例化，根据虚拟机规定抛出InstantiationError异常
		panic("java.lang.InstantiationError")
	}
	ref := class.NewObject()
	frame.OperandStack().PushRef(ref)
}