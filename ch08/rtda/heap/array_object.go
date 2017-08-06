package heap
//别针对引用类型数组和7种基本类型数组返 回具体的数组数据。
func (self *Object) Bytes() []int8 {
	return self.data.([]int8)
}
func (self *Object) Shorts() []int16 {
	return self.data.([]int16)
}
func (self *Object) Ints() []int32 {
	return self.data.([]int32)
}
func (self *Object) Longs() []int64 {
	return self.data.([]int64)
}
func (self *Object) Chars() []uint16 {
	return self.data.([]uint16)
}
func (self *Object) Floats() []float32 {
	return self.data.([]float32)
}
func (self *Object) Doubles() []float64 {
	return self.data.([]float64)
}
func (self *Object) Refs() []*Object {
	return self.data.([]*Object)
}
//，为什么返回数组数据的方法有8个，但却只 有一个统一的ArrayLength（）方法呢？
// 答案是，这些方法主要是供 <t>aload、<t>astore和arraylength指令使用的。<t>aload和<t>astore系 列指令各有8条，
// 所以针对每种类型都提供一个方法，返回相应的 数组数据。因为arraylength指令只有一条，所以ArrayLength（）方法 需要
// 自己判断数组类型。

func (self *Object) ArrayLength() int32 {
	switch self.data.(type) {
	case []int8:
		return int32(len(self.data.([]int8)))
	case []int16:
		return int32(len(self.data.([]int16)))
	case []int32:
		return int32(len(self.data.([]int32)))
	case []int64:
		return int32(len(self.data.([]int64)))
	case []uint16:
		return int32(len(self.data.([]uint16)))
	case []float32:
		return int32(len(self.data.([]float32)))
	case []float64:
		return int32(len(self.data.([]float64)))
	case []*Object:
		return int32(len(self.data.([]*Object)))
	default:
		panic("Not array!")
	}
}
