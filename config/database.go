package config

import (
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

// general get database
func GetDatabase(dialector gorm.Dialector, maxIdleConns, maxOpenConns int) *gorm.DB {
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		logrus.Fatal(err)
		return db
	}

	db.Use(dbresolver.Register(dbresolver.Config{
		Policy:            dbresolver.RandomPolicy{},
		TraceResolverMode: true,
	}).SetConnMaxIdleTime(time.Hour).
		SetConnMaxLifetime(24 * time.Hour).
		SetMaxIdleConns(maxIdleConns).
		SetMaxOpenConns(maxOpenConns))
	return db
}
