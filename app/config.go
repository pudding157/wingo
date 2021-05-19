package app

import (
	"winapp/models"

	"github.com/spf13/viper"
)

// Config struct
type Config struct {
	Env string `mapstructure:"env"`
	vp  *viper.Viper
	Db  *Database
}

// Database struct
type Database struct {
	Bank        models.Bank
	User        []models.User
	Otp_history models.Otp_history
}

// NewConfig func
func NewConfig(env string) *Config {
	return &Config{Env: env}
}

//Init func
func (c *Config) Init() error {
	vp := viper.New()
	vp.AddConfigPath("./config")
	vp.SetConfigName(c.Env)
	c.vp = vp
	// c.Db = &Database{
	// 	Merchants: []models.Merchant{{
	// 		ID:           1,
	// 		Name:         "test",
	// 		BangkAccount: "1234",
	// 		Username:     "abc1234",
	// 		Password:     "1234",
	// 	}},
	// 	Products: []models.Product{},
	// }
	c.Db = &Database{}
	if err := vp.ReadInConfig(); err != nil {
		return err
	}
	if err := c.binding(); err != nil {
		return err
	}
	return nil
}

func (c *Config) binding() error {
	if err := c.vp.Unmarshal(&c); err != nil {
		return err
	}
	return nil
}
