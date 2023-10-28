package postgre

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"time"
)

type Config struct {
	Uname      string
	Pass       string
	NameDB     string
	Host       string
	Port       string
	SSL        string
	DriverName string
}

func Migrate(cfg *Config) (*sql.DB, error) {
	log.Infoln("Migrate database...")
	db, err := InitDB(cfg)
	if err != nil {
		log.Errorf("Failed init database:%s", err.Error())
	}
	mig, err := migrate.New(
		"file://internal/repository/postgre/sql",
		fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=%s",
			cfg.DriverName, cfg.Uname, cfg.Pass, cfg.Host, cfg.Port, cfg.NameDB, cfg.SSL))
	if err != nil {
		log.Errorf("Failed database migration:%s", err.Error())
		return nil, err
	}
	err = mig.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Errorf("Error Up migrate database:%s", err.Error())
		return nil, err
	}
	return db, nil
}

func InitDB(cfg *Config) (*sql.DB, error) {
	log.Infoln("Init database...")
	db, err := sql.Open(cfg.DriverName, fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Uname, cfg.NameDB, cfg.Pass, cfg.SSL))
	if err != nil {
		log.Errorf("Failed init database:%s", err.Error())
		return nil, err
	}

	log.Infoln("Ping database...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		log.Errorf("Failed ping database:%s", err.Error())
		return nil, err
	}
	return db, nil
}
