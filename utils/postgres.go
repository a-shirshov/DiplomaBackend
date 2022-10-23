package utils

import (
	"fmt"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

func buildConnectionString(user,password, host, port, database string) string {
	connectionString := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", user, password, host, port, database)
	return connectionString
}

func InitPostgres() (*sqlx.DB, error) {
	user := viper.GetString("POSTGRES_USER")
	database := viper.GetString("POSTGRES_DB")
	password := viper.GetString("POSTGRES_PASSWORD")
	host := viper.GetString("postgres.host")
	port := fmt.Sprintf("%d",viper.GetInt("postgres.port"))

	config, err := pgxpool.ParseConfig(buildConnectionString(user,password,host,port,database))
	if err != nil {
		log.Print(err)
		return nil, err
	}

	nativeDB := stdlib.OpenDB(*config.ConnConfig)
	nativeDB.SetMaxIdleConns(100)
	nativeDB.SetMaxOpenConns(100)
	
	return sqlx.NewDb(nativeDB,"pgx"), nil
	
}

// func Prepare(db *pgxpool.Pool) error {
// 	tx, err := db.Begin(context.Background())
// 	if err != nil {
// 		return err
// 	}

// 	for _, query := range queries {
// 		_, err := tx.Prepare(context.Background(), query.Name, query.Query)
// 		if err != nil {
// 			tx.Rollback(context.Background())
// 			return err
// 		}
// 	}
// 	tx.Commit(context.Background())
// 	return nil
// }