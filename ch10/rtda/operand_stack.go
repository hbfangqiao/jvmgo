package rtda

import "math"
import "jvmgo/ch10/rtda/heap"

type OperandStack struct {
	size  uint //栈顶位置
	slots []Slot
}

func newOperandStack(maxStack uint) *OperandStack {
	if maxStack > 0 {
		return &OperandStack{
			slots: make([]Slot, maxStack),
		}
	}
	return nil
}

/*往栈顶放一个int变量*/
func (self *OperandStack) PushInt(val int32) {
	self.slots[self.size].num = val
	self.size++
}
/*返回栈顶的int变量*/
func (self *OperandStack) PopInt() int32 {
	self.size--
	return self.slots[self.size].num
}

/*往栈顶push一个java float变量*/
func (self *OperandStack) PushFloat(val float32) {
	bits := math.Float32bits(val)
	self.slots[self.size].num = int32(bits)
	self.size++
}
func (self *OperandStack) PopFloat() float32 {
	self.size--
	bits := uint32(self.slots[self.size].num)
	return math.Float32frombits(bits)
}
/*向栈顶push一个java long变量*/
func (self *OperandStack) PushLong(val int64) {
	self.slots[self.size].num = int32(val)
	self.slots[self.size + 1].num = int32(val >> 32)
	self.size += 2
}
func (self *OperandStack) PopLong() int64 {
	self.size -= 2
	low := uint32(self.slots[self.size].num)
	high := uint32(self.slots[self.size + 1].num)
	return int64(high) << 32 | int64(low)
}
/*往栈顶push一个double类型*/
func (self *OperandStack) PushDouble(val float64) {
	bits := math.Float64bits(val)
	self.PushLong(int64(bits))
}
func (self *OperandStack) PopDouble() float64 {
	bits := uint64(self.PopLong())
	return math.Float64frombits(bits)
}

func (self *OperandStack) PushRef(ref *heap.Object) {
	self.slots[self.size].ref = ref
	self.size++
}
func (self *OperandStack) PopRef() *heap.Object {
	self.size--
	ref := self.slots[self.size].ref
	self.slots[self.size].ref = nil//帮助Go的垃圾回收器回收heap.Object结构体实例
	return ref
}

func (self *OperandStack) PushSlot(slot Slot) {
	self.slots[self.size] = slot
	self.size++
}
func (self *OperandStack) PopSlot() Slot {
	self.size --
	return self.slots[self.size]
}
/*0则获取栈顶Slot，1获取倒数第二个Slot*/
func (self *OperandStack) GetRefFromTop(n uint) *heap.Object {
	return self.slots[self.size - 1 - n].ref
}

func (self *OperandStack) PushBoolean(val bool) {
	if val {
		self.PushInt(1)
	} else {
		self.PushInt(0)
	}
}

func (self *OperandStack) PopBoolean() bool {
	return self.PopInt() == 1
}

func (self *OperandStack) Clear() {
	self.size = 0
	for i := range self.slots {
		self.slots[i].ref = nil
	}
}