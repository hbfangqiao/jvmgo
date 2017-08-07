package references

import "jvmgo/ch10/instructions/base"
import "jvmgo/ch10/rtda"
import "jvmgo/ch10/rtda/heap"

// Invoke a class (static) method
type INVOKE_STATIC struct{ base.Index16Instruction }

func (self *INVOKE_STATIC) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	methodRef := cp.GetConstant(self.Index).(*heap.MethodRef)
	resolvedMethod := methodRef.ResolvedMethod()
	//resolvedMethod不能是类初始化方法。类 初始化方法只能由Java虚拟机调用，不能使用invokestatic指令调用。
	// 这一规则由class文件验证器保证，这里不做检查
	if !resolvedMethod.IsStatic() {//不是静态方法
		panic("java.lang.IncompatibleClassChangeError")
	}


	class := resolvedMethod.Class()
	//类还没初始化，则先初始化类
	if !class.InitStarted() {
		frame.RevertNextPC()//将指令重新指向当前指令
		base.InitClass(frame.Thread(), class)//初始化类
		return//终止执行当前指令
	}

	base.InvokeMethod(frame, resolvedMethod)
}
