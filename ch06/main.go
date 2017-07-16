package main

import "fmt"
import "strings"
import "jvmgo/ch06/classfile"
import "jvmgo/ch06/classpath"

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {//如果输入了-version
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" { //如果输了了-help
		//printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOperation, cmd.cpOption)
	className := strings.Replace(cmd.class,".","/",-1)
	//读取并解析class文件
	cf := loadClass(className,cp)
	//查找类的main方法
	mainMethod := getMainMethod(cf)
	if mainMethod != nil {
		//执行main方法
		interpret(mainMethod)
	} else {
		fmt.Printf("Main method not found in class %s\n",cmd.class)
	}

}
func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile {
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		panic(err)
	}
	cf, err := classfile.Parse(classData)
	if err != nil {
		panic(err)
	}
	return cf
}

func getMainMethod(cf *classfile.ClassFile) *classfile.MemberInfo {
	for _, m := range cf.Methods() {
		fmt.Println(m.Name(),m.Descriptor())
		if m.Name() == "main" && m.Descriptor() == "([Ljava/lang/String;)V" {
			return m
		}
	}
	return nil
}


