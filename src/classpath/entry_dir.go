package classpath

import (
	"io/ioutil"
	"path/filepath"
)

//DirEntry
/**
目录形式的类路径结构
成员变量说明：
absDir:存放类的绝对路径
*/
type DirEntry struct {
	absDir string
}

/**
通过该方法得到DirEntry实体。
Golang中没有设计构造函数. 取而代之的, 设计Golang的大师希望你用普通函数去实现构造的任务.
*/
func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path) //通过filepath.Abs()方法获取绝对路径
	if err != nil {
		panic(err) //When you panic in Go, you’re freaking out, it’s not someone else problem, it’s game over man.
	}
	return &DirEntry{absDir}
}

func (self *DirEntry) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(self.absDir, className)
	data, err := ioutil.ReadFile(fileName)
	return data, self, err
}
func (self *DirEntry) String() string {
	return self.absDir
}
