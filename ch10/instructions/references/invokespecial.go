package references

import "jvmgo/ch10/instructions/base"
import "jvmgo/ch10/rtda"
import "jvmgo/ch10/rtda/heap"

// Invoke instance method;
// special handling for superclass, private, and instance initialization method invocations
type INVOKE_SPECIAL struct{ base.Index16Instruction }

func (self *INVOKE_SPECIAL) Execute(frame *rtda.Frame) {
	//先拿到当前类、当前常量池、方法符号引用，然后解析符号引
	//用，拿到解析后的类和方法
	currentClass := frame.Method().Class()//当前类
	cp := currentClass.ConstantPool()//当前常量池
	methodRef := cp.GetConstant(self.Index).(*heap.MethodRef)//方法符号引用
	resolvedClass := methodRef.ResolvedClass()//方法引用所属的类
	resolvedMethod := methodRef.ResolvedMethod()//解析后的方法

	if resolvedMethod.Name() == "<init>" && resolvedMethod.Class() != resolvedClass {
		//如果是构造方法，方法所属的类，必须是方法引用所属的类
		panic("java.lang.NoSuchMethodError")
	}
	if resolvedMethod.IsStatic() {
		//是静态方法，则抛出IncompatibleClassChangeError异常
		panic("java.lang.IncompatibleClassChangeError")
	}
	//从操作数栈弹出this引用
	ref := frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount() - 1)
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
	//确保protected方法只能被声明该方法的类或子类调用
	if resolvedMethod.IsProtected() &&
		resolvedMethod.Class().IsSuperClassOf(currentClass) &&
		resolvedMethod.Class().GetPackageName() != currentClass.GetPackageName() &&
		ref.Class() != currentClass &&
		!ref.Class().IsSubClassOf(currentClass) {

		panic("java.lang.IllegalAccessError")
	}

	methodToBeInvoked := resolvedMethod
	if currentClass.IsSuper() &&
		resolvedClass.IsSuperClassOf(currentClass) &&
		resolvedMethod.Name() != "<init>" {
		//当前类ACC_SUPER被设置，要调用的是超类中的方法，并且不是构造方法则
		methodToBeInvoked = heap.LookupMethodInClass(currentClass.SuperClass(),
			methodRef.Name(), methodRef.Descriptor())
	}
	//如果查找过程失败，或者找到的方法是抽象的，抛出 AbstractMethodError异常。
	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}

	base.InvokeMethod(frame, methodToBeInvoked)
}
