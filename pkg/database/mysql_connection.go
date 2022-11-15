package database

import (
	"github.com/BounkBU/doonungfangpleng/config"
	"github.com/BounkBU/doonungfangpleng/pkg/util"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func NewMySQLDatabaseConnection(config config.Database) (*sqlx.DB, error) {
	mysqlUrl := util.NewConnectionUrlBuilder(config)
	db, err := sqlx.Connect("mysql", mysqlUrl)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
