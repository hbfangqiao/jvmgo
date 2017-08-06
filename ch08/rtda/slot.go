package rtda

import "jvmgo/ch08/rtda/heap"
/*根据虚拟机规范，局部变量表的每个元素Slot至少可以容纳一个int或引用值*/
type Slot struct {
	num	int32//存放整数
	ref	*heap.Object//存放引用
}
