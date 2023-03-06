package setting

import "time"

type App struct {
	JwtSecret string
	PageSize  int
	PrefixUrl string

	RuntimeRootPath string

	// ExportSavePath string
	// FontSavePath string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
}

var AppSetting = &App{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Database struct {
	// Type        string
	User        string
	Password    string
	Host        string
	Port        string
	Name        string
	TablePrefix string
}

var DatabaseSetting = &Database{}

type Redis struct {
	Host        string
	Port        string
	Password    string
	MaxIdle     string
	MaxActive   string
	IdleTimeout time.Duration
}

var RedisSetting = &Redis{}

func setup() {

}
