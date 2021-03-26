package stx

import (
	"bufio"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"

	"CrawlerAudioFiction/db"

	"gopkg.in/yaml.v2"
)

type MysqlConfig struct {
	Host   string `yaml:"host"`
	Port   string `yaml:"port"`
	Name   string `yaml:"name"`
	Pwd    string `yaml:"pwd"`
	DBname string `yaml:"dbname"`
	Table  string `yaml:"table"`
}

type ServiceConfig struct {
	Mysql MysqlConfig `yaml:"mysql"`
}

type ServiceContext struct {
	RW sync.RWMutex
	DB *sql.DB
}

func NewServiceContext(file string) *ServiceContext {
	sc := newServiceConfig(file)
	return &ServiceContext{
		DB: db.NewMysqlConn(sc.Mysql.Name, sc.Mysql.Pwd, sc.Mysql.Host, sc.Mysql.Port, sc.Mysql.DBname),
	}
}

func newServiceConfig(file string) ServiceConfig {
	var sct ServiceConfig
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal("ioutil.ReadAll err:", err)
	}
	if err := yaml.Unmarshal(bytes, &sct); err != nil {
		log.Fatal("yaml unmarshal error :", err)
	}
	fmt.Println(sct)
	return sct
}

func ReadFile(file string) {

	fileinfo, err := os.Open(file)
	if err != nil {
		log.Println("文件打开失败")
		return
	}
	defer fileinfo.Close()
	scanner := bufio.NewScanner(fileinfo)
	for scanner.Scan() {
		line := scanner.Text()
		_ = strings.Split(line, ",")

	}
	return
}
