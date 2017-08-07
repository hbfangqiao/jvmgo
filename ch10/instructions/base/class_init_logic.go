package base

import "jvmgo/ch10/rtda"
import "jvmgo/ch10/rtda/heap"

func InitClass(thread *rtda.Thread, class *heap.Class) {
	//把类的initStarted状态设 置成true以免进入死循环
	class.StartInit()
	//准备执行类的初始化方法，把本类的 <clinit> 方法推入栈中
	scheduleClinit(thread, class)
	//初始化超类，将超类的<clinit>方法推入栈中---递归调用
	initSuperClass(thread, class)
}

func scheduleClinit(thread *rtda.Thread, class *heap.Class) {
	//获取类的<clinit>方法
	clinit := class.GetClinitMethod()
	if clinit != nil {
		// 把本类的 <clinit> 方法推入栈中
		newFrame := thread.NewFrame(clinit)
		thread.PushFrame(newFrame)
	}
}

func initSuperClass(thread *rtda.Thread, class *heap.Class) {
	if !class.IsInterface() {//初始化的类不是接口
		superClass := class.SuperClass()
		//拥有超类，并且超类没有初始化，则递归调用初始化超类
		if superClass != nil && !superClass.InitStarted() {
			InitClass(thread, superClass)
		}
	}
}