package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/PutskouDzmitry/DbTr/pkg/const_db"
)

//GetConnection it's return a new connection in db
func GetConnection(host, port, user, dbname, password, sslmode string) (*gorm.DB, error) {
	dsn := fmt.Sprintf(const_db.AddInfoForConnection,
		host, port, user, dbname, password, sslmode)
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf(const_db.TroubleWithConnection, err)
	}
	return connection, nil
}
