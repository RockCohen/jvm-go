package classpath

import (
	"os"
	"path/filepath"
)

// Classpath
/**
三个 Entry 变量分别代表：
1. bootClasspath : 启动类路径
2. extClasspath  : 扩展类路径
3. userClasspath : 用户类路径
*/
type Classpath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClasspath Entry
}

// Parse /**
// jreOption : 启动类路径选项
// cpOption  : 用户类路径选项
func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}
func (self *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class" // className格式：java.lang.Object
	//启动类路径读取
	if data, entry, err := self.bootClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	//扩展类路径读取
	if data, entry, err := self.extClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	// 用户类路径读取
	return self.userClasspath.readClass(className)
}

func (self *Classpath) String() string {
	return self.userClasspath.String()
}

/**
启动类路径与扩展类路径
*/
func (self *Classpath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)
	// jre/lib/*
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	self.bootClasspath = newWildcardEntry(jreLibPath)
	// jre/lib/ext/*
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	self.extClasspath = newWildcardEntry(jreExtPath)
}

/**
优先使用用户输入的-Xjre选项作为jre目录。如果没有输入该
选项，则在当前目录下寻找jre目录。如果找不到，尝试使用
JAVA_HOME环境变量
*/
func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}
	if exists("./jre") {
		return "./jre"
	}
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}
	panic("Can not find jre folder!")
}

/**
判断目录存在与否
*/
func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

/**
用户类路径
*/
func (self *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	self.userClasspath = newEntry(cpOption)
}
