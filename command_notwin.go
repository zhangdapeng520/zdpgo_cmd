//go:build !windows
// +build !windows

package zdpgo_cmd

var preExecHookFn func(*Command)
