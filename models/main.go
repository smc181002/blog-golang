package models

import (
	"log"
	"net/url"
	"time"

	"github.com/smc181002/blog-golang/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

type Model struct {
    ID        uint `json:"id" gorm:"primarykey"`
    CreatedAt time.Time `json:"createdAt"`
    UpdatedAt time.Time `json:"updatedAt"`
    DeletedAt gorm.DeletedAt `json:"deletedAt,omitempty" gorm:"index"`
}

func DBInitialize(ServerConfig config.Config) {
    dsn := url.URL {
        User: url.UserPassword(
            ServerConfig.Db.User, 
            ServerConfig.Db.Password,
        ),
        Scheme: ServerConfig.Db.Scheme,
        Host: ServerConfig.Db.Host,
        Path: ServerConfig.Db.Path,
        RawQuery: ServerConfig.Db.RawQuery,
    }

    var err error
    Db, err = gorm.Open(postgres.Open(dsn.String()), &gorm.Config{})
    if err != nil {
        log.Panicln("Database connection failed", err)
        return 
    }

    Db.AutoMigrate(&User{})
}
