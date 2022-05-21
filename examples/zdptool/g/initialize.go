package g

import (
	"github.com/zhangdapeng520/zdpgo_file"
	"github.com/zhangdapeng520/zdpgo_random"
	"zdptool/config"
)

/*
@Time : 2022/5/21 11:17
@Author : 张大鹏
@File : initialize.go
@Software: Goland2021.3.1
@Description:
*/

func init() {
	InitConfig()
	InitRandom()
	InitFile()
}

func InitConfig() {
	Config = &config.Config{
		Debug:       false,
		LogFilePath: "c:/tmp",
	}
}

func InitRandom() {
	Random = zdpgo_random.NewWithConfig(zdpgo_random.Config{
		Debug:       Config.Debug,
		LogFilePath: Config.LogFilePath,
	})
}

func InitFile() {
	File = zdpgo_file.NewWithConfig(zdpgo_file.Config{
		Debug:       Config.Debug,
		LogFilePath: Config.LogFilePath,
	})
}
