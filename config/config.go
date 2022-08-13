package config

import (
	"net/url"
	"os"
)

type Database struct {
    User string
    Password string
    Scheme string
    Host string
    Path string
    RawQuery string
}

type Config struct {
    Db Database
    Port int
}

var ServerConfig = Config{
    Port: 8080,
    Db: Database{
        User: "postgres",
        Password: os.Getenv("POSTGRES_PASSWORD"),
        Scheme: "postgres",
        Host: "localhost:5432",
        Path: "gorm_http",
        RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
    },
}
