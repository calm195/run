package global

import (
	"run/config"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Config *config.Config
	Vp     *viper.Viper
	Log    *zap.Logger
	Db     *gorm.DB
)
