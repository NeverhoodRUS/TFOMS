package databaseworkers

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/beego/beego"
	_ "github.com/lib/pq"

	"tfoms_server/static/exceptions"
)

var (
	host     string = ""
	port     int
	user     string = ""
	password string = ""
	dbname   string = ""

	max_conn  int = 40
	idle_conn int = 10

	postgreConn *sql.DB // Глобальное соединение sql

	DB_CONN_ERROR error
)

func init() {
	host = beego.AppConfig.String("db_host")
	port, _ = beego.AppConfig.Int("db_port")
	user = beego.AppConfig.String("db_user")
	password = beego.AppConfig.String("db_password")
	dbname = beego.AppConfig.String("db_name")

	max_conn = beego.AppConfig.DefaultInt("db_max_conn", 50)
	idle_conn = beego.AppConfig.DefaultInt("db_idle_conn", 10)

	DB_CONN_ERROR = errors.New(exceptions.DB_connection_error)
}

func GetPostgresSession() *sql.DB {
	if postgreConn == nil {
		psqlInfo := fmt.Sprintf(`host=%s port=%d user=%s password=%s dbname=%s sslmode=disable`,
			host, port, user, password, dbname)

		db, err := sql.Open("postgres", psqlInfo)
		if err != nil {
			return nil
		}

		db.SetConnMaxLifetime(30 * time.Minute)
		db.SetMaxOpenConns(max_conn)
		db.SetMaxIdleConns(idle_conn)

		//		err = db.Ping()
		//		if err != nil {
		//			return nil
		//		}

		postgreConn = db
	}

	return postgreConn
}
