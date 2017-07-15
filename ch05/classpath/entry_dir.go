package classpath

import "io/ioutil"
import "path/filepath"

type DirEntry struct {
	absDir string
}

func newDirEntry(path string) *DirEntry {
	//将path转换成绝对路径
	absDir,err := filepath.Abs(path)
	if err != nil {
		//终止程序运行
		panic(err)
	}
	return &DirEntry{absDir}
}
/*
	将类名凭借到路径上，读取文件
 */
func (self *DirEntry) readClass(className string) ([]byte,Entry,error){
	fileName := filepath.Join(self.absDir,className)
	data, err := ioutil.ReadFile(fileName)
	return data,self,err
}
func (self *DirEntry) String() string {
	return self.absDir
}
