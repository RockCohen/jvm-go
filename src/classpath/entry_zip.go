package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

// ZipEntry
/**
  ZipEntry : ZIP或JAR文件形式的类路径
  absPath: 存放ZIP或JAR文件的绝对路径
*/
type ZipEntry struct {
	absPath string
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath}
}
func (self *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	/**
		这里的r的类型为:ReadCloser,其结构说明如下：
	    type ReadCloser struct {
			f *os.File
			Reader
		}
	*/
	r, err := zip.OpenReader(self.absPath)
	if err != nil {
		return nil, nil, err
	}
	defer r.Close() //确保打开的文件关闭
	/**
	关于Reader的结构体如下：
	type Reader struct {
		r             io.ReaderAt
		File          []*File
		Comment       string
		decompressors map[uint16]Decompressor
	}
	下面的f代表一个[]*File
	*/
	for _, f := range r.File {
		if f.Name == className {
			rc, err := f.Open()
			if err != nil {
				return nil, nil, err
			}
			defer rc.Close()
			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}
			return data, self, err
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}
func (self *ZipEntry) String() string {
	return self.absPath
}
