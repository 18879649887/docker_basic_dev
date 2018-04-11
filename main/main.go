package main

import (
	"fmt"
	"os"
	"os/exec"
	"bytes"
	"../flag"
	"../build"
)

/**
 * wocker version -v    输出版本号
 * wocker build -n container_name -p prot:port  开放容器内外端口映射
 * 达成目标: 
 *	   linux底层的namespaces隔离和cgroups隔离;
 *     为隔离后的容器添加属于自己的ip
 *     利用宿主机上的telnet可以访问wocker的ip及端口号
 *     局域网内的其他机器可以访问wocker ip:port
 */

func main() {
	var cmd *exec.Cmd;
	if len(os.Args) == 1 {
		fmt.Println("please input command!\n");
	} else {
		cmd = exec.Command("/bin/sh", "-c", Array2String(os.Args[1:]));
		if out, err := cmd.Output(); err != nil {
			//err情况下,分为两种情况:wocker开头与其他linux识别不了的模式开头
			if os.Args[1] == "wocker" {
				flagstr := new(flag.SFlag);
				flagstr = flagstr.Resolve(os.Args[2:]);

				//业务层,根据指令进行判断.注意包的循环引用问题的发生,业务层不能既用来做全局动作,又用来分发任务
				switch(flagstr.Command) {
					case "version": flag.ShowVersion(); break;
					case "build": build.Build(flagstr);break;      
					default: fmt.Println("cmd developing"); break;
				}
			} else {
				os.Exit(1);
			}
		} else {
			//输出linux能识别的指令结果
			fmt.Println(out);
		}
	}
}

func Array2String(arr []string) string {
	var buffer bytes.Buffer;
	for _, cmdstring := range arr {
		buffer.WriteString(cmdstring + " ");
	} 
	return buffer.String();
}