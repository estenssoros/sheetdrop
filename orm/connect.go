package orm

import (
	"context"
	"os"
	"time"

	"github.com/estenssoros/sheetdrop/constants"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	_ "gorm.io/driver/mysql" //mysql driver
	"gorm.io/gorm"
)

// ConnectTimeout times out after 5 seconds
func ConnectTimeout() (*gorm.DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	ch := make(chan struct {
		conn *gorm.DB
		err  error
	}, 1)
	go func() {
		conn, err := connect()
		ch <- struct {
			conn *gorm.DB
			err  error
		}{conn, errors.Wrap(err, "connect db")}
	}()
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case pack := <-ch:
		return pack.conn, pack.err
	}
}

// Connect wrapper for connect timeout
func Connect() (*gorm.DB, error) {
	start := time.Now()
	defer func() {
		logrus.Infof("latency: %v", time.Since(start))
	}()
	return ConnectTimeout()
}

// Connect connect to a database environment
func connect() (*gorm.DB, error) {
	return connectConfig(&config{
		Database: os.Getenv(constants.EnvDataBaseName),
		Host:     os.Getenv(constants.EnvDataBaseHost),
		User:     os.Getenv(constants.EnvDataBaseUser),
		Password: os.Getenv(constants.EnvDataBasePassword),
		Dialect:  os.Getenv(constants.EnvDatabaseDialect),
	})
}

func connectConfig(conf *config) (*gorm.DB, error) {
	logrus.Info(conf)
	url, err := conf.URL()
	if err != nil {
		return nil, errors.Wrap(err, "conf.URL")
	}
	db, err := gorm.Open(conf.driverName(), url)
	if err != nil {
		return nil, errors.Wrap(err, "open db")
	}
	db.DB().SetMaxIdleConns(0)
	db.LogMode(true)
	if err := db.DB().Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
