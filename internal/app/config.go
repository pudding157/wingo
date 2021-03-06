package app

import (
	"fmt"
	"log"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

// Config struct
type Config struct {
	Env  string `mapstructure:"env"`
	vp   *viper.Viper
	DB   *gorm.DB
	R    *redis.Client
	T    string
	UI   int
	ROLE string
}

// Database struct
// type Database struct {
// 	Bank        models.Bank
// 	User        []models.User
// 	Otp_history models.Otp_history
// }

// NewConfig func
func NewConfig(env string) *Config {
	return &Config{Env: env}
}

//Init func
func (c *Config) Init() error {
	vp := viper.New()
	// vp.AddConfigPath("./config")
	// vp.SetConfigName(c.Env)
	vp.SetConfigFile(".env")
	c.vp = vp
	if err := vp.ReadInConfig(); err != nil {
		return err
	}

	c.DB = c.connectDatabase()
	c.R = c.connectRedis()
	if err := c.binding(); err != nil {
		fmt.Println("binding err : => ", err)
		return err
	}

	Config_db := Config_db{c.DB}
	Config_db.Init()

	return nil
}

func (c *Config) binding() error {
	if err := c.vp.Unmarshal(&c); err != nil {
		return err
	}
	return nil
}
func (c *Config) connectDatabase() *gorm.DB {
	//db:3306
	fmt.Println("vp sql => ", c.vp.GetString("SQL_DB"))

	db, err := gorm.Open(c.vp.GetString("SQL_DIALECT"), c.vp.GetString("SQL_DB")) //127.0.0.1:3306
	if err != nil {
		log.Fatal(err)
	}
	return db
}
func (c *Config) connectRedis() *redis.Client {
	log.Printf("hello redis")
	//rediss
	client := redis.NewClient(&redis.Options{
		Addr:     c.vp.GetString("REDIS_DB"),
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()

	if err != nil {
		log.Fatal("redis error => ", err)
	}
	fmt.Println("pass ? => ?", pong, err)

	return client
}
