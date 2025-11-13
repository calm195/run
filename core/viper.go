package core

import (
	"flag"
	"fmt"
	"path/filepath"
	"run/global"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Viper() *viper.Viper {
	configPath := getConfigPath()

	v := viper.New()
	v.SetConfigFile(configPath)
	v.SetConfigType("yaml")

	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal: error config file: %w", err))
	}
	v.WatchConfig()

	if err = v.Unmarshal(&global.Config); err != nil {
		panic(fmt.Errorf("fatal error unmarshal config: %w", err))
	}

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&global.Config); err != nil {
			fmt.Println(err)
		}
	})

	return v
}

// getConfigPath
//
//	@Description: 获取配置文件路径，默认为dev目录下配置文件。
//	其中，
//	命令行参数 `-f filename` 用于指定配置文件，默认为 `core.DefaultConfigFileName`
//		`-p` 用于开启发布模式，不需要传值。
//	@return configPath
func getConfigPath() string {
	configDir := DevDir
	fileName := DefaultConfigFileName
	fmt.Println("config dir:", configDir)
	fmt.Println("config file:", fileName)

	flag.Parse()
	if flag.Lookup(ProdFlag) != nil {
		configDir = ProdDir
	}
	if flag.Lookup(FileFlag) != nil {
		fileName = *flag.String(FileFlag, DefaultConfigFileName, "config file name")
		fileName = fileName + ConfigFileSuffix
	}
	fmt.Println("config dir:", configDir)
	fmt.Println("config file:", fileName)

	configPath := filepath.Join(configDir, fileName)
	return configPath
}
