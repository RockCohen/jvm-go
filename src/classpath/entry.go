package classpath

import (
	"os"
	"strings"
)

const pathListSeparator = string(os.PathListSeparator)

// Entry
/**
在加载主类之前，jvm首先会加载它的超类，如果没有明确的继承关系，那么其父类便是:java.lang.Object。
执行main方法之前还需要加载java.lang.String类以及java.lang.String[].
于是虚拟机面临这样一个问题：虚拟机如何寻找需要加载的类的。

Java虚拟机规范给了足够的自由：从何处加载class文件并没有作硬性的规定。可以来自：
1. 文件系统中的zip或者jar文件
2. 网络、数据库
3. 运行期间生成的class文件

Java虚拟机规范并没有规定虚拟机去哪里加载类，于是实现方式由虚拟机实现者自定。通常的实现是：根据类路径来搜索类。通常分为三种类：
1. 启动类路径 默认对应jre/lib
2. 扩展类路径 默认对应jre/lib/ext
3. 用户类路径 默认当前目录

类路径接口表示:
将类路径对应的类读取到内存中。
*/
type Entry interface {
	/**
	负责读取和加载class文件
	参数格式:
	java/lang/Object.class
	*/
	readClass(className string) ([]byte, Entry, error)
	// String
	/**
	toString方法
	*/
	String() string
}

/**
Entry构造函数
Entry有四种具体的方法来构造分别为：
1. CompositeEntry
2. WildcardEntry(CompositeEntry)
3. ZipEntry
4. DirEntry 目录形式的类路径
*/
func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}
	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}
	return newDirEntry(path)

}
