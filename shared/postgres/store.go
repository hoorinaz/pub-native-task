package postgres

import (
	"fmt"
	"github.com/caarlos0/env"
	"github.com/gocraft/dbr/v2"
	_ "github.com/lib/pq"
	"log"

	"time"
)

type (
	Conf struct {
		Host        string `env:"DB_PG_HOST" envDefault:"localhost:5432"`
		Username    string `env:"DB_PG_USER" envDefault:"postgres"`
		Password    string `env:"DB_PG_PASS" envDefault:"rootPassword"`
		DBName      string `env:"DB_PG_NAME" envDefault:"postgres"`
		SSLMode     string `env:"DB_PG_SSL"  envDefault:"disable"`
		MaxOpenConn int    `env:"DB_PG_MAX_OPEN_CONN"  envDefault:"5"`
		MaxIdleConn int    `env:"DB_PG_MAX_IDLE_CONN"  envDefault:"5"`
		MaxConnLife int    `env:"DB_PG_MAX_CONN_LIFE"  envDefault:"5000"`
	}
)

func (c *Conf) String() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=%s",
		c.Username,
		c.Password,
		c.Host,
		c.DBName,
		c.SSLMode,
	)
}

func NewSession() *dbr.Session {
	c := new(Conf)
	_ = env.Parse(c)
	connection, err := dbr.Open("postgres", c.String(), nil)
	if nil != err {
		log.Print("error opening postgres connection", err.Error())
		panic(err.Error())
	}

	if err = connection.Ping(); nil != err {
		log.Println("error dialing postgres db", err.Error())
		panic(err.Error())
	}

	connection.SetMaxOpenConns(c.MaxOpenConn)
	connection.SetMaxIdleConns(c.MaxIdleConn)
	connection.SetConnMaxLifetime(time.Duration(c.MaxConnLife) * time.Millisecond)

	return connection.NewSession(nil)
}
