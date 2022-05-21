package g

import (
	"github.com/zhangdapeng520/zdpgo_file"
	"github.com/zhangdapeng520/zdpgo_random"
	"zdptool/config"
)

/*
@Time : 2022/5/21 11:16
@Author : 张大鹏
@File : global.go
@Software: Goland2021.3.1
@Description:
*/

var (
	Config *config.Config
	Random *zdpgo_random.Random
	File   *zdpgo_file.File
)
