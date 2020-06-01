package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

//app节的配置
type App struct {
	PageSize        int
	JwtSecret       string
	RuntimeRootPath string
	LogSavePath     string
	LogSaveName     string
	LogFileExt      string

	DateStrFormat  string
	DateTimeFormat string
	TimeFormat     string
	DateFormat     string
}

var AppSetting = &App{}

type Serve struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServeSetting = &Serve{}

var Cfg *ini.File

func Setup() {
	var err error
	Cfg, err = ini.Load("conf/conf.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}

	mapTo("app", AppSetting)
	mapTo("server", ServeSetting)
}
func mapTo(section string, v interface{}) {
	err := Cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
