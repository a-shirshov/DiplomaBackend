package utils

import (
	"github.com/jackc/pgx"
	"github.com/spf13/viper"
)

func InitPostgresDB() (*pgx.ConnPool, error) {
	config := pgx.ConnConfig{
		User:                 viper.GetString("POSTGRES_USER"),
		Database:             viper.GetString("POSTGRES_DATABASE"),
		Password:             viper.GetString("POSTGRES_PASSWORD"),
		Host: 				  viper.GetString("postgres.host"),
		Port: 				  uint16(viper.GetInt("postgres.port")),	
		PreferSimpleProtocol: false,
	}
	connPoolConfig := pgx.ConnPoolConfig{
		ConnConfig:     config,
		MaxConnections: 100,
		AfterConnect:   nil,
		AcquireTimeout: 0,
	}
	return pgx.NewConnPool(connPoolConfig)
}

func Prepare(db *pgx.ConnPool) error {
	for _, query := range queries {
		_, err := db.Prepare(query.Name, query.Query)
		if err != nil {
			return err
		}
	}
	return nil
}