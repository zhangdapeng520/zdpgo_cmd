module github.com/zhangdapeng520/zdpgo_cmd

go 1.17

require (
	github.com/zhangdapeng520/zdpgo_env v0.1.1
	github.com/zhangdapeng520/zdpgo_file v1.0.1
	github.com/zhangdapeng520/zdpgo_log v1.3.4
	github.com/zhangdapeng520/zdpgo_random v1.1.2
	github.com/zhangdapeng520/zdpgo_shell v0.1.1
)

require golang.org/x/text v0.3.7 // indirect

replace (
	github.com/zhangdapeng520/zdpgo_env v0.1.1 => ../zdpgo_env
	github.com/zhangdapeng520/zdpgo_file v1.0.1 => ../zdpgo_file
	github.com/zhangdapeng520/zdpgo_shell v0.1.1 => ../zdpgo_shell
)
