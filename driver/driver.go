package driver

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
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
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)
	log.Println("Connect success!")
	conn = db
}

func GetConnectionStringFromEnviroment() (connString string) {
	//connString = os.Getenv("CLEARDB_DATABASE_URL")
	connString = ""
	if len(connString) == 0 {
		connString = "be598cab4597a2:2dbfb2f8@tcp(us-cdbr-iron-east-02.cleardb.net:3306)/heroku_87f7c0b2de20104?parseTime=True&loc=Local"
		//connString = "root:123456@tcp(localhost:3306)/simpledeploy?parseTime=true"
		//connString = "root:123456@tcp(localhost:3306)/simpledeploy?parseTime=True&loc=Local"
		log.Println("Not setting CLEARDB_DATABASE_URL. Setting to:", connString)
	}
	return
}

func GetConn() *sqlx.DB {
	return conn
}

func SetConn(newConn *sqlx.DB) {
	conn = newConn
}
