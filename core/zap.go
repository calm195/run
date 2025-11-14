package core

import (
	"fmt"
	"os"
	"run/global"
	"run/util"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Zap 获取 zap.Logger
func Zap() (log *zap.Logger) {
	if ok, _ := util.PathExists(global.Config.Zap.Director); !ok {
		_ = os.Mkdir(global.Config.Zap.Director, os.ModePerm)
		fmt.Printf("create %v directory\n", global.Config.Zap.Director)
	}
	levels := global.Config.Zap.Levels()
	length := len(levels)
	cores := make([]zapcore.Core, 0, length)
	for i := 0; i < length; i++ {
		core := NewZapCore(levels[i])
		cores = append(cores, core)
	}
	logger := zap.New(zapcore.NewTee(cores...))
	if global.Config.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	zap.ReplaceGlobals(logger)
	return logger
}
