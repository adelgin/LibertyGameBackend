package config

import (
	"fmt"
	"time"
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

func (c *Db) ConnectionString() string {
	uri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		c.Host, c.Port,
		c.User, c.Name,
		c.Password,
	)

	return uri
}
