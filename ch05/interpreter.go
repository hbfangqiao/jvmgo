package main

import "fmt"
import "jvmgo/ch05/classfile"
import "jvmgo/ch05/instructions"
import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

func interpret(methodInfo *classfile.MemberInfo) {
	//获取MemberInfo的Code属性
	codeAttr := methodInfo.CodeAttribute()
	//获取执行方法所需要的局部变量表
	maxLocals := codeAttr.MaxLocals()
	//获取执行方法所需要的栈空间
	maxStack := codeAttr.MaxStack()
	//获取方法的字节码
	bytecode := codeAttr.Code()
	//新建栈
	thread := rtda.NewThread()
	frame := thread.NewFrame(maxLocals,maxStack)
	thread.PushFrame(frame)

	defer catchErr(frame)
	loop(thread,bytecode)
}
/**
解释器目前还没有办法优雅地结 束运行。因为每个方法的最后一条指令都是某个return指令，而还 没有实现return指令，
所以方法在执行过程中必定会出现错误，此 时解释器逻辑会转到catchErr）函数，
 */
func catchErr(frame *rtda.Frame) {
	if r := recover(); r != nil {
		fmt.Printf("LocalVars:%v\n",frame.LocalVars())
		fmt.Printf("OperandStack:%v\n",frame.OperandStack())
		panic(r)
	}
}

func loop(thread *rtda.Thread,bytecode []byte) {
	frame := thread.PopFrame()
	reader := &base.BytecodeReader{}

	for {
		//计算pc
		pc := frame.NextPC()
		thread.SetPC(pc)
		//解码指令
		reader.Reset(bytecode,pc)
		opcode := reader.ReadUint8()
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())
		//执行指令
		fmt.Printf("pc:%2d inst:%T %v\n",pc,inst,inst)
		inst.Execute(frame)
	}
}
