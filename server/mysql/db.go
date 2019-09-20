package mysql

import (
	"database/sql"
	"fabric-orderer-seek/server/helpers"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type DBMgr struct {
	db        *sql.DB
	stmtTx    *sql.Stmt
	stmtBlock *sql.Stmt
}

var (
	blockSQL = "INSERT INTO block VALUES(?,?,?,?,?)"
	txSQL    = "INSERT INTO transaction VALUES(?,?,?,?,?,?,?)"
	dbInfo   = helpers.GetAppConf().Conf.DB
)

var dbMgr *DBMgr

func init() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbInfo.User, dbInfo.Password, dbInfo.Host, dbInfo.Port, dbInfo.Name))
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	db.SetMaxOpenConns(dbInfo.MaxOpenConns)
	db.SetMaxIdleConns(dbInfo.MaxIdleConns)
	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	stmtBlock, err := db.Prepare(blockSQL) // ? = placeholder
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	stmtTx, err := db.Prepare(txSQL) // ? = placeholder
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	dbMgr = &DBMgr{db, stmtTx, stmtBlock}
}

func CloseDB() {
	if err := dbMgr.db.Close(); err != nil {
		panic(err)
	}
	if err := dbMgr.stmtTx.Close(); err != nil {
		panic(err)
	}
	if err := dbMgr.stmtBlock.Close(); err != nil {
		panic(err)
	}
}

func GetDB() *sql.DB {
	return dbMgr.db
}

func GetStmtTx() *sql.Stmt {
	return dbMgr.stmtTx
}

func GetStmtBlock() *sql.Stmt {
	return dbMgr.stmtBlock
}
