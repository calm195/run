package config

type SystemConfig struct {
	Port         int    `mapstructure:"port" json:"port" yaml:"port"`
	RouterPrefix string `mapstructure:"router-prefix" json:"router-prefix" yaml:"router-prefix"`
}
