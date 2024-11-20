package configs

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"time"
)

var (
	GlobalConfig = &Configuration{}
)

func init() {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Can't read config file: %v", err)
		return
	}

	config := &Configuration{}
	err = viper.Unmarshal(config)
	if err != nil {
		log.Fatalf("Can't unmarshal config: %v", err)
		return
	}

	GlobalConfig = config

	fmt.Println("\033[34m================ Loaded Configuration ================\033[0m")
	fmt.Printf("\033[32mHost:           \033[0m%s\n", GlobalConfig.Host)
	fmt.Printf("\033[32mPort:           \033[0m%d\n", GlobalConfig.Port)
	fmt.Printf("\033[32mJWT Secret:     \033[0m%s\n", GlobalConfig.JwtSecret)
	fmt.Printf("\033[32mJWT Issuer:     \033[0m%s\n", GlobalConfig.JwtIssuer)
	fmt.Printf("\033[32mJWT Expiration: \033[0m%s\n", GlobalConfig.JwtExp)
	fmt.Println("\033[34m=====================================================\033[0m")
	return
}

type Configuration struct {
	Host      string        `mapstructure:"HOST"`
	Port      int           `mapstructure:"PORT"`
	JwtSecret string        `mapstructure:"JWT_SECRET"`
	JwtIssuer string        `mapstructure:"JWT_ISSUER"`
	JwtExp    time.Duration `mapstructure:"JWT_EXPIRATION"`
}
