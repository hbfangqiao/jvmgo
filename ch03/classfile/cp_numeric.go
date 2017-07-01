package classfile

import "math"

/*存放java中int型常量，比int更小的boolean,byte,short,char也存放在其中*/
type ConstantIntegerInfo struct {
	val int32
}

func (self *ConstantIntegerInfo) readInfo(reader *ClassReader){
	bytes := reader.readUnit32()
	self.val = int32(bytes)
}

type ConstantFloatInfo struct {
	val float32
}

func (self *ConstantFloatInfo) readInfo (reader *ClassReader) {
	bytes := reader.readUnit32()
	self.val = math.Float32frombits(bytes)
}

type ConstantLongInfo struct {
	val int64
}

func (self *ConstantLongInfo) readInfo (reader *ClassReader) {
	bytes := reader.readUnit64()
	self.val = int64(bytes)
}

type ConstantDoubleInfo struct {
	val float64
}

func (self *ConstantDoubleInfo) readInfo (reader *ClassReader) {
	bytes := reader.readUnit64()
	self.val = math.Float64frombits(bytes)
}

