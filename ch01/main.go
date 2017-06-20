package main

import "fmt"

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
	fmt.Printf("classpath:%s class:%s args:%v \n",
	cmd.cpOption,cmd.class,cmd.args)
}
