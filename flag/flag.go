package flag

import (
	"flag"
)

var (
	Version      = flag.Bool("v", false, "版本号")
	PathConfFile = flag.String("c", "./config.toml", "指定配置文件")
)

func InitFlag() {
	// 解析参数
	flag.Parse() // 解析命令行参数
}
