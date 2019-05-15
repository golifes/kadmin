package config

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/xormplus/xorm"
	"sync"
)

type Config struct {
	DB struct {
		Host    string
		User    string
		Pwd     string
		Db      string
		Show    bool
		Port    int
		MaxOpen int
		MaxIdle int
	}
	Redis struct {
		Dns      string
		MinIdle  int
		PoolSize int
	}

	Mgo struct {
		Dns string
	}
}

var (
	EngDb  *xorm.Engine
	EngRds *redis.Client
)

/**
每个服务里面必须调用这个函数
*/

var global *Config

func LoadGlobalConfig(path string) {
	var once sync.Once
	if global == nil {
		global = &Config{}
	}
	once.Do(func() {
		Load(path, &global)
	})
	global.loadDb()
	//config.loadRedis()
	//config.loadMgo()
}

func GetGlobalConfig() *Config {
	if global == nil {
		return &Config{}
	}

	return global
}

//mysql
func (c *Config) LoadDb() *xorm.Engine {
	fmt.Println(c.DB)
	var err error
	dns := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		c.DB.User,
		c.DB.Pwd,
		c.DB.Host,
		c.DB.Db)

	EngDb, err = xorm.NewEngine("mysql", dns)

	ping := EngDb.Ping()
	if ping != nil || err != nil {
		panic(ping)
	}
	EngDb.SetMaxIdleConns(c.DB.MaxIdle)
	EngDb.SetMaxOpenConns(c.DB.MaxOpen)
	EngDb.ShowSQL(c.DB.Show)

	return EngDb
}

//redis
func (c *Config) loadRedis() {

	EngRds = redis.NewClient(&redis.Options{
		Addr:         c.Redis.Dns,
		Password:     "", // no password set
		DB:           0,  // use default DB
		PoolSize:     c.Redis.PoolSize,
		MinIdleConns: c.Redis.MinIdle,
	})
	_, err := EngRds.Ping().Result()

	//a := EngRds.Info()
	//demo(fmt.Sprintf("%s", a))

	if err != nil {
		panic(err)
	}
}
