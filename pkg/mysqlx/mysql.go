package mysqlx

import (
	"fmt"
	"time"

	"github.com/jacobintern/meme_coin/pkg/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Client struct {
	DB *gorm.DB
}

func NewClient(conf *conf.Config) *Client {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True",
		conf.Database.MySql.User,
		conf.Database.MySql.Password,
		conf.Database.MySql.Host,
		conf.Database.MySql.Port,
		conf.Database.MySql.Database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		QueryFields: true,
	})
	if err != nil {
		panic(err)
	}

	context, err := db.DB()
	if err != nil {
		panic(err)
	}

	context.SetConnMaxLifetime(120 * time.Second)
	context.SetMaxIdleConns(10)
	context.SetMaxOpenConns(25)

	return &Client{DB: db}
}
