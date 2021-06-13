package config

type Config struct {
	User     UserConfig     `json:"user"`
	Document DocumentConfig `json:"document"`
	Database DatabaseConfig `json:"database"`
}

type UserConfig struct {
	Userdb string `json:"userdb"`
}

type DocumentConfig struct {
	Documentdb string `json:"documentdb"`
}

type DatabaseConfig struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
}

const configPath = `C:\Users\gmalik\Documents\Workspace\Learning\watermark-config.yml`

var Configuration Config

func LoadConfig() error {
	b := GetBuilder(DefaultBuilder)
	b.SetConfigPath(configPath)

	if err := b.ReadConfig(&Configuration); err != nil {
		return err
	}
	return nil
}
