package util

import "github.com/spf13/viper"

type Config struct{
	DBDriver string `mapstructure:"DB_DRIVER"`
	DBUrl string `mapstructure:"DB_SOURCE"`
	DBAddress string `mapstructure:"DB_ADDRESS"`
}

func LoadConfig(path string) (Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil{
		return Config{}, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil{
		return Config{}, err
	}

	return config, nil

}