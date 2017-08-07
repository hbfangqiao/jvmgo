package heap

import (
	"jvmgo/ch10/classfile"
	"strings"
)
/*类信息 放入方法区*/
type Class struct {
	accessFlags       uint16        //类的访问标识
	name              string        //类名 完全限定名 例:java/lang/Object
	superClassName    string        //超类名 完全限定名
	interfaceNames    []string      //接口名 完全限定名
	constantPool      *ConstantPool //运行时常量池指针
	fields            []*Field      //字段表
	methods           []*Method     //方法表
	loader            *ClassLoader  //类加载器指针
	superClass        *Class        //超类指针
	interfaces        []*Class      //接口指针
	instanceSlotCount uint          //实例变量占据空间大小
	staticSlotCount   uint          //类变量占据的空间大小
	staticVars        Slots         //存放静态变量
	initStarted       bool          //判断类的<clinit>方法是否已经开始执行
	jClass            *Object       // java.lang.Class实例
	sourceFile        string
}

func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = cf.AccessFlags()
	class.name = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfaceNames()
	class.constantPool = newConstantPool(class, cf.ConstantPool())
	class.fields = newFields(class, cf.Fields())
	class.methods = newMethods(class, cf.Methods())
	class.sourceFile = getSourceFile(cf)
	return class
}
/*－－－－－－－＃判断访问标识符是否被设置－－－－－－－－－－*/
func (self *Class) IsPublic() bool {
	return 0 != self.accessFlags & ACC_PUBLIC
}

func (self *Class) IsFinal() bool {
	return 0 != self.accessFlags & ACC_FINAL
}

func (self *Class) IsSuper() bool {
	return 0 != self.accessFlags & ACC_SUPER
}

func (self *Class) IsInterface() bool {
	return 0 != self.accessFlags & ACC_INTERFACE
}

func (self *Class) IsAbstract() bool {
	return 0 != self.accessFlags & ACC_ABSTRACT
}

func (self *Class) IsSynthetic() bool {
	return 0 != self.accessFlags & ACC_SYNTHETIC
}

func (self *Class) IsAnnotation() bool {
	return 0 != self.accessFlags & ACC_ANNOTATION
}

func (self *Class) IsEnum() bool {
	return 0 != self.accessFlags & ACC_ENUM
}
/*是否是基本类型的类*/
func (self *Class) IsPrimitive() bool {
	_, ok := primitiveTypes[self.name]
	return ok
}

/*－－－－－－判断访问标识符是否被设置＃－－－－－－－－－－*/

//getters
func (self *Class) Name() string {
	return self.name
}

func (self *Class) ConstantPool() *ConstantPool {
	return self.constantPool
}

func (self *Class) Fields() []*Field {
	return self.fields
}

func (self *Class) Methods() []*Method {
	return self.methods
}

func (self *Class) SourceFile() string {
	return self.sourceFile
}

func (self *Class) StaticVars() Slots {
	return self.staticVars
}

func (self *Class) SuperClass() *Class {
	return self.superClass
}

func (self *Class) Loader() *ClassLoader {
	return self.loader
}

func (self *Class) InitStarted() bool {
	return self.initStarted
}

func (self *Class) JClass() *Object {
	return self.jClass
}

func (self *Class) StartInit() {
	self.initStarted = true
}

func (self *Class) isAccessibleTo(other *Class) bool {
	return self.IsPublic() || self.GetPackageName() == other.GetPackageName()
}
/*从类的全限定名中截取包名*/
func (self *Class) GetPackageName() string {
	//当类定义在默认包中的话，它的包名是空字符串
	if i := strings.LastIndex(self.name, "/"); i >= 0 {
		return self.name[:i]
	}
	return ""
}

func (self *Class) JavaName() string {
	return strings.Replace(self.name, "/", ".", -1)
}

func (self *Class) GetMainMethod() *Method {
	return self.getStaticMethod("main", "([Ljava/lang/String;)V")
}

func (self *Class) GetClinitMethod() *Method {
	return self.getStaticMethod("<clinit>", "()V")
}

func getSourceFile(cf *classfile.ClassFile) string {
	if sfAttr := cf.SourceFileAttribute(); sfAttr != nil {
		return sfAttr.FileName()
	}
	return "Unknown"
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

/*根据方法名和描述符查找字段*/
func (self *Class) getMethod(name, descriptor string, isStatic bool) *Method {
	for c := self; c != nil; c = c.superClass {
		for _, method := range c.methods {
			if method.IsStatic() == isStatic &&
				method.name == name &&
				method.descriptor == descriptor {
				return method
			}
		}
	}
	return nil
}

/*根据字段名和描述符查找字段*/
func (self *Class) getField(name, descriptor string, isStatic bool) *Field {
	for c := self; c != nil; c = c.superClass {
		for _, field := range c.fields {
			if field.IsStatic() == isStatic &&//是否都是静态的
				field.name == name &&//是否同名
				field.descriptor == descriptor {
				//描述符是否相同
				return field
			}
		}
	}
	return nil
}

func (self *Class) isJlObject() bool {
	return self.name == "java/lang/Object"
}

func (self *Class) isJlCloneable() bool {
	return self.name == "java/lang/Cloneable"
}

func (self *Class) isJioSerializable() bool {
	return self.name == "java/io/Serializable"
}

func (self *Class) NewObject() *Object {
	return newObject(self)
}

func (self *Class) ArrayClass() *Class {
	arrayClassName := getArrayClassName(self.name)
	return self.loader.LoadClass(arrayClassName)
}

func (self *Class) GetInstanceMethod(name, descriptor string) *Method {
	return self.getMethod(name, descriptor, false)
}

func (self *Class) GetRefVar(fieldName, fieldDescriptor string) *Object {
	field := self.getField(fieldName, fieldDescriptor, true)
	return self.staticVars.GetRef(field.slotId)
}
func (self *Class) SetRefVar(fieldName, fieldDescriptor string, ref *Object) {
	field := self.getField(fieldName, fieldDescriptor, true)
	self.staticVars.SetRef(field.slotId, ref)
}
