package main

import (
	"classfile"
	"classpath"
	"flag"
	"fmt"
	"os"
	"rtda"
	"strings"
)

// Cmd
/**
该结构体来表示命令行参数与选项
成员说明：
1. helpFlag 帮助选项
2. versionFlag 版本选项
3. cpOption 类路径选项
4. XjreOption Java虚拟机将使用JDK的启动类路径来寻找和加载Java标准库中的类.该参数指定加载的jre的目录。
5. class 指定类路径
6. args 参数
*/
type Cmd struct {
	helpFlag    bool
	versionFlag bool
	cpOption    string
	XjreOption  string
	class       string
	args        []string
}

/**
使用flag包来帮助解析命令行的参数与选项
基本的流程是：
1. 对指定的选项进行设置名称，默认值，提示等
2. 调用flag.Parse()进行解析，解析失败时调用printUsage()方法进行提示
3. 通过flag.Args()方法来获取命令行参数，对命令行参数进行解析。
   - para-1:类名
   - para-2++:其他参数
4. 返回cmd结构体

关于flag的使用，参考官方文档：https://golang.org/pkg/flag/
*/
func parseCmd() *Cmd {
	//Go语言的推导声明，编译器会自动根据右值类型推断出左值的对应类型。
	cmd := &Cmd{}
	flag.Usage = printUsage
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	flag.StringVar(&cmd.XjreOption, "Xjre", "", "path to jre")
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}
	return cmd
}

/**
该函数打印提示信息
*/
func printUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}

/**
目前版本的JVM啥也没实现，直接打印输出信息表示JVM启动
*/
func startJVM(cmd *Cmd) {
	//cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	//fmt.Printf("classpath:%v class:%v args:%v\n",
	//	cp, cmd.class, cmd.args)
	//className := strings.Replace(cmd.class, ".", "/", -1)
	//cf := loadClass(className, cp)
	//fmt.Printf(cmd.class)
	//printClassInfo(cf)
	//frame := rtda.NewFrame(100, 100)
	//testLocalVars(frame.LocalVars())
	//testOperandStack(frame.OperandStack())
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	className := strings.Replace(cmd.class, ".", "/", -1)
	cf := loadClass(className, cp)
	mainMethod := getMainMethod(cf)
	if mainMethod != nil {
		interpret(mainMethod)
	} else {
		fmt.Printf("Main method not found in class %s\n", cmd.class)
	}

}

func getMainMethod(cf *classfile.ClassFile) *classfile.MemberInfo {
	for _, m := range cf.Methods() {
		if m.Name() == "main" && m.Descriptor() == "([Ljava/lang/String;)V" {
			return m
		}
	}
	return nil
}
func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile {
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		panic(err)
	}
	cf, err := classfile.Parse(classData)
	if err != nil {
		panic(err)
	}
	return cf
}

func printClassInfo(cf *classfile.ClassFile) {
	fmt.Printf(" version:  %v.%v\n", cf.MajorVersion(), cf.MinorVersion())
	fmt.Printf("constants count: %v\n", len(cf.ConstantPool()))
	fmt.Printf("access flags: 0x%x\n", cf.AccessFlags())
	fmt.Printf("this class: %v\n", cf.ClassName())
	fmt.Printf("super class: %v\n", cf.SuperClassName())
	fmt.Printf("interfaces: %v\n", cf.InterfaceNames())
	fmt.Printf("fields count: %v\n", len(cf.Fields()))
	for _, f := range cf.Fields() {
		fmt.Printf(" %s\n", f.Name())
	}
	fmt.Printf("methods count: %v\n", len(cf.Methods()))
	for _, m := range cf.Methods() {
		fmt.Printf(" %s\n", m.Name())
	}
}
func testLocalVars(vars rtda.LocalVars) {
	vars.SetInt(0, 100)
	vars.SetInt(1, -100)
	vars.SetLong(2, 2997924580)
	vars.SetLong(4, -2997924580)
	vars.SetFloat(6, 3.1415926)
	vars.SetDouble(7, 2.71828182845)
	vars.SetRef(9, nil)
	println(vars.GetInt(0))
	println(vars.GetInt(1))
	println(vars.GetLong(2))
	println(vars.GetLong(4))
	println(vars.GetFloat(6))
	println(vars.GetDouble(7))
	println(vars.GetRef(9))
}
func testOperandStack(ops *rtda.OperandStack) {
	ops.PushInt(100)
	ops.PushInt(-100)
	ops.PushLong(2997924580)
	ops.PushLong(-2997924580)
	ops.PushFloat(3.1415926)
	ops.PushDouble(2.71828182845)
	ops.PushRef(nil)
	println(ops.PopRef())
	println(ops.PopDouble())
	println(ops.PopFloat())
	println(ops.PopLong())
	println(ops.PopLong())
	println(ops.PopInt())
	println(ops.PopInt())
}
