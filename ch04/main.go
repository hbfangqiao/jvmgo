package main

import "fmt"
import "jvmgo/ch04/rtda"

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {//如果输入了-version
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" { //如果输了了-help
		//printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	frame := rtda.NewFrame(100,100)
	testLocalVars(frame.LocalVars())
	testOperandStack(frame.OperandStack())

}

func testLocalVars(vars rtda.LocalVars) {
	vars.SetInt(0,100)
	vars.SetInt(1,-100)
	vars.SetLong(2,2999999333)
	vars.SetLong(4,-1222223233)
	vars.SetFloat(6,3.1415926)
	vars.SetDouble(7,2.1918238132)
	vars.SetRef(9,nil)
	println(vars.GetInt(0))
	println(vars.GetInt(1))
	println(vars.GetLong(2))
	println(vars.GetLong(4))
	println(vars.GetFloat(6))
	println(vars.GetDouble(7))
	println(vars.GetRef(9))
}

func testOperandStack(ops *rtda.OperandStack) {
	ops.PushInt(100)
	ops.PushLong(-100)
	ops.PushLong(29999991919)
	ops.PushLong(-20198199828)
	ops.PushFloat(3.11231231)
	ops.PushDouble(2.12312300123)
	ops.PushRef(nil)
	println(ops.PopRef())
	println(ops.PopDouble())
	println(ops.PopFloat())
	println(ops.PopLong())
	println(ops.PopLong())
	println(ops.PopInt())
	println(ops.PopInt())
}
