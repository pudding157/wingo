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
	Env string `mapstructure:"env"`
	vp  *viper.Viper
	DB  *gorm.DB
	R   *redis.Client
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
	// vp := viper.New()
	// vp.AddConfigPath("./config")
	// vp.SetConfigName(c.Env)
	// c.vp = vp
	// if err := vp.ReadInConfig(); err != nil {
	// 	return err
	// }

	c.DB = connectDatabase()
	c.R = connectRedis()
	// if err := c.binding(); err != nil {
	// 	fmt.Println("binding err : => ", err)
	// 	return err
	// }

	return nil
}

func (c *Config) binding() error {
	if err := c.vp.Unmarshal(&c); err != nil {
		return err
	}
	return nil
}
func connectDatabase() *gorm.DB {
	//db:3306
	db, err := gorm.Open("mysql", "root:helloworld@tcp(localhost:6603)/godb?charset=utf8&parseTime=True") //127.0.0.1:3306
	if err != nil {
		log.Fatal(err)
	}
	return db
}
func connectRedis() *redis.Client {
	log.Printf("hello redis")
	//rediss
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
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
