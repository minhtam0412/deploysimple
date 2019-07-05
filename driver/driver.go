package driver

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
)

var (
	conn                 *sqlx.DB
	CLEARDB_DATABASE_URL string
)

func init() {
	CLEARDB_DATABASE_URL = GetConnectionStringFromEnviroment()
	db, err := sqlx.Connect("mysql", CLEARDB_DATABASE_URL)
	if err != nil {
		log.Fatalln("Error in Connect:", err)
		return
	}
	if err := db.Ping(); err != nil {
		log.Fatalln("Error in Ping:", err)
		return
	}
	conn.SetMaxOpenConns(100)
	conn.SetMaxIdleConns(10)
	log.Println("Connect success!")
}

func GetConnectionStringFromEnviroment() (connString string) {
	connString = os.Getenv("CLEARDB_DATABASE_URL");
	if len(connString) == 0 {
		log.Println("Not setting CLEARDB_DATABASE_URL. Setting to:", CLEARDB_DATABASE_URL)
		connString = "be598cab4597a2:2dbfb2f8@tcp(us-cdbr-iron-east-02.cleardb.net:3306)/heroku_87f7c0b2de20104"
	}
	return
}

func GetConn() *sqlx.DB {
	return conn
}

func SetConn(newConn *sqlx.DB) {
	conn = newConn
}
