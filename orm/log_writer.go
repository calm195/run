package orm

import (
	"fmt"
	"run/config"
	"run/global"

	"gorm.io/gorm/logger"
)

type Writer struct {
	writer logger.Writer
	config config.PgsqlConfig
}

func NewWriter(config config.PgsqlConfig) *Writer {
	return &Writer{config: config}
}

// Printf 格式化打印日志
func (c *Writer) Printf(message string, data ...any) {
	if c.config.LogZap {
		logMessage := fmt.Sprintf(message, data...)
		switch c.config.LogLevel() {
		case logger.Silent:
			global.Log.Debug(logMessage)
		case logger.Error:
			global.Log.Error(logMessage)
		case logger.Warn:
			global.Log.Warn(logMessage)
		case logger.Info:
			global.Log.Info(logMessage)
		default:
			global.Log.Info(logMessage)
		}
	}
}
