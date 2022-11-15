package util

import (
	"fmt"

	"github.com/BounkBU/doonungfangpleng/config"
)

func NewConnectionUrlBuilder(db config.Database) string {
	url := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		db.Username,
		db.Password,
		db.Hostname,
		db.Port,
		db.Database,
	)
	return url
}
