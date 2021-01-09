package utils

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type GormConfig interface {
	GetPoolMin() uint64
	GetPoolMax() uint64
	GetLogMode() bool
}

func GormMysqlJinzhu(dns string, poolMin int, poolMax int, logMode bool) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", dns)
	if err != nil {
		return nil, err
	}

	db.DB().SetMaxOpenConns(poolMax)
	db.DB().SetMaxIdleConns(poolMin)

	if err := db.DB().Ping(); err != nil {
		return nil, fmt.Errorf("Ping the database=%s err=%v\n", dns, err)
	}

	db.LogMode(logMode)

	return db, nil
}
