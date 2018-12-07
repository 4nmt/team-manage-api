package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-gorp/gorp"
	_ "github.com/lib/pq" //import postgres
)

//DB ...
type DB struct {
	*sql.DB
}

const (
	//DbUser ...
	// DbUser = "postgres"
	DbUser = "oypwfmnxlfdtrw"

	//DbPassword ...
	// DbPassword = "example"
	DbPassword = "707a89591c093c910f1873f6a3258dd466b646573537a4a0b14bfc19035dc452"

	//DbName ...
	// DbName = "team_manage_app"
	DbName = "d7a8rbk65f03hv"

	//DbHost ...
	DbHost = "ec2-174-129-41-12.compute-1.amazonaws.com"

	//DbPort ...
	// DbName = "team_manage_app"
	DbPort = "5432"
)

var db *gorp.DbMap

//Init ...
func Init() {

	// dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
	// 	DbUser, DbPassword, DbName)

		dbinfo := fmt.Sprintf("user=%s dbname=%s sslmode=%s password=%s host=%s port=%s",
	DbUser,
	DbName,
	"require",
	DbPassword,
	DbHost,
	DbPort,
	

	var err error
	db, err = ConnectDB(dbinfo)
	if err != nil {
		log.Fatal(err)
	}

}

//ConnectDB ...
func ConnectDB(dataSourceName string) (*gorp.DbMap, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	//dbmap.TraceOn("[gorp]", log.New(os.Stdout, "golang-gin:", log.Lmicroseconds)) //Trace database requests
	return dbmap, nil
}

//GetDB ...
func GetDB() *gorp.DbMap {
	return db
}
