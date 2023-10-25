package postgre

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
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

func InitdDB(cfg *Config) (*sql.DB, error) {
	log.Infoln("Init database...")
	db, err := sql.Open(cfg.DriverName, fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Uname, cfg.NameDB, cfg.Pass, cfg.SSL))
	if err != nil {
		log.Errorf("Failed init database:%s", err.Error())
		return nil, err
	}
	log.Infoln("Ping database...")
	err = db.Ping()
	if err != nil {
		log.Errorf("Failed ping database:%s", err.Error())
		return nil, err
	}
	return db, nil
}
