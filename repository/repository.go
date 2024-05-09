package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/spf13/viper"
)

var conn *pgx.Conn

func Connect(conf *viper.Viper) error {
	var err error
	psqlConf := conf.GetString("DATABASE")
	conn, err = pgx.Connect(context.Background(), psqlConf)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return err
	}

	return nil
}
