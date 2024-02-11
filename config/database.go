package config

import (
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

// general get database
func GetDatabase(dialector gorm.Dialector) *gorm.DB {
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		logrus.Fatal(err)
		return db
	}

	maxConnIdel := viper.GetInt("DB_POSTGRES_MAX_CONN_IDLE")
	if maxConnIdel == 0 {
		maxConnIdel = 10
	}

	maxConnOpen := viper.GetInt("DB_POSTGRES_MAX_CONN_OPEN")
	if maxConnOpen == 0 {
		maxConnOpen = 10
	}

	db.Use(dbresolver.Register(dbresolver.Config{
		Policy:            dbresolver.RandomPolicy{},
		TraceResolverMode: true,
	}).SetConnMaxIdleTime(time.Hour).
		SetConnMaxLifetime(24 * time.Hour).
		SetMaxIdleConns(maxConnIdel).
		SetMaxOpenConns(maxConnOpen))
	return db
}
