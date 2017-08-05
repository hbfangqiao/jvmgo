package classpath

import "errors"
import "strings"

//Entry切片
type CompositeEntry []Entry
func newCompositeEntry(pathList string) CompositeEntry {
	compositeEntry := []Entry{}
	for _,path := range strings.Split(pathList,pathListSeparator){
		//将每一个path取出来，创建对应的entry,并添加到compositeEntry
		entry := newEntry(path)
		compositeEntry = append(compositeEntry,entry)
	}
	return compositeEntry
}

func (self CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	for _, entry := range self {
		//使用所有的entr去读取文件，读取到了就返回
		data, from, err := entry.readClass(className)
		if err == nil {
			return data,from,err
		}
	}
	return nil,nil,errors.New("class not found: "+className)
}

func (self CompositeEntry) String() string {
	strs := make([]string,len(self))
	for i, entry := range self {
		strs[i] = entry.String()
	}
	//将多个str用分隔符拼接起来
	return strings.Join(strs,pathListSeparator)
}
