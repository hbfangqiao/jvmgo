package control

import "jvmgo/ch08/instructions/base"
import "jvmgo/ch08/rtda"

// Return void from method
type RETURN struct{ base.NoOperandsInstruction }
//把当前帧从Java虚拟机栈中弹出即可
func (self *RETURN) Execute(frame *rtda.Frame) {
	frame.Thread().PopFrame()
}

// Return reference from method
type ARETURN struct{ base.NoOperandsInstruction }

func (self *ARETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	ref := currentFrame.OperandStack().PopRef()
	invokerFrame.OperandStack().PushRef(ref)
}

// Return double from method
type DRETURN struct{ base.NoOperandsInstruction }

func (self *DRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	val := currentFrame.OperandStack().PopDouble()
	invokerFrame.OperandStack().PushDouble(val)
}

// Return float from method
type FRETURN struct{ base.NoOperandsInstruction }

func (self *FRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()//弹出当前帧
	invokerFrame := thread.TopFrame()//获取栈顶帧
	val := currentFrame.OperandStack().PopFloat()//将当前帧计算的float值弹出
	invokerFrame.OperandStack().PushFloat(val)//将计算出的float推入栈顶帧
}

// Return int from method
type IRETURN struct{ base.NoOperandsInstruction }

func (self *IRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	val := currentFrame.OperandStack().PopInt()
	invokerFrame.OperandStack().PushInt(val)
}

// Return double from method
type LRETURN struct{ base.NoOperandsInstruction }

func (self *LRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	val := currentFrame.OperandStack().PopLong()
	invokerFrame.OperandStack().PushLong(val)
}
