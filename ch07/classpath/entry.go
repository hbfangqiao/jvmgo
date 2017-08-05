package classpath

import "os"
import "strings"

/*
	分隔符windows下是;Linux,MacOS下是：
 */
const pathListSeparator  = string(os.PathListSeparator)

type Entry interface {
	readClass(className string) ([]byte,Entry,error)
	String() string
}

func newEntry(path string) Entry {
	if strings.Contains(path,pathListSeparator) {
		//path中包含分割符windows下是; Mac和Linux下是：
		return newCompositeEntry(path)
	}
	if strings.HasSuffix(path,"*") {
		//path中包含通配符
		return newWildcardEntry(path)
	}
	if strings.HasSuffix(path,".jar") || strings.HasSuffix(path,".JAR") ||
	strings.HasSuffix(path,".zip") || strings.HasSuffix(path,".ZIP") {
		return newZipEntry(path)
	}
	return newDirEntry(path)
}
