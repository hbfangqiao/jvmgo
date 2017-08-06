package references

import "jvmgo/ch09/instructions/base"
import "jvmgo/ch09/rtda"
import "jvmgo/ch09/rtda/heap"

// Create new multidimensional array
/*multianewarray指令创建多维数组。
第一个操作数是个uint16索引，通过这个索引可以从运行时常量池中找到一个类符号引用，解析这个引用就 可以得到多维数组类
第二个操作数是个uint8整数，表示数组维度
这两个操作数在字节码中紧跟在指令操作码后面
multianewarray指令还需要从操作数栈中弹出n个整数，分别代表每一个维度的数组长度*/
type MULTI_ANEW_ARRAY struct {
	index      uint16
	dimensions uint8
}

func (self *MULTI_ANEW_ARRAY) FetchOperands(reader *base.BytecodeReader) {
	self.index = reader.ReadUint16()
	self.dimensions = reader.ReadUint8()
}
/*Execute() 方法根据数组类、数组维度和各个维度的数组长度创建多维数组*/
func (self *MULTI_ANEW_ARRAY) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(uint(self.index)).(*heap.ClassRef)
	//通过数组引用创建数组类-->创建多维数组的第一个结构体 int[3]型
	arrClass := classRef.ResolvedClass()

	stack := frame.OperandStack()
	//通过数组的维度，从操作数栈中弹出counts一个[]int32 数组
	counts := popAndCheckCounts(stack, int(self.dimensions))
	//通过每个维度数组的长度，已经数组类，创建多维数组
	arr := newMultiDimensionalArray(counts, arrClass)
	stack.PushRef(arr)
}
/*从操作数栈中弹出n个int值，并且确保它们都大于等于0。如果其中任何一个小于0，则抛出 NegativeArraySizeException异常*/
func popAndCheckCounts(stack *rtda.OperandStack, dimensions int) []int32 {
	counts := make([]int32, dimensions)
	for i := dimensions - 1; i >= 0; i-- {
		counts[i] = stack.PopInt()
		if counts[i] < 0 {
			panic("java.lang.NegativeArraySizeException")
		}
	}

	return counts
}
/*多维数组的int[3][2][4] int[3]中的每个元素都装的int[2] int[3][2]中的每个元素也是装的int[4]*/
func newMultiDimensionalArray(counts []int32, arrClass *heap.Class) *heap.Object {
	//取第一个counts，即当前数组的维度
	count := uint(counts[0])
	//根据第当前的维度，创建数组对象*Object型对象
	arr := arrClass.NewArray(count)
	//递归调用，创建[[[...[Ljava.lang.String;
	if len(counts) > 1 {
		//获取*Object中的data 转换为arr
		refs := arr.Refs()
		//将arr[]中的每一个元素填上一个 数组
		for i := range refs {
			//会在运行时动态生成多个Class结构体int[3][2][4] 在之前已经生成了int[3] 会再生成int[][] 和int[][][]
			refs[i] = newMultiDimensionalArray(counts[1:], arrClass.ComponentClass())
		}
	}

	return arr
}
