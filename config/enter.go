package config

type Config struct {
	Pgsql  PgsqlConfig  `mapstructure:"pgsql" json:"pgsql" yaml:"pgsql"`
	Zap    ZapConfig    `mapstructure:"zap" json:"zap" yaml:"zap"`
	System SystemConfig `mapstructure:"system" json:"system" yaml:"system"`
}
