package main

import (
	"flag"
	"fmt"
	"os"
)

type Cmd struct {
	helpFlag    bool
	versionFlag bool
	cpOption    string
	XjreOption  string
	class       string
	args        []string
}

func parseCmd() *Cmd {
	/**
	 *结构体初始化 https://www.cnblogs.com/liyutian/p/10050320.html
	 * 1、指针类型
	 *    var 变量 = new (结构体名称)
	 *    var 变量 = &结构体名称{}
	 *    var 变量 = &结构体名称{成员A:值，成员B:值}
	 * 2、值类型
	 *    var 变量= 结构体{成员A:值，成员B:值}
	 */
	cmd := &Cmd{}

	/**
	 * https://www.jianshu.com/p/f9cf46a4de0e
	 *&cmd.helpFlag 用来接收命令行中输入的 -u 后面的参数值
	 *"help" 就是命令中 helpFlag指定参数
	 *false 是默认值
	 *"print help message" 默认
	 */
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

func printUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}
