package heap
/*如何知道静态变量和实例变量需要多少空 间，以及哪个字段对应Slots中的哪个位置呢？
第一个问题比较好解决，只要数一下类的字段即可。假设某个 类有m个静态字段和n个实例字段，那么静态变量和实例变量所需 的空间大小就分别是m'和n'。这里要注意两点。首先，类是可以继承
的。也就是说，在数实例变量时，要递归地数超类的实例变量；其次，long和double字段都占据两个位置，所以m'>=m，n'>=n。
第二个问题也不算难，在数字段时，给字段按顺序编上号就可
以了。这里有三点需要要注意。首先，静态字段和实例字段要分开
编号，否则会混乱。其次，对于实例字段，一定要从继承关系的最 顶端，也就是java.lang.Object开始编号，否则也会混乱。最后，编号 时也要考虑long和double类型。
*/
type Object struct {
	class *Class      //对象的Class指针
	data  interface{} //实例变量
	extra interface{} //extra字段用来记录Object结构体实例的额外信息,暂时用它来记录类对象对应的Class结构体指针
}

func newObject(class *Class) *Object {
	return &Object{
		class: class,
		data:  newSlots(class.instanceSlotCount),
	}
}

// getters

func (self *Object) Class() *Class {
	return self.class
}

func (self *Object) Fields() Slots {
	return self.data.(Slots)
}

func (self *Object) Extra() interface{} {
	return self.extra
}

func (self *Object) SetExtra(extra interface{}) {
	self.extra = extra
}

func (self *Object) IsInstanceOf(class *Class) bool {
	return class.isAssignableFrom(self.class)
}

// reflection
func (self *Object) GetRefVar(name, descriptor string) *Object {
	field := self.class.getField(name, descriptor, false)
	slots := self.data.(Slots)
	return slots.GetRef(field.slotId)
}
/*/*通过域名，修饰符，给对象的引用类型实例变量赋值为ref*/
func (self *Object) SetRefVar(name, descriptor string, ref *Object) {
	field := self.class.getField(name, descriptor, false)
	slots := self.data.(Slots)
	slots.SetRef(field.slotId, ref)
}