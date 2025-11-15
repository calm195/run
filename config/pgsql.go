package config

import (
	"strings"

	"gorm.io/gorm/logger"
)

type PgsqlConfig struct {
	Path               string `mapstructure:"path" json:"path" yaml:"path"`                            // 数据库主机地址
	Port               string `mapstructure:"port" json:"port" yaml:"port"`                            // 数据库端口
	Username           string `mapstructure:"username" json:"username" yaml:"username"`                // 登录用户名
	Password           string `mapstructure:"password" json:"password" yaml:"password"`                // 登录密码
	Dbname             string `mapstructure:"db-name" json:"db-name" yaml:"db-name"`                   // 数据库名
	Config             string `mapstructure:"config" json:"config" yaml:"config"`                      // 连接参数
	MaxIdleConnections int    `mapstructure:"max-idle" json:"max-idle" yaml:"max-idle"`                // 空闲中的最大连接数
	MaxOpenConnections int    `mapstructure:"max-open" json:"max-open" yaml:"max-open"`                // 打开到数据库的最大连接数
	LogMode            string `mapstructure:"log-mode" json:"log-mode" yaml:"log-mode" default:"info"` // 是否开启Gorm全局日志
	LogZap             bool   `mapstructure:"log-zap" json:"log-zap" yaml:"log-zap"`                   // 是否通过zap写入日志文件
}

func (c *PgsqlConfig) LogLevel() logger.LogLevel {
	switch strings.ToLower(c.LogMode) {
	case "silent":
		return logger.Silent
	case "error":
		return logger.Error
	case "warn":
		return logger.Warn
	case "info":
		return logger.Info
	default:
		return logger.Info
	}
}

// Dsn 基于配置文件获取 dsn
func (c *PgsqlConfig) Dsn() string {
	return "host=" + c.Path + " user=" + c.Username + " password=" + c.Password + " dbname=" + c.Dbname + " port=" + c.Port + " " + c.Config
}

// DefaultDsn 基于配置文件获取 默认dsn
func (c *PgsqlConfig) DefaultDsn() string {
	defaultDbName := "postgres"
	return "host=" + c.Path + " user=" + c.Username + " password=" + c.Password + " dbname=" + defaultDbName + " port=" + c.Port + " " + c.Config
}

// LinkDsn 根据 dbname 生成 dsn
func (c *PgsqlConfig) LinkDsn(dbname string) string {
	return "host=" + c.Path + " user=" + c.Username + " password=" + c.Password + " dbname=" + dbname + " port=" + c.Port + " " + c.Config
}
