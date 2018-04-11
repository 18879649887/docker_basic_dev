package build

import (
	"fmt"
	"../flag"
	"syscall"
	"../net"
)

func Build(flagstring *flag.SFlag) {
	//利用解析过来的参数=>构建namespace,查看cgroups分配资源的规则,构建虚拟容器内网络的互通
}