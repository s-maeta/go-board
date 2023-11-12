package database

import (
	"board/config"
	"fmt"
	"log"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var m *migrate.Migrate
var err error

func Init() {
	// config設定を取得
	cfg := config.GetConfig()

	// DBの接続先設定
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		cfg.Db.User,
		cfg.Db.Password,
		cfg.Db.Host,
		cfg.Db.Port,
		cfg.Db.Database,
	)

	dsn_option := fmt.Sprintf(
		"?charset=%s&parseTime=%t&loc=%s",
		cfg.Db.Charset,
		cfg.Db.ParseTime,
		cfg.Db.Loc,
	)

	dsn_mysql := dsn + dsn_option

	// DBに接続(指定回数回接続試行する)
	count := 10
	for count > 1 {
		if db, err = gorm.Open(mysql.Open(dsn_mysql), &gorm.Config{}); err != nil {
			time.Sleep(time.Second * 2)
			count--
			log.Printf("retry... count:%v\n", count)
			continue
		}
		err = nil
		break
	}
	if err != nil {
		panic(err)
	}

	// マイグレーション設定
	dsn_migration := fmt.Sprintf(
		"%s://%s",
		cfg.Db.Type,
		dsn,
	)
	m, err = migrate.New(
		cfg.Migrate.FilePath,
		dsn_migration,
	)
	if err != nil {
		panic(err)
	}

	// マイグレーションの実行
	err = m.Up()
	// エラーメッセージが「no change」の場合もスキップ
	if err != nil && err.Error() != "no change" {
		log.Printf("m.Up() Error Message: %s\n", err)
	}
}

func GetDB() *gorm.DB {
	return db
}

func GetM() *migrate.Migrate {
	return m
}

func Close() {
	getDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	getDB.Close()
}
