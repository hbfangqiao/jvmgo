package main

import "fmt"
import "strings"
import "jvmgo/ch07/classpath"
import "jvmgo/ch07/rtda/heap"

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
	classLoader := heap.NewClassLoader(cp, cmd.verboseClassFlag)

	className := strings.Replace(cmd.class, ".", "/", -1)
	mainClass := classLoader.LoadClass(className)
	mainMethod := mainClass.GetMainMethod()
	if mainMethod != nil {
		interpret(mainMethod, cmd.verboseInstFlag)
	} else {
		fmt.Printf("Main method not found in class %s\n", cmd.class)
	}
}





