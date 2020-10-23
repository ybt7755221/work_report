package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"work_report/config"
	"sync"
)

//var Engine *xorm.Engine
var EngineGroup map[string]*xorm.Engine
var once *sync.Once

//自动加载mysql连接
func init() {
	EngineGroup = map[string]*xorm.Engine{}
	for key, _ := range config.MysqlConfMap {
		EngineGroup[key] = connect(key)
	}
	fmt.Println("Engine Group:")
}

func GetDB(key string) *xorm.Engine {
	//如果自动加载mysql实效，重新加载
	if EngineGroup[key] == nil {
		fmt.Println("reInit : " + key)
		connect(key)
	}
	db, _ := EngineGroup[key].Clone()
	return db
}

//连接数据库--单例模式
func connect(key string) *xorm.Engine {
	var err error
	confMap := config.MysqlConfMap
	if err != nil {
		fmt.Println("connect db "+key+" Error: ", err.Error())
	}
	addrStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&loc=Local",
		confMap[key].User,
		confMap[key].Passwd,
		confMap[key].Host,
		confMap[key].Port,
		confMap[key].Name,
		confMap[key].Charset,
	)
	fmt.Println("DB Addr : " + addrStr)
	engine, err := xorm.NewEngine("mysql", addrStr)
	engine.SetMaxOpenConns(confMap[key].OpenMax)
	engine.SetMaxIdleConns(confMap[key].IdleMax)
	engine.ShowSQL(true)
	if err != nil {
		fmt.Println("Connect DB "+key+" Error :", err.Error())
	} else {
		fmt.Println("Connect DB " + key + " Success!")
	}
	return engine
}
