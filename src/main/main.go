package main

import (
	"fmt"
	"path"
)

/**
执行方法入口
*/
func main() {
	fmt.Println(path.Join("c:", "aa", "bb", "cc.txt"))
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}
