package config

type Config struct {
	Pgsql PgsqlConfig `mapstructure:"pgsql" json:"pgsql" yaml:"pgsql"`
}
