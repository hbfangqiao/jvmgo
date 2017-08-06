package heap
/**判端self 是否可以赋值给other*/
func (self *Class) isAssignableFrom(other *Class) bool {
	//在三种情况下，S类型的引用值可以赋值给T类型：S 和T是同一类型；T是类且S是T的子类；或者T是接口且S实现了T接 口
	//t = s 能否成立
	s, t := other, self
	if s == t {
		return true
	}
	if !s.IsArray() {//s不是Array
		if !s.IsInterface() {//s不是接口
			// s是class 时
			if !t.IsInterface() {// t = s
				//s是class t是class时 s必须是t的子类
				return s.IsSubClassOf(t)
			} else {
				//s是class t是interface时，s必须实现了t接口
				return s.IsImplements(t)
			}
		} else {
			// s 是接口时
			if !t.IsInterface() {
				//s是接口 t是类时,t必须是object类
				return t.isJlObject()
			} else {
				//s是接口 t是接口时，t必须是s的超接口
				return t.isSuperInterfaceOf(s)
			}
		}
	} else {//s是array时
		if !t.IsArray() {//t不是array
			if !t.IsInterface() {
				// s是array t是class时，t必须是Object
				return t.isJlObject()
			} else {
				// s是array t是接口 则t必须实现了Cloneable()，Serializable()接口
				return t.isJlCloneable() || t.isJioSerializable()
			}
		} else {
			// s是array t也是array
			sc := s.ComponentClass()
			tc := t.ComponentClass()
			//s的数组元素必须和t的数组元素一样，或者t的数组元素可以赋值给s的数组元素
			return sc == tc || tc.isAssignableFrom(sc)
		}
	}

	return false
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

// iface extends self

func (self *Class) isSuperInterfaceOf(iface *Class) bool {

	return iface.isSubInterfaceOf(self)

}