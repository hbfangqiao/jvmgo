package heap

import "fmt"
import "jvmgo/ch06/classfile"
import "jvmgo/ch06/classpath"

type ClassLoader struct {
	cp       *classpath.Classpath //Classpath指针
	classMap map[string]*Class    //已经加载的类，key是类的完全限定名。classMap当作方法区的具体实现
}

func NewClassLoader(cp *classpath.Classpath) *ClassLoader {
	return &ClassLoader{
		cp:        cp,
		classMap: make(map[string]*Class),
	}
}

func (self *ClassLoader) LoadClass(name string) *Class {
	if class, ok := self.classMap[name]; ok {//查看类是否已经被加载
		return class
	}
	return self.loadNonArrayClass(name)
}

func (self *ClassLoader) loadNonArrayClass(name string) *Class {
	//找到class文件并把数据读取到内存
	data, entry := self.readClass(name)
	//解析class文件，生成虚拟机可以使用的类数据,并放入方法区
	class := self.defineClass(data)
	//进行链接
	link(class)
	fmt.Printf("[Loaded %s from %s]\n",name, entry)
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
	if class.name != "java/lang/Object" {//递归调用LoadClass()加载超类
		class.superClass = class.loader.LoadClass(class.superClassName)
	}
}

func resolveInterfaces(class *Class) {
	interfaceCount := len(class.interfaceNames)
	if interfaceCount >0 {//加载类的每一个接口
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