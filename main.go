package main

import (
	"run/core"
	"run/global"
	"run/orm"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func main() {
	global.Vp = core.Viper()
	global.Log = core.Zap()
	global.Db = orm.Gorm()
}
