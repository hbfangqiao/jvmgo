package classpath

import "os"
import "path/filepath"

type Classpath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClasspath Entry
}
/*
	将启动类路径，扩展类路径，用户类路径解析成Entry
 */
func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp;
}
func (self *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	if data, entry, err := self.bootClasspath.readClass(className); err == nil {
		return data,entry,err
	}
	if data, entry, err := self.extClasspath.readClass(className); err == nil {
		return data,entry,err
	}
	return self.userClasspath.readClass(className)
}
func (self *Classpath) String() string {
	return self.userClasspath.String()
}

func (self *Classpath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)
	// jre/lib/*
	jreLibPath := filepath.Join(jreDir,"lib","*")
	self.bootClasspath = newWildcardEntry(jreLibPath)
	// jre/lib/ext/*
	jreExtPath := filepath.Join(jreDir,"lib","ext","*")
	self.extClasspath = newWildcardEntry(jreExtPath)
}

func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}
	if exists("./jre") {
		return "./jre"
	}
	if jh := os.Getenv("JAVA_HOME"); jh!= "" {
		return filepath.Join(jh,"jre")
	}
	panic("Can not find jre folder")
}

//判断目录是否存在
func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func (self *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	self.userClasspath = newEntry(cpOption)
}

