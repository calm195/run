package main

import (
	"fmt"
	"run/core"
	"run/global"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func main() {
	fmt.Println("begin")
	global.Vp = core.Viper()
	fmt.Println("end")
}
