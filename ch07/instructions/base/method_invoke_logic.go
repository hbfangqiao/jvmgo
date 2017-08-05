package base

import "fmt"
import "jvmgo/ch07/rtda"
import "jvmgo/ch07/rtda/heap"



func InvokeMethod(invokerFrame *rtda.Frame, method *heap.Method) {
	//创建新的帧并推入Java虚拟机栈
	thread := invokerFrame.Thread()
	newFrame := thread.NewFrame(method)
	thread.PushFrame(newFrame)
	//传递参数 要确定方法的参数在局部 变量表中占用多少位置。注意，这个数量并不一定等于从Java代码 中看到的参数个数，原因有两个。第一，long和double类型的参数要
	//占用两个位置。第二，对于实例方法，Java编译器会在参数列表的 前面添加一个参数，这个隐藏的参数就是this引用。假设实际的参 数占据n个位置，依次把这n个变量从调用者的操作数栈中弹出，放
	//进被调用方法的局部变量表中，参数传递就完成了。

	argSlotSlot := int(method.ArgSlotCount())
	if argSlotSlot > 0 {
		for i := argSlotSlot - 1; i >= 0; i-- {
			slot := invokerFrame.OperandStack().PopSlot()
			newFrame.LocalVars().SetSlot(uint(i), slot)
		}
	}
	
	// hack!
	if method.IsNative() {
		if method.Name() == "registerNatives" {
			thread.PopFrame()
		} else {
			panic(fmt.Sprintf("native method: %v.%v%v\n",
				method.Class().Name(), method.Name(), method.Descriptor()))
		}
	}
}
