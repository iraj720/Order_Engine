package confs

import (
	"database/sql"
	"os"

	"github.com/uptrace/bun/extra/bundebug"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func SetupPgConnection() (*bun.DB, *bun.DB) {
	rdbUser := os.Getenv("RDB_USER")
	rdbPass := os.Getenv("RDB_PASS")
	rdbName := os.Getenv("RDB_NAME")
	rdbHost := os.Getenv("RDB_HOST")
	rdbPort := os.Getenv("RDB_PORT")

	wdbUser := os.Getenv("WDB_USER")
	wdbPass := os.Getenv("WDB_PASS")
	wdbName := os.Getenv("WDB_NAME")
	wdbHost := os.Getenv("WDB_HOST")
	wdbPort := os.Getenv("WDB_PORT")

	rpgconn := pgdriver.NewConnector(
		pgdriver.WithNetwork("tcp"),
		pgdriver.WithAddr(rdbHost+":"+rdbPort),
		pgdriver.WithUser(rdbUser),
		pgdriver.WithPassword(rdbPass),
		pgdriver.WithDatabase(rdbName),
		pgdriver.WithInsecure(true))

	wpgconn := pgdriver.NewConnector(
		pgdriver.WithNetwork("tcp"),
		pgdriver.WithAddr(wdbHost+":"+wdbPort),
		pgdriver.WithUser(wdbUser),
		pgdriver.WithPassword(wdbPass),
		pgdriver.WithDatabase(wdbName),
		pgdriver.WithInsecure(true))

	//pgdriver.WithApplicationName("rabex"))

	rsqldb := sql.OpenDB(rpgconn)
	wsqldb := sql.OpenDB(wpgconn)

	rdb := bun.NewDB(rsqldb, pgdialect.New())
	wdb := bun.NewDB(wsqldb, pgdialect.New())
	rdb.SetMaxIdleConns(0)
	rdb.SetMaxOpenConns(1)
	wdb.SetMaxIdleConns(0)
	wdb.SetMaxOpenConns(1)
	rdb.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	wdb.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

	if err := rdb.Ping(); err != nil {
		panic("read connection could not connect to database: " + err.Error())
	}

	if err := wdb.Ping(); err != nil {
		panic("write connection could not connect to database: " + err.Error())
	}

	return rdb, wdb
}

func ClosePgConnection(db *bun.DB) error {
	return db.Close()
}
