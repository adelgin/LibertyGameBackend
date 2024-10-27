package config

import (
	"fmt"
	"time"

	"github.com/jessevdk/go-flags"
)

type Db struct {
	Host            string        `long:"host" env:"HOST" description:"Postgres host" required:"yes"`
	Port            string        `long:"port" env:"PORT" description:"Postgres port" required:"yes"`
	User            string        `long:"user" env:"USER" description:"Postgres user" required:"yes"`
	Password        string        `long:"password" env:"PASSWORD" description:"Postgres password" required:"yes"`
	Name            string        `long:"name" env:"NAME" description:"Postgres name" required:"yes"`
	MaxOpenConns    int           `long:"max-open-conns" env:"MAX_OPEN_CONNS" default:"25" description:"maximum of open database connections"`
	MaxIdleConns    int           `long:"max-idle-conns" env:"MAX_IDLE_CONNS" default:"10" description:"maximum of idle database connections"`
	ConnMaxLifeTime time.Duration `long:"conn-max-life-time" env:"CONN_MAX_LIFE_TIME" default:"5m" description:"database max connection life time"`
}

type LogConfig struct {
	Level string `short:"l" long:"level" env:"LEVEL" description:"Logger level" required:"yes" default:"error"` // std: trace, debug, info, warning, error, fatal, panic
}

type MainBackendConfig struct {
	Host string `long:"host" env:"HOST" description:"Main backend host" required:"yes"`
	Port string `long:"port" env:"PORT" description:"Main backend port" required:"yes"`
}

func Parse() (*Config, error) {
	var config Config
	p := flags.NewParser(&config, flags.HelpFlag|flags.PassDoubleDash)

	_, err := p.ParseArgs([]string{})
	if err != nil {
		return nil, err
	}

	return &config, nil
}

type Config struct {
	Environment       string             `long:"environment" env:"ENVIRONMENT" default:"test"`
	Debug             bool               `long:"debug" env:"DEBUG"`
	Timeout           int                `long:"timeout" env:"TIMEOUT" default:"1000000"`
	MainBackendConfig *MainBackendConfig `group:"Main backend args" namespace:"mainbackend" env-namespace:"MAIN_BACKEND"`
	Log               *LogConfig         `group:"Logger args" namespace:"logger" env-namespace:"LOGGER"`
	Db                *Db                `group:"database args" namespace:"db" env-namespace:"POSTGRES"`
}

func (c *Db) ConnectionString() string {
	uri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		c.Host, c.Port,
		c.User, c.Name,
		c.Password,
	)

	return uri
}
