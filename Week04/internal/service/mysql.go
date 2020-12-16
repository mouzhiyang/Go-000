package service

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"

	"../../conf"
)

var db *gorm.DB

var dbWrite *gorm.DB

// Setup initializes the database instance
func Setup() {
	var err error
	db, err = gorm.Open(conf.DBDataRead.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&loc=Local",
		conf.DBDataRead.User,
		conf.DBDataRead.Password,
		conf.DBDataRead.Host,
		conf.DBDataRead.Name))

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return ""
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(5)
	db.DB().SetMaxOpenConns(10)

	/////////////////////////////////////////
	///             Wr                ///////
	/////////////////////////////////////////

	dbWrite, err = gorm.Open(conf.DBDataWrite.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&loc=Local",
		conf.DBDataWrite.User,
		conf.DBDataWrite.Password,
		conf.DBDataWrite.Host,
		conf.DBDataWrite.Name))

	if err != nil {
		log.Fatalf("models.Setup err DBWrite: %v", err)
	}

	gorm.DefaultTableNameHandler = func(dbWrite *gorm.DB, defaultTableName string) string {
		return ""
	}

	dbWrite.SingularTable(true)
	dbWrite.DB().SetMaxIdleConns(5)
	dbWrite.DB().SetMaxOpenConns(10)
}

// CloseDB closes database connection (unnecessary)
func CloseDB() {
	defer db.Close()

	defer dbWrite.Close()

}
