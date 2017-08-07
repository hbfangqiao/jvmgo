package lang

import "fmt"
import "jvmgo/ch10/native"
import "jvmgo/ch10/rtda"
import "jvmgo/ch10/rtda/heap"

const jlThrowable = "java/lang/Throwable"

/*记录Java虚拟机栈帧信息*/
type StackTraceElement struct {
	fileName   string //类所 在的文件名
	className  string //声明方法的类名
	methodName string //方法名
	lineNumber int    //帧正在执行哪行代码
}

func (self *StackTraceElement) String() string {
	return fmt.Sprintf("%s.%s(%s:%d)",
		self.className, self.methodName, self.fileName, self.lineNumber)
}

func init() {
	native.Register(jlThrowable, "fillInStackTrace", "(I)Ljava/lang/Throwable;", fillInStackTrace)
}

// private native Throwable fillInStackTrace(int dummy);
// (I)Ljava/lang/Throwable;
func fillInStackTrace(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()
	frame.OperandStack().PushRef(this)

	stes := createStackTraceElements(this, frame.Thread())
	this.SetExtra(stes)
}

/*由于栈顶两帧正在执行 fillInStackTrace（int）和fillInStackTrace（）方法，所以需要跳过这两
帧。这两帧下面的几帧正在执行异常类的构造函数，所以也要跳
过，具体要跳过多少帧数则要看异常类的继承层次*/
func createStackTraceElements(tObj *heap.Object, thread *rtda.Thread) []*StackTraceElement {
	//跳过异常类的构造函数，fillInStackTrace（int）和fillInStackTrace（）方法
	skip := distanceToObject(tObj.Class()) + 2
	//拿到完整的Java虚拟机栈
	frames := thread.GetFrames()[skip:]
	stes := make([]*StackTraceElement, len(frames))
	for i, frame := range frames {
		stes[i] = createStackTraceElement(frame)
	}
	return stes
}

/*计算跳过的帧数*/
func distanceToObject(class *heap.Class) int {
	distance := 0
	for c := class.SuperClass(); c != nil; c = c.SuperClass() {
		distance++
	}
	return distance
}
/*根据帧创建StackTraceElement实例*/
func createStackTraceElement(frame *rtda.Frame) *StackTraceElement {
	method := frame.Method()
	class := method.Class()
	return &StackTraceElement{
		fileName:   class.SourceFile(),
		className:  class.JavaName(),
		methodName: method.Name(),
		lineNumber: method.GetLineNumber(frame.NextPC() - 1),
	}
}
