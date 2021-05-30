package main

import (
	"flag"
	"fmt"
	"os"
)

// Cmd
/**
该结构体来表示命令行参数与选项
*/
type Cmd struct {
	helpFlag    bool
	versionFlag bool
	cpOption    string
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
	fmt.Printf("classpath:%s class:%s args:%v\n",
		cmd.cpOption, cmd.class, cmd.args)
}

/**
执行方法入口
*/
func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}
