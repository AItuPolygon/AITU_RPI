package app

import (
	"os"
	"time"

	yaml "gopkg.in/yaml.v3"
)

const (
	ConfigFilePath   = "./build"
	ConfigFileName   = "config"
	ConfigFileFormat = "yaml"
)

type Config struct {
	App         *AppConfig  `yaml:"app"`
	Virtual     *Virtual    `yaml:"virtual"`
	Cert        *CertConfig `yaml:"cert"`
	TasksConfig string      `yaml:"tasks"`
	Modbus      string      `default:"/dev/ttyUSB0" yaml:"modbus"`
}

type Virtual struct {
	Addr string `yaml:"addr"`
	Port string `yaml:"port"`
}

type AppConfig struct {
	Port            int           `yaml:"port"`
	ShutdownTimeout time.Duration `yaml:"shutdown_timeout"`
}

type CertConfig struct {
	DirPath    string    `yaml:"directory_path"`
	RootCAPath string    `yaml:"root_ca_path"`
	ServerCert *KeyPairs `yaml:"server"`
}

type KeyPairs struct {
	PublicKey  string `yaml:"public_key"`
	PrivateKey string `yaml:"private_key"`
}

func (app *App) InitConfig() (err error) {
	if yamlRaw, err := os.ReadFile("config.yml"); err == nil {
		app.Config = new(Config)
		if err = yaml.Unmarshal(yamlRaw, app.Config); err == nil {
			return nil
		}
	}

	return err
}

func (app *App) getRegisterParams(cmd string, params map[string]interface{}) (uint8, uint8, uint8) {

	return 0, 0, 0
}
