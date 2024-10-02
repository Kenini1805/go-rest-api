package postgres

import (
	"fmt"
	"time"

	_ "github.com/jackc/pgx/stdlib" // pgx driver.

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/Kenini1805/go-rest-api/config"
)

const (
	maxOpenConns    = 60
	connMaxLifetime = 120
	maxIdleConns    = 30
	connMaxIdleTime = 20
)

// Return new Postgresql db instance.
func NewPsqlDB(c *config.Config) (*gorm.DB, error) {
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		c.Postgres.PostgresqlHost,
		c.Postgres.PostgresqlPort,
		c.Postgres.PostgresqlUser,
		c.Postgres.PostgresqlDbname,
		c.Postgres.PostgresqlPassword,
	)

	// Open a new connection to the database using GORM
	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Configure the connection pool
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(maxOpenConns)
	sqlDB.SetConnMaxLifetime(connMaxLifetime * time.Second)
	sqlDB.SetMaxIdleConns(maxIdleConns)
	sqlDB.SetConnMaxIdleTime(connMaxIdleTime * time.Second)

	// Ping the database to verify connection
	if err = sqlDB.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Gagal untuk keluar koneksi database")
	}
	dbSQL.Close()
}
