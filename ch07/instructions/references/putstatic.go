package references

import "jvmgo/ch07/instructions/base"
import "jvmgo/ch07/rtda"
import "jvmgo/ch07/rtda/heap"

/**给类的某个静态变量赋值，需要两个操作数，
第一个操作数是uint16 来自字节码。通过这个索引可以从当前类的运行时常量池中找到一个字段符号引用，解析这个符号引用
就可以知道要给类的哪个静态变量赋值。
第二个操作数是要赋给静 态变量的值，从操作数栈中弹出。*/
type PUT_STATIC struct {
	base.Index16Instruction
}

func (self *PUT_STATIC) Execute(frame *rtda.Frame) {
	currentMethod := frame.Method()//获取当前方法
	currentClass := currentMethod.Class()//获取当前类
	cp := currentClass.ConstantPool()//获取常量池
	fieldRef := cp.GetConstant(self.Index).(*heap.FieldRef)//解析字段符号引用，获取需要操作的静态变量
	field := fieldRef.ResolvedField()
	class := field.Class()

	//类还没初始化，则先初始化类
	if !class.InitStarted() {
		frame.RevertNextPC()//将指令重新指向当前指令
		base.InitClass(frame.Thread(), class)//初始化类
		return//终止执行当前指令
	}


	if !field.IsStatic() {//解析后的字段是实例字段
		panic("java.lang.IncompatibleClassChangeError")
	}
	if field.IsFinal() {//如果是final字段，这里实际操作的是静态常量，只能在类初始化方法中给它赋值
		if currentClass != class || currentMethod.Name() != "<clinit>" {//类初始化方法<clinit> 由编译器生成
			panic("java.lang.IllegalAccessError")
		}
	}

	//根据字段类型从操作数栈中弹出相应的值，然后赋给静态变量
	descriptor := field.Descriptor()
	slotId := field.SlotId()
	slots := class.StaticVars()
	stack := frame.OperandStack()
	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I': slots.SetInt(slotId, stack.PopInt())
	case 'F': slots.SetFloat(slotId, stack.PopFloat())
	case 'J': slots.SetLong(slotId, stack.PopLong())
	case 'D': slots.SetDouble(slotId, stack.PopDouble())
	case 'L', '[': slots.SetRef(slotId, stack.PopRef())

	}
}