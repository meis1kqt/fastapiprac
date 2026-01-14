package storage

import (

	"github.com/jmoiron/sqlx"
)


func Connect(databaseURl string) (*sqlx.DB, error) {

	db, err  := sqlx.Connect("postgres",databaseURl)

	if err != nil {
		return nil, err 
	}
	 
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(5 )

	return db, nil
}