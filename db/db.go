package db

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq" // INIT PostgreSQL drivers
	"github.com/xo/dburl"
)

// DB is the database
var DB *gorm.DB

// parseDBURL returns parses a database url into individual components
func parseDBURL(url string) (string, string, error) {
	u, err := dburl.Parse(url)
	if err != nil {
		return "", "", err
	}

	dsn := u.DSN
	switch u.Driver {
	case "postgres":
		dsn = fmt.Sprintf("%s sslmode=disable", dsn)
	default:
		return "", "", errors.New("dialect not supported")
	}

	return u.Driver, dsn, nil
}

// Init sets up the database
func Init(url string) error {
	dialect, uri, err := parseDBURL(url)
	if err != nil {
		return err
	}

	conn, err := gorm.Open(dialect, uri)
	if err != nil {
		return err
	}

	DB = conn
	return nil
}
