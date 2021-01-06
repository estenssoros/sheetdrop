package orm

import (
	"context"
	"os"
	"time"

	"github.com/estenssoros/sheetdrop/constants"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// ConnectConfig times out after 5 seconds
func ConnectConfig(config *Config) (*gorm.DB, error) {
	if config.Database == "" {
		return nil, errors.New("blank database")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	ch := make(chan struct {
		db  *gorm.DB
		err error
	}, 1)
	go func() {
		db, err := openConnection(config)
		ch <- struct {
			db  *gorm.DB
			err error
		}{db, errors.Wrap(err, "openConnection")}
	}()
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case pack := <-ch:
		return pack.db, pack.err
	}
}

type connector interface {
	URL() string
	dialector() (gorm.Dialector, error)
	database() string
}

func openConnection(conn connector) (*gorm.DB, error) {
	// start := time.Now()
	// defer func() {
	// 	logrus.Infof("latency: %v", time.Since(start))
	// }()
	// logrus.Infof("connecting to: %s", conn.database())
	dialect, err := conn.dialector()
	if err != nil {
		return nil, errors.Wrap(err, "conn.dialector")
	}
	db, err := gorm.Open(dialect, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, errors.Wrap(err, "gorm.Open")
	}
	// logrus.Infof("dialect: %s", db.Dialector.Name())
	sqlDB, err := db.DB()
	if err != nil {
		return nil, errors.Wrap(err, "db.DB")
	}
	sqlDB.SetMaxIdleConns(1)
	sqlDB.SetMaxOpenConns(10)
	return db, nil
}

// Connect connect to a database by name
func Connect() (*gorm.DB, error) {
	// if err := godotenv.Load(); err != nil {
	// 	logrus.Warning("could not find .env")
	// }
	config := &Config{
		Dialect:  os.Getenv(constants.EnvDatabaseDialect),
		Database: os.Getenv(constants.EnvDataBaseName),
		Host:     os.Getenv(constants.EnvDataBaseHost),
		User:     os.Getenv(constants.EnvDataBaseUser),
		Password: os.Getenv(constants.EnvDataBasePassword),
	}
	db, err := ConnectConfig(config)
	if err != nil {
		return nil, errors.Wrap(err, "ConnectConfig")
	}
	return db, nil
}
