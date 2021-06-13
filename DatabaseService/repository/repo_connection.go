package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/apoorvabhargava/kubernetes-watermark-application/database-service/config"
	"github.com/jackc/pgx/v4"
)

type DBConnection struct {
	client   *pgx.Conn
	dbconfig config.Config
}

func (dbc *DBConnection) CreateConnection(appEnv string) error {
	dbName := ""
	if strings.ToLower(appEnv) == "user" {
		dbName = dbc.dbconfig.User.Userdb
	} else {
		dbName = dbc.dbconfig.Document.Documentdb
	}
	client, err := pgx.Connect(context.Background(), getConnectionString(dbName, dbc.dbconfig.Database.Host, dbc.dbconfig.Database.UserName, dbc.dbconfig.Database.Password,
		dbc.dbconfig.Database.Port))
	if err != nil {
		return err
	}
	dbc.client = client
	return nil
}

func getConnectionString(dbName, hostname, username, password string, port string) string {
	connectionString := fmt.Sprintf("postgres://{%s}:{%s}@{%s}:{%s}/{%s}", username, password, hostname, port, dbName)
	return connectionString
}
