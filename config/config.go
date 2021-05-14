package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	fileman "xcheck/filemanager"
	"xcheck/gin"
	timeutil "xcheck/utils"
)

const (
	DEVELOPMENT string = "dev"
	PRODUCTION  string = "prod"
)

type ServerConfig struct {
	Port string `toml:"port"`
}

type DatabaseConfig struct {
	Host     string `toml:"host"`
	Username string `toml:"username"`
	Pass     string `toml:"pass"`
	Port     string `toml:"port"`
	DbName   string `toml:"dbname"`
}

type App struct {
	StageMode     string `toml:"stageMode"`
	Version       string `toml:"version"`
	MailSender    string `toml:"mailSender"`
	MailPass      string `toml:"mailPass"`
	SmtpHost      string `toml:"smtpHost"`
	SmtpPort      string `toml:"smtpPort"`
	ReportMail    string `toml:"reportMail"`
	Pwd           string
	PidFile       string
	StartDateTime string
	StartUnix     int64
	GinVersion    string
}

type ProjectConfig struct {
	ConfDir    string
	DotEnvDir  string
	HttpServer ServerConfig   `toml:"http_server"`
	DB         DatabaseConfig `toml:"db"`
	App        App            `toml:"app"`
}

type Service struct {
	Name string `toml:"name"`
	Url  string `toml:"url"`
}

type ServiceArray struct {
	Interval int32     `toml:"interval"`
	Services []Service `toml:"service"`
}

func (p *ProjectConfig) init() {
	p.ConfDir = "app.conf"
	p.DotEnvDir = ".env"
	p.App.StageMode = PRODUCTION
	p.App.Pwd = fileman.Pwd()
	p.App.PidFile = p.App.Pwd + "/" + "killer.pid"
	p.App.StartUnix = timeutil.Unix()
	p.App.StartDateTime = timeutil.Now()
	p.App.GinVersion = gin.Version
}

var CONFIG *ProjectConfig
var Services *ServiceArray

func LoadConfig() {
	var c ProjectConfig
	c.init()
	if !fileman.FileExists(c.ConfDir) {
		panic("there is no config file " + c.ConfDir)
	}

	if !fileman.FileExists(c.ConfDir) {
		panic("there is no config environment file " + c.DotEnvDir)
	}

	if _, err := toml.DecodeFile(c.ConfDir, &c); err != nil {
		fmt.Println(err)
		return
	}

	if _, err := toml.DecodeFile(c.DotEnvDir, &c); err != nil {
		fmt.Println(err)
		return
	}

	var services ServiceArray
	if _, err := toml.DecodeFile("services.conf", &services); err != nil {
		fmt.Println(err)
		return
	}

	Services = &services
	CONFIG = &c

	//fmt.Println("loaded services:")
	//fmt.Println(*Services)
}
