package glob

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

var (
	Config *DataBases
)

type DataBases struct {
	Database []string `yaml:"Database"`
	Address  string   `yaml:"Address"`
	Username string   `yaml:"Username"`
	Password string   `yaml:"Password"`
}

func init() {
	InitConfig()
}

func InitConfig() {
	LoadConfig("app.yaml")
}

// LoadConfig ...
func LoadConfig(file string) {
	if Config != nil {
		return
	}

	if _, err := os.Stat(file); os.IsNotExist(err) {
		return
	}

	Config = new(DataBases)

	viper.SetConfigType("yaml")
	viper.SetConfigFile(file)
	err := viper.ReadInConfig()
	if err != nil {
		log.Panicf("Fatal error config file: %v", err)
	}

	viper.Unmarshal(&Config)

}