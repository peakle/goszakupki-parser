package manager

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/peakle/goszakupki-parser/pkg/provider"
	sg "github.com/wakeapp/go-sql-generator"
)

type config struct {
	Host     string
	Username string
	Pass     string
	Port     string
	DBName   string
}

// SQLManager - manage connect to db
type SQLManager struct {
	conn *sql.DB
}

// UpsertLots - upsert new lots
func (m *SQLManager) UpsertLots(lots []provider.Lot) {
	//TODO
}

// GetLots - get lots by filters
func (m *SQLManager) GetLots(entry provider.EntryDto) ([]provider.ResponseLot, error) {
	//TODO

	return nil, nil
}

// InitManager - init connect to db
func InitManager() *SQLManager {
	m := &SQLManager{}

	m.open(&config{
		Host:     os.Getenv("MYSQL_HOST"),
		Username: os.Getenv("MYSQL_USER"),
		Pass:     os.Getenv("MYSQL_PASSWORD"),
		Port:     "3306",
		DBName:   os.Getenv("MYSQL_DATABASE"),
	})

	return m
}

// Close - close connect to db
func (m *SQLManager) Close() {
	m.conn.Close()
}

func (m *SQLManager) insert(dataInsert *sg.InsertData) int {
	if len(dataInsert.ValuesList) == 0 {
		return 0
	}

	sqlGenerator := sg.MysqlSqlGenerator{}

	query, args, err := sqlGenerator.GetInsertSql(*dataInsert)
	if err != nil {
		fmt.Printf("error occurred on generate insert sql: %s \n", err.Error())

		return 0
	}

	var stmt *sql.Stmt
	stmt, err = m.conn.Prepare(query)
	if err != nil {
		fmt.Printf("error occurred on prepare stmt: %s \n", err.Error())

		return 0
	}
	defer func() {
		_ = stmt.Close()
	}()

	var result sql.Result
	result, err = stmt.Exec(args...)
	if err != nil {
		fmt.Printf("error occurred on execute stmt: %s \n", err.Error())

		return 0
	}

	ra, _ := result.RowsAffected()

	return int(ra)
}

func (m *SQLManager) upsert(dataUpsert *sg.UpsertData) int {
	if len(dataUpsert.ValuesList) == 0 {
		return 0
	}

	sqlGenerator := sg.MysqlSqlGenerator{}

	query, args, err := sqlGenerator.GetUpsertSql(*dataUpsert)
	if err != nil {
		fmt.Printf("error occurred on Generate query: %v, %s \r\n", dataUpsert, err.Error())

		return 0
	}

	var stmt *sql.Stmt
	stmt, err = m.conn.Prepare(query)
	if err != nil {
		fmt.Printf("error occurred on Prepare query: %s, %s \r\n", query, err.Error())

		return 0
	}
	defer func() {
		_ = stmt.Close()
	}()

	var result sql.Result
	result, err = stmt.Exec(args...)
	if err != nil {
		fmt.Printf("error occurred on Exec query, args: %v, %s \r\n", args, err.Error())

		return 0
	}

	ra, _ := result.RowsAffected()

	return int(ra)
}

func (m *SQLManager) open(c *config) {
	var conn *sql.DB
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?collation=utf8_unicode_ci", c.Username, c.Pass, c.Host, c.Port, c.DBName)
	if conn, err = sql.Open("mysql", dsn); err != nil {
		fmt.Printf("error occurred on open connection to db: %s \n", err.Error())

		os.Exit(1)
	}

	m.conn = conn
}
