package config

import (
	"fmt"
	"io/fs"
	"log"
	"path"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	DatabseHost  string `mapstructure:"DATABASE_HOST"`
	DatabaseUser string `mapstructure:"DATABASE_USER"`
	DatabasePass string `mapstructure:"DATABASE_PASS"`
	DatabaseName string `mapstructure:"DATABASE_NAME"`
	DatabasePort string `mapstructure:"DATABASE_PORT"`

	TokenSecret string `mapstructure:"TOKEN_SECRET"`

	Port string `mapstructure:"PORT"`
}

func (c *Config) String() string {
	return "DatabaseHost: " + c.DatabseHost + "\n" +
		"DatabaseUser: " + c.DatabaseUser + "\n" +
		"DatabasePass: " + c.DatabasePass + "\n" +
		"DatabaseName: " + c.DatabaseName + "\n" +
		"DatabasePort: " + c.DatabasePort + "\n" +
		"TokenSecret: " + c.TokenSecret + "\n" +
		"Port: " + c.Port
}

var AppConfig *Config

func LoadConfig(configPath string) *Config {
	viper.AddConfigPath(configPath)
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		switch err.(type) {
		case viper.ConfigFileNotFoundError:
		case *fs.PathError:
			viper.SetDefault("DATABASE_HOST", "db")
			viper.SetDefault("DATABASE_PORT", "5432")
			viper.SetDefault("DATABASE_USER", "postgres")
			viper.SetDefault("DATABASE_PASS", "")
			viper.SetDefault("DATABASE_NAME", "")
			viper.SetDefault("PORT", "8080")
			viper.SetDefault("TOKEN_SECRET", GenerateTokenSecret())

			err = viper.WriteConfigAs(path.Join(configPath, ".env"))

			if err != nil {
				log.Fatalf("Error writing config file, %s", err)
			}
			break
		default:
			log.Printf("Error reading config file, %s", err)
			break
		}

	}

	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		fmt.Println("Server is running on port", viper.Get("PORT"))
	})

	viper.WatchConfig()

	return AppConfig
}

func GetConfig() *Config {
	if AppConfig == nil {
		log.Fatal("Configuration not loaded. Call LoadConfig first.")
	}
	return AppConfig
}

func SetVariable(key string, value string) *Config {
	viper.Set(key, value)
	return AppConfig
}
