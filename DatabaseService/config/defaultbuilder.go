package config

import (
	"errors"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type defaultBuilder struct {
	V *viper.Viper
}

func newdefaultBuilder() *defaultBuilder {
	viperInit := viper.New()

	if defaultLoadEnv {
		viperInit.AutomaticEnv()
	}
	viperInit.SetConfigType(defaultConfigType)
	viperInit.WatchConfig()
	return &defaultBuilder{V: viperInit}
}

func (builder *defaultBuilder) SetConfigPath(path string) {
	builder.V.SetConfigFile(path)
}

func (builder *defaultBuilder) SetConfigType(configType string) {
	builder.V.SetConfigType(configType)
}

func (builder *defaultBuilder) SetReadEnv(readEnv bool) {
	if readEnv {
		builder.V.AutomaticEnv()
	}
}

func (builder *defaultBuilder) ReadConfig(input interface{}) error {
	if err := builder.V.ReadInConfig(); err != nil {
		errMsg := "Error occured while reading the config file"
		return errors.New(errMsg)
	}
	err := builder.V.Unmarshal(&input)
	if err != nil {
		return errors.New("Error while converting the config to input config fields")
	}
	builder.V.OnConfigChange(func(in fsnotify.Event) {
		if in.Op == fsnotify.Write {
			if err := viper.Unmarshal(&input); err != nil {
			}
		}
	})

	return nil
}
