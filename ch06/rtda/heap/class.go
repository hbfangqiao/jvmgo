package heap

import (
	"jvmgo/ch06/classfile"
	"strings"
)
/*类信息 放入方法区*/
type Class struct {
	accessFlags       uint16//类的访问标识
	name              string//类名 完全限定名 例:java/lang/Object
	superClassName    string//超类名 完全限定名
	interfaceNames    []string//接口名 完全限定名
	constantPool      *ConstantPool//运行时常量池指针
	fields            []*Field//字段表
	methods           []*Method//方法表
	loader            *ClassLoader//类加载器指针
	superClass        *Class//超类指针
	interfaces        []*Class//接口指针
	instanceSlotCount uint//实例变量占据空间大小
	staticSlotCount   uint//类变量占据的空间大小
	staticVars        Slots//存放静态变量
}

func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = cf.AccessFlags()
	class.name = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfaceNames()
	class.constantPool = newConstantPool(class,cf.ConstantPool())
	class.fields = newFields(class,cf.Fields())
	class.methods = newMethods(class, cf.Methods())
	return class
}
/*－－－－－－－＃判断访问标识符是否被设置－－－－－－－－－－*/
func (self *Class) IsPublic() bool {
	return 0 != self.accessFlags&ACC_PUBLIC
}

func (self *Class) IsFinal() bool {
	return 0 != self.accessFlags&ACC_FINAL
}

func (self *Class) IsSuper() bool {
	return 0 != self.accessFlags&ACC_SUPER
}

func (self *Class) IsInterface() bool {
	return 0 != self.accessFlags&ACC_INTERFACE
}

func (self *Class) IsAbstract() bool {
	return 0 != self.accessFlags&ACC_ABSTRACT
}

func (self *Class) IsSynthetic() bool {
	return 0 != self.accessFlags&ACC_SYNTHETIC
}

func (self *Class) IsAnnotation() bool {
	return 0 != self.accessFlags&ACC_ANNOTATION
}

func (self *Class) IsEnum() bool {
	return 0 != self.accessFlags&ACC_ENUM
}
/*－－－－－－判断访问标识符是否被设置＃－－－－－－－－－－*/

//getters
func (self *Class) ConstantPool() *ConstantPool {
	return self.constantPool
}

func (self *Class) StaticVars() Slots {
	return self.staticVars
}

func (self *Class) isAccessibleTo(other *Class) bool {
	return self.IsPublic() || self.getPackageName() == other.getPackageName()
}
/*从类的全限定名中截取包名*/
func (self *Class) getPackageName() string {
	//当类定义在默认包中的话，它的包名是空字符串
	if i := strings.LastIndex(self.name,"/"); i >= 0 {
		return self.name[:i]
	}
	return ""
}

func (self *Class) GetMainMethod() *Method {
	return self.getStaticMethod("main", "([Ljava/lang/String;)V")
}


func (self *Class) getStaticMethod(name, descriptor string) *Method {
	for _, method := range self.methods {
		if method.IsStatic() &&
			method.name == name &&
			method.descriptor == descriptor {
			return method
		}
	}
	return nil

}

func (self *Class) NewObject() *Object {
	return newObject(self)
}