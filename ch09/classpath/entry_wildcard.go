package classpath

import "os"
import "path/filepath"
import "strings"

func newWildcardEntry(path string) CompositeEntry {
	//将path末尾的通配符 “*”截取
	baseDir := path[:len(path)-1]
	compositeEntry := []Entry{}
	walkFn := func(path string,info os.FileInfo,err error) error {
		//根据后缀名选出jar文件，
		if err != nil {
			return err
		}
		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}
		if strings.HasSuffix(path,".jar") || strings.HasSuffix(path,".JAR") {
			jarEntry := newZipEntry(path)
			compositeEntry = append(compositeEntry,jarEntry)
		}
		return nil
	}
	filepath.Walk(baseDir, walkFn)
	return compositeEntry
}
