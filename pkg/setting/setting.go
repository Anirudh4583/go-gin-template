package setting

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type App struct {
	JwtSecret string `mapstructure:"JwtSecret"`
	PageSize  int    `mapstructure:"PageSize"`
	PrefixUrl string `mapstructure:"PrefixUrl"`

	RuntimeRootPath string `mapstructure:"RuntimeRootPath"`

	// ExportSavePath string
	// FontSavePath string

	LogSavePath string    `mapstructure:"LogSavePath"`
	LogSaveName string    `mapstructure:"LogSaveName"`
	LogFileExt  string    `mapstructure:"LogFileExt"`
	TimeFormat  time.Time `mapstructure:"TimeFormat"`
}

type Server struct {
	RunMode      string        `mapstructure:"RunMode"`
	HttpPort     int           `mapstructure:"HttpPort"`
	ReadTimeout  time.Duration `mapstructure:"ReadTimeout"`
	WriteTimeout time.Duration `mapstructure:"WriteTimeout"`
}

type Database struct {
	// Type        string
	User        string `mapstructure:"User"`
	Password    string `mapstructure:"Password"`
	Host        string `mapstructure:"Host"`
	Port        string `mapstructure:"Port"`
	Name        string `mapstructure:"Name"`
	TablePrefix string `mapstructure:"TablePrefix"`
}

type Redis struct {
	Host        string        `mapstructure:"Host"`
	Port        string        `mapstructure:"Port"`
	Password    string        `mapstructure:"Password"`
	MaxIdle     string        `mapstructure:"MaxIdle"`
	MaxActive   string        `mapstructure:"MaxActive"`
	IdleTimeout time.Duration `mapstructure:"IdleTimeout"`
}

type Setting struct {
	AppSetting      App      `mapstructure:"App"`
	ServerSetting   Server   `mapstructure:"Server"`
	DatabaseSetting Database `mapstructure:"Database"`
	RedisSetting    Redis    `mapstructure:"Redis"`
}

var Config = &Setting{}

func Setup() {
	viper.SetConfigName(".env")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Cannot load the specified config file; %s", err)
	}

	if err := viper.Unmarshal(Config); err != nil {
		log.Fatalf("Cannot unmarshal the config into struct; %s", err)
	}

	Config.ServerSetting.ReadTimeout = Config.ServerSetting.ReadTimeout * time.Second
	Config.ServerSetting.WriteTimeout = Config.ServerSetting.WriteTimeout * time.Second
	Config.RedisSetting.IdleTimeout = Config.RedisSetting.IdleTimeout * time.Second
}
