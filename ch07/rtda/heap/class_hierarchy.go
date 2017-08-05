package heap
/**判端self 是否可以赋值给other*/
func (self *Class) isAssignableFrom(other *Class) bool {
	//在三种情况下，S类型的引用值可以赋值给T类型：S 和T是同一类型；T是类且S是T的子类；或者T是接口且S实现了T接 口
	s, t := other, self
	if s == t {
		return true
	}
	if !t.IsInterface() {
		return s.IsSubClassOf(t)
	} else {
		return s.IsImplements(t)//s实现了t接口
	}
}
/**判断S是否是T的子类，实际上也就是判断T是否是S的（直接或 间接）超类*/
func (self *Class) IsSubClassOf(other *Class) bool {
	for c := self.superClass; c != nil; c = c.superClass {
		if c == other {
			return true
		}
	}
	return false
}

func (self *Class) IsImplements(iface *Class) bool {
	for c := self; c != nil; c = c.superClass {
		for _, i := range c.interfaces {
			if i == iface || i.isSubInterfaceOf(iface) {
				return true
			}
		}
	}
	return false
}

/**判断S是否实现了T接口，就看S或S的（直接或间接）超类是否 实现了某个接口T'，T'要么是T，要么是T的子接口。 */
func (self *Class) isSubInterfaceOf(iface *Class) bool {
	//每个接口（Class结构体）会持有父接口（Class结构体）
	for _, superInterface := range self.interfaces {
		//递归判断 所有的父接口中是否有 和iface完全相同的接口
		if superInterface == iface || superInterface.isSubInterfaceOf(iface) {
			return true
		}
	}
	return false
}

func (self *Class) IsSuperClassOf(other *Class) bool {
	return other.IsSubClassOf(self)
}