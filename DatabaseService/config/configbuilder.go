package config

const (
	DefaultBuilder    = "default"
	defaultLoadEnv    = true
	defaultConfigType = "yml"
)

type ConfigBuilder interface {
	SetConfigPath(path string)
	SetConfigType(configType string)
	SetReadEnv(shouldRead bool)
	ReadConfig(config interface{}) error
}

func GetBuilder(builderType string) ConfigBuilder {
	if builderType == DefaultBuilder {
		return newdefaultBuilder()
	}
	return newdefaultBuilder()
}

func GetDefaultConfiguration(input interface{}, configPath string) error {
	b := GetBuilder(DefaultBuilder)
	b.SetConfigPath(configPath)

	return b.ReadConfig(&input)
}
