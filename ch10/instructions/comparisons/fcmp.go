package comparisons
import "jvmgo/ch10/instructions/base"
import "jvmgo/ch10/rtda"
// Compare float
type FCMPG struct{ base.NoOperandsInstruction }
type FCMPL struct{ base.NoOperandsInstruction }

func (self *FCMPG) Execute(frame *rtda.Frame) {
	//当两个float变量中至少有一个是NaN时,用fcmpg指 令比较的结果是1
	_fcmp(frame, true)
}
func (self *FCMPL) Execute(frame *rtda.Frame) {
	//当两个float变量中至少有一个是NaN时,用fcmpl指令比较的结果是-1
	_fcmp(frame, false)
}


func _fcmp(frame *rtda.Frame, gFlag bool) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else if v1 < v2 {
		stack.PushInt(-1)
	} else if gFlag {//当有一个数是NAN 不能比较时
		stack.PushInt(1)
	} else {
		stack.PushInt(-1)
	}
}
