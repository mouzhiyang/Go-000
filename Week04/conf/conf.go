package conf

import (
	"log"

	"github.com/go-ini/ini"
)

type Database struct {
	Type         string
	User         string
	Password     string
	Host         string
	Name         string
	MaxIdleConns int
	MaxOpenConns int
}

var DBDataWrite = &Database{}
var DBDataRead = &Database{}

var cfg *ini.File
var confFile string

func SetUp(conf string) {
	confFile = conf
	var err error
	cfg, err = ini.Load(confFile)
	if err != nil {
		log.Fatalln("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}

	mapTo("dbdatawrite", DBDataWrite)
	mapTo("dbdataread", DBDataRead)

}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
