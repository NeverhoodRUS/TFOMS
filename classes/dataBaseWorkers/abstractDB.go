package databaseworkers

import (
	"database/sql"
	"errors"
	"fmt"
	"tfoms_server/static/exceptions"
	"time"

	"github.com/beego/beego/v2/core/config"

	_ "github.com/lib/pq"
)

var (
	host     string = ""
	port     int
	user     string = ""
	password string = ""
	dbname   string = ""

	max_conn  int = 40
	idle_conn int = 10

	db *sql.DB // Глобальное соединение sql

	DB_CONN_ERROR error
)

func init() {
	host, _ = config.String("db_host")
	port, _ = config.Int("db_port")
	user, _ = config.String("db_user")
	password, _ = config.String("db_password")
	dbname, _ = config.String("db_name")

	max_conn = config.DefaultInt("db_max_conn", 50)
	idle_conn = config.DefaultInt("db_idle_conn", 10)

	DB_CONN_ERROR = errors.New(exceptions.DB_connection_error)
}

func getPostgresSession() *sql.DB {
	if db == nil {
		psqlInfo := fmt.Sprintf(`host=%s port=%d user=%s password=%s dbname=%s sslmode=disable`, host, port, user, password, dbname)

		dbOpen, err := sql.Open("postgres", psqlInfo)
		if err != nil {
			return nil
		}

		dbOpen.SetConnMaxLifetime(30 * time.Minute)
		dbOpen.SetMaxOpenConns(max_conn)
		dbOpen.SetMaxIdleConns(idle_conn)

		//		err = db.Ping()
		//		if err != nil {
		//			return nil
		//		}

		db = dbOpen
	}

	return db
}

type AbstracterDB interface {
	GetRowStruct(any ...any) *interface{}
}

func GetAllRows(table string) (*sql.Rows, error) {
	return getPostgresSession().Query("select * from ?", table)
}

func getRow(table string, where string) *sql.Row {
	var query string = "select * from " + table + " where " + where
	return getPostgresSession().QueryRow(query)
}

// func rowsToMap(rows *sql.Rows) (*[]map[string]interface{}, error) {
// 	columns, _ := rows.Columns()
// 	columnLength := len(columns)
// 	cache := make([]interface{}, columnLength) //temporarily store each row of data
// 	for index := range cache {                 // initialize a pointer for each column
// 		var a interface{}
// 		cache[index] = &a
// 	}
// 	var list []map[string]interface{} //returned slice
// 	for rows.Next() {
// 		_ = rows.Scan(cache...)

// 		item := make(map[string]interface{})
// 		for i, data := range cache {
// 			item[columns[i]] = *data.(*interface{})
// 		}
// 		list = append(list, item)
// 	}
// 	_ = rows.Close()
// 	return &list, nil
// }
