package classpath

import "archive/zip"
import "errors"
import "io/ioutil"
import "path/filepath"

type ZipEntry struct {
	absPath string
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if (err != nil) {
		panic(err)
	}
	return &ZipEntry{absPath}
}

func (self *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	//通过absPath 打开zip文件
	r, err := zip.OpenReader(self.absPath)
	if err != nil {
		return nil, nil, err
	}
	defer r.Close()
	//遍历r中的文件
	for _, f := range r.File {
		if f.Name == className {
			//发现与className相同的文件
			rc, err := f.Open()
			if err != nil {
				return nil,nil,err
			}
			defer rc.Close()
			//读取改文件的内容
			data, err := ioutil.ReadAll(rc)
			if err !=nil {
				return nil,nil,err
			}
			return data,self,nil

		}
	}
	//没有发现zip中有同类名的文件
	return nil,nil, errors.New("class not found:"+className)
}

func (self *ZipEntry) String() string {
	return self.absPath
}
