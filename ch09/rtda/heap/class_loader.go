package heap

import "fmt"
import "jvmgo/ch09/classfile"
import "jvmgo/ch09/classpath"

type ClassLoader struct {
	cp          *classpath.Classpath //Classpath指针
	verboseFlag bool
	classMap    map[string]*Class    //已经加载的类，key是类的完全限定名。classMap当作方法区的具体实现
}

func NewClassLoader(cp *classpath.Classpath, verboseFlag bool) *ClassLoader {
	loader := &ClassLoader{
		cp:          cp,
		verboseFlag: verboseFlag,
		classMap:    make(map[string]*Class),
	}

	loader.loadBasicClasses()
	loader.loadPrimitiveClasses()
	return loader
}

func (self *ClassLoader) loadBasicClasses() {
	//先加载java.lang.Class类，会触发 java.lang.Object等类和接口的加载。
	jlClassClass := self.LoadClass("java/lang/Class")
	//遍历classMap，给已经加载的每一个类关联类对象。
	for _, class := range self.classMap {
		if class.jClass == nil {
			class.jClass = jlClassClass.NewObject()
			class.jClass.extra = class
		}
	}
}

func (self *ClassLoader) loadPrimitiveClasses() {
	for primitiveType, _ := range primitiveTypes {
		self.loadPrimitiveClass(primitiveType) // primitiveType是 void、 int、 float等
	}
}

func (self *ClassLoader) loadPrimitiveClass(className string) {
	//void和基本类型类,也是在运行时生成的
	class := &Class{
		accessFlags: ACC_PUBLIC, // todo
		name:        className,
		loader:      self,
		initStarted: true,
	}
	class.jClass = self.classMap["java/lang/Class"].NewObject()
	class.jClass.extra = class
	self.classMap[className] = class
}

func (self *ClassLoader) LoadClass(name string) *Class {
	if class, ok := self.classMap[name]; ok {
		//查看类是否已经被加载
		return class
	}

	var class *Class
	if name[0] == '[' { //如果是数组类，调用loadArrayClass()加载类
		class = self.loadArrayClass(name)
	} else {//如果不是数组类，调用loadNonArrayClass()加载类
		class = self.loadNonArrayClass(name)
	}
	//看java.lang.Class是否已经加载,如果是则给类关联类对象。
	if jlClassClass, ok := self.classMap["java/lang/Class"]; ok {
		class.jClass = jlClassClass.NewObject()
		class.jClass.extra = class
	}

	return class

}

func (self *ClassLoader) loadArrayClass(name string) *Class {
	//需要生成一个Class结构体，即虚拟机运行时生成的类
	class := &Class{
		accessFlags: ACC_PUBLIC,//
		name:        name,
		loader:      self,
		initStarted: true,
		superClass:  self.LoadClass("java/lang/Object"),
		interfaces: []*Class{
			self.LoadClass("java/lang/Cloneable"),
			self.LoadClass("java/io/Serializable"),
		},
	}
	self.classMap[name] = class
	return class
}

func (self *ClassLoader) loadNonArrayClass(name string) *Class {
	//找到class文件并把数据读取到内存
	data, entry := self.readClass(name)
	//解析class文件，生成虚拟机可以使用的类数据,并放入方法区
	class := self.defineClass(data)
	//进行链接
	link(class)
	if self.verboseFlag {
		fmt.Printf("[Loaded %s from %s]\n", name, entry)
	}
	return class
}
/*调用Classpath的ReadClass()方法*/
func (self *ClassLoader) readClass(name string) ([]byte, classpath.Entry) {
	data, entry, err := self.cp.ReadClass(name)
	if err != nil {
		panic("java.lang.ClassNotFoundException: " + name)
	}
	//返回data的同时返回entry是为了打印类加载信息
	return data, entry
}

func (self *ClassLoader) defineClass(data []byte) *Class {
	//将class文件数据转换成Class结构体
	class := parseClass(data)
	class.loader = self
	//superClass存放超类名
	resolveSuperClass(class)
	//interfaces存放直接接口表
	resolveInterfaces(class)
	self.classMap[class.name] = class
	return class
}

func parseClass(data []byte) *Class {
	cf, err := classfile.Parse(data)
	if err != nil {
		panic("java.lang.ClassFormatError")
	}
	return newClass(cf)
}

func resolveSuperClass(class *Class) {
	if class.name != "java/lang/Object" {
		//递归调用LoadClass()加载超类
		class.superClass = class.loader.LoadClass(class.superClassName)
	}
}

func resolveInterfaces(class *Class) {
	interfaceCount := len(class.interfaceNames)
	if interfaceCount > 0 {
		//加载类的每一个接口
		class.interfaces = make([]*Class, interfaceCount)
		for i, interfaceName := range class.interfaceNames {
			class.interfaces[i] = class.loader.LoadClass(interfaceName)
		}
	}
}
/*类的链接*/
func link(class *Class) {
	verify(class)//验证阶段
	prepare(class)//准备阶段，给类分配空间并给予初始值
}

func verify(class *Class) {
	//todo
}

func prepare(class *Class) {
	calcInstanceFieldSlotIds(class)//计算实例字段的个数，并给它们编号
	calcStaticFieldSlotIds(class)//计算静态字段的个数，并给它们编号
	allocAndInitStaticVars(class)
}
/*计算实例字段的个数*/
func calcInstanceFieldSlotIds(class *Class) {
	slotId := uint(0)
	if class.superClass != nil {
		//发生在加载类的准备阶段，会递归加载超类。故这里的superclass的instanceSlotCount，已经计算出来
		slotId = class.superClass.instanceSlotCount
	}
	for _, field := range class.fields {
		if !field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.instanceSlotCount = slotId
}

func calcStaticFieldSlotIds(class *Class) {
	slotId := uint(0)
	for _, field := range class.fields {
		if field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.staticSlotCount = slotId
}
/*给类变量分配空间，然后给它们初始值*/
func allocAndInitStaticVars(class *Class) {
	class.staticVars = newSlots(class.staticSlotCount)
	for _, field := range class.fields {
		//如果静态变量属于基本类型或String类型，有final修饰符， 且它的值在编译期已知，则该值存储在class文件常量池中。
		if field.IsStatic() && field.IsFinal() {
			//从常量池中加载常量值，并给静态变量赋值
			initStaticFinalVar(class, field)
		}
	}
}
/**从常量池中加载常量值，并给静态变量赋值*/
func initStaticFinalVar(class *Class, field *Field) {
	vars := class.staticVars
	cp := class.constantPool
	cpIndex := field.ConstValueIndex()
	slotId := field.SlotId()
	if cpIndex > 0 {
		switch field.Descriptor() {
		case "Z", "B", "C", "S", "I"://Z boolean B byte C char S short I int
			val := cp.GetConstant(cpIndex).(int32)
			vars.SetInt(slotId, val)
		case "J"://J long
			val := cp.GetConstant(cpIndex).(int64)
			vars.SetLong(slotId, val)
		case "F"://F float
			val := cp.GetConstant(cpIndex).(float32)
			vars.SetFloat(slotId, val)
		case "D":
			val := cp.GetConstant(cpIndex).(float64)
			vars.SetDouble(slotId, val)
		case "Ljava/lang/String;":
			goStr := cp.GetConstant(cpIndex).(string)
			jStr := JString(class.Loader(), goStr)
			vars.SetRef(slotId, jStr)
		}
	}
}