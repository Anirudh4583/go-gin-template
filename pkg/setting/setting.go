package setting

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Setting struct {
	// APP
	AppJwtSecret       string `mapstructure:"APP_JWT_SECRET"`
	AppPageSize        int    `mapstructure:"APP_PAGE_SIZE"`
	AppPrefixUrl       string `mapstructure:"APP_PREFIX_URL"`
	AppRuntimeRootPath string `mapstructure:"APP_RUNTIME_ROOTPATH"`
	// ExportSavePath string
	// FontSavePath string
	AppLogSavePath string `mapstructure:"APP_LOG_SAVE_PATH"`
	AppLogSaveName string `mapstructure:"APP_LOG_SAVE_NAME"`
	AppLogFileExt  string `mapstructure:"APP_LOG_FILE_EXT"`
	AppTimeFormat  string `mapstructure:"APP_TIME_FORMAT"`

	// Server
	ServerRunMode      string `mapstructure:"SERVER_RUN_MODE"`
	ServerHttpPort     int    `mapstructure:"SERVER_HTTP_PORT"`
	ServerReadTimeout  string `mapstructure:"SERVER_READ_TIMEOUT"`
	ServerWriteTimeout string `mapstructure:"SERVER_WRITE_TIMEOUT"`

	// Database
	// Type        string
	DBUser        string `mapstructure:"DB_USER"`
	DBPassword    string `mapstructure:"DB_PASSWORD"`
	DBHost        string `mapstructure:"DB_HOST"`
	DBPort        string `mapstructure:"DB_PORT"`
	DBName        string `mapstructure:"DB_NAME"`
	DBTablePrefix string `mapstructure:"DB_TABLE_PREFIX"`

	// Redis
	RedisHost        string `mapstructure:"REDIS_HOST"`
	RedisPort        string `mapstructure:"REDIS_PORT"`
	RedisPassword    string `mapstructure:"REDIS_PASSWORD"`
	RedisMaxIdle     string `mapstructure:"REDIS_MAX_IDLE"`
	RedisMaxActive   string `mapstructure:"REDIS_MAX_ACTIVE"`
	RedisIdleTimeout string `mapstructure:"REDIS_IDLE_TIMEOUT"`
}

var Config = &Setting{}

func Setup() {
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Cannot load the specified config file; %s", err)
	}

	if err := viper.Unmarshal(Config); err != nil {
		log.Fatalf("Cannot unmarshal the config into struct; %s", err)
	}

	fmt.Println(Config)
}
