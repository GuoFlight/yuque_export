package conf

import (
	"github.com/BurntSushi/toml"
	"log"
)

var (
	GConf ConfigFileStruct
)

type ConfigFileStruct struct {
	Cookie     string `toml:"cookie"`
	XCsrfToken string `toml:"x_csrf_token"`
	FileDir    string `toml:"file_dir"`
}

// ParseConfig 解析配置文件
func ParseConfig(pathConfFile string) {
	if _, err := toml.DecodeFile(pathConfFile, &GConf); err != nil {
		log.Fatal(err)
	}
}
