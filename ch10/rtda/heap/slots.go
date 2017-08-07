package heap

import "math"
/*表示类变量和实例变量*/
type Slot struct {
	num int32
	ref *Object
}

type Slots []Slot 

func newSlots(slotCount uint) Slots {
	return make([]Slot,slotCount)
}

/*存取Int变量*/
func (self Slots) SetInt(index uint, val int32) {
	self[index].num = val
}
func (self Slots) GetInt(index uint) int32 {
	return self[index].num
}

/*存取java float变量,将float转变int存取*/
func (self Slots) SetFloat(index uint, val float32) {
	bits := math.Float32bits(val)
	self[index].num = int32(bits)
}
func (self Slots) GetFloat(index uint) float32 {
	bits := uint32(self[index].num)
	return math.Float32frombits(bits)
}

/*存取long变量,需要用2个Slot存取*/
func (self Slots) SetLong(index uint,val int64){
	self[index].num = int32(val)
	self[index+1].num = int32(val >> 32)
}
func (self Slots) GetLong(index uint) int64 {
	low := uint32(self[index].num)
	high := uint32(self[index+1].num)
	return int64(high)<<32 | int64(low)
}

/*存取double变量,转成long类型，按照long变量处理*/
func (self Slots) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	self.SetLong(index, int64(bits))
}
func (self Slots) GetDouble(index uint) float64 {
	bits := uint64(self.GetLong(index))
	return math.Float64frombits(bits)
}

/*引用值，直接存取*/
func (self Slots) SetRef(index uint,ref *Object) {
	self[index].ref = ref
}
func (self Slots) GetRef(index uint) *Object {
	return self[index].ref
}