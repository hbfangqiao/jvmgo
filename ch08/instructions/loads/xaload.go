package loads

import "jvmgo/ch08/instructions/base"
import "jvmgo/ch08/rtda"
import "jvmgo/ch08/rtda/heap"

/*<t>aload系列指令按索引取数组元素值，然后推入操作数栈*/

// Load reference from array
type AALOAD struct{ base.NoOperandsInstruction }

func (self *AALOAD) Execute(frame *rtda.Frame) {
	//从操作数栈中弹出第一个操作数：数组索引
	stack := frame.OperandStack()
	index := stack.PopInt()
	//弹出第二个操作数：数组引用
	arrRef := stack.PopRef()
	//如果数组引用是null，则抛出 NullPointerException异常
	checkNotNil(arrRef)
	//根据数组引用，返回引用数组
	refs := arrRef.Refs()
	//如果数组索引小于0，或者大于等于数组长度，则抛出 ArrayIndexOutOfBoundsException
	checkIndex(len(refs), index)
	//如果一切正常，则按索引取出数组元素，推入操作数栈顶
	stack.PushRef(refs[index])
}

// Load byte or boolean from array
type BALOAD struct{ base.NoOperandsInstruction }

func (self *BALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	bytes := arrRef.Bytes()
	checkIndex(len(bytes), index)
	stack.PushInt(int32(bytes[index]))
}

// Load char from array
type CALOAD struct{ base.NoOperandsInstruction }

func (self *CALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	chars := arrRef.Chars()
	checkIndex(len(chars), index)
	stack.PushInt(int32(chars[index]))
}

// Load double from array
type DALOAD struct{ base.NoOperandsInstruction }

func (self *DALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	doubles := arrRef.Doubles()
	checkIndex(len(doubles), index)
	stack.PushDouble(doubles[index])
}

// Load float from array
type FALOAD struct{ base.NoOperandsInstruction }

func (self *FALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	floats := arrRef.Floats()
	checkIndex(len(floats), index)
	stack.PushFloat(floats[index])
}

// Load int from array
type IALOAD struct{ base.NoOperandsInstruction }

func (self *IALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	ints := arrRef.Ints()
	checkIndex(len(ints), index)
	stack.PushInt(ints[index])
}

// Load long from array
type LALOAD struct{ base.NoOperandsInstruction }

func (self *LALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	longs := arrRef.Longs()
	checkIndex(len(longs), index)
	stack.PushLong(longs[index])
}

// Load short from array
type SALOAD struct{ base.NoOperandsInstruction }

func (self *SALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	shorts := arrRef.Shorts()
	checkIndex(len(shorts), index)
	stack.PushInt(int32(shorts[index]))
}

func checkNotNil(ref *heap.Object) {
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
}
func checkIndex(arrLen int, index int32) {
	if index < 0 || index >= int32(arrLen) {
		panic("ArrayIndexOutOfBoundsException")
	}
}
