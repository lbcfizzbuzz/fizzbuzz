package datastore

import (
	"database/sql"
	"github.com/lbcfizzbuzz/fizzbuzz/models"
)

// mySQLDatastore a datastore using a MySQL driver
type mySQLDatastore struct {
	dsn string
	db  *sql.DB
}

// Init initialize a new connection pool
func (datastore *mySQLDatastore) Init() error {
	db, err := sql.Open("mysql", datastore.dsn)
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}
	datastore.db = db
	return nil
}

// FindByMostUsedQueryString retrieves the most used requests
func (datastore *mySQLDatastore) FindByMostUsedQueryString() (models.Request, error) {
	request := models.Request{}
	err := datastore.db.QueryRow(`SELECT int1_param, int2_param, limit_param, str1_param, str2_param,
										COUNT(id) AS occurrence
										FROM requests
										GROUP BY int1_param, int2_param, limit_param, str1_param, str2_param
										ORDER BY occurrence DESC LIMIT 1`).Scan(&request.Int1Param,
		&request.Int2Param,
		&request.LimitParam,
		&request.Str1Param,
		&request.Str2Param,
		&request.Count)

	if err != nil && err != sql.ErrNoRows {
		return request, err
	}
	return request, nil
}

// Store takes a Request as parameter and store it to the database
func (datastore *mySQLDatastore) Store(request *models.Request) error {
	tx, err := datastore.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(`INSERT INTO requests (int1_param, int2_param, limit_param, str1_param, str2_param)
							VALUES (?, ?, ?, ?, ?)`,
		request.Int1Param,
		request.Int2Param,
		request.LimitParam,
		request.Str1Param,
		request.Str2Param)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
