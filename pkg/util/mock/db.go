package mock

import (
	"fmt"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBConnection() (*gorm.DB, sqlmock.Sqlmock) {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		fmt.Println()

		panic("error ga bisa kloneks ")
	}

	mysql := mysql.New(mysql.Config{
		DSN:                       "sqlmock_db_0",
		DriverName:                "mysql",
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	})

	db, err := gorm.Open(mysql, &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db, mock

}
