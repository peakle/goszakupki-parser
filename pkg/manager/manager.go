package manager

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

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

// InsertPurchase - insert new lots/trends/purchase
func (m *SQLManager) InsertPurchase(lots []*provider.Purchase) {
	const TableName = "Purchase"

	dataInsert := &sg.InsertData{
		TableName: TableName,
		Fields: []string{
			"id",
			"fz",
			"customer",
			"customer_link",
			"customer_inn",
			"customer_region",
			"bidding_region",
			"customer_activity_field",
			"bidding_volume",
			"bidding_count",
			"purchase_target",
			"registry_bidding_number",
			"contract_price",
			"participation_security_amount",
			"execution_security_amount",
			"published_at",
			"requisition_deadline_at",
			"contract_start_at",
			"contract_end_at",
			"playground",
			"purchase_link",
		},
		IsIgnore: true,
	}

	for _, lot := range lots {
		dataInsert.Add([]string{
			lot.ID,
			lot.Fz,
			lot.Customer,
			lot.CustomerLink,
			lot.CustomerInn,
			lot.CustomerRegion,
			lot.BiddingRegion,
			lot.CustomerActivityField,
			lot.BiddingVolume,
			lot.BiddingCount,
			lot.PurchaseTarget,
			lot.RegistryBiddingNumber,
			lot.ContractPrice,
			lot.ParticipationSecurityAmount,
			lot.ExecutionSecurityAmount,
			lot.PublishedAt,
			lot.RequisitionDeadlineAt,
			lot.ContractStartAt,
			lot.ContractEndAt,
			lot.Playground,
			lot.PurchaseLink,
		})
	}

	m.insert(dataInsert)
}

// GetLots - get lots by filters
func (m *SQLManager) GetLots(entry provider.EntryDto) ([]provider.Purchase, error) {
	whereCond := parseParamsPurchase(entry)

	query := fmt.Sprintf(`
		SELECT
			id,
			fz,
			customer,
			customer_link,
			customer_inn,
			customer_region,
			bidding_region,
			customer_activity_field,
			bidding_volume,
			bidding_count,
			purchase_target,
			registry_bidding_number,
			contract_price,
			participation_security_amount,
			execution_security_amount,
			published_at,
			requisition_deadline_at,
			contract_start_at,
			contract_end_at,
			playground,
			purchase_link
		FROM Purchase
		WHERE 1
			%s
	`, whereCond)

	rows, err := m.conn.Query(query)
	if err != nil {
		if err == sql.ErrNoRows {
			return make([]provider.Purchase, 0, 0), nil
		}

		log.Printf("on query GetLots: %s, query: %s \n", err.Error(), query)

		return make([]provider.Purchase, 0, 0), err
	}

	result := make([]provider.Purchase, 0, 10)
	var row provider.Purchase

	for rows.Next() {
		err := rows.Scan(
			&row.ID,
			&row.Fz,
			&row.Customer,
			&row.CustomerLink,
			&row.CustomerInn,
			&row.CustomerRegion,
			&row.BiddingRegion,
			&row.CustomerActivityField,
			&row.BiddingVolume,
			&row.BiddingCount,
			&row.PurchaseTarget,
			&row.RegistryBiddingNumber,
			&row.ContractPrice,
			&row.ParticipationSecurityAmount,
			&row.ExecutionSecurityAmount,
			&row.PublishedAt,
			&row.RequisitionDeadlineAt,
			&row.ContractStartAt,
			&row.ContractEndAt,
			&row.Playground,
			&row.PurchaseLink,
		)

		if strings.HasPrefix(row.PurchaseLink, "https://") == false {
			row.PurchaseLink = fmt.Sprintf("https://zakupki.gov.ru/epz/order/notice%s", row.PurchaseLink)
		}

		if err != nil {
			log.Println("on GetLots: on scan: " + err.Error())
			return make([]provider.Purchase, 0, 0), err
		}

		result = append(result, row)
	}

	return result, nil
}

//rebuild to queryBuilder
// parseParamsPurchase - filters data
func parseParamsPurchase(entry provider.EntryDto) string {
	const tableName = "Purchase"
	var cond string

	if date, ok := entry["to_date"]; ok && date != "" {
		cond += fmt.Sprintf(" AND %s.published_at <= '%s' ", tableName, date)
	}

	if date, ok := entry["from_date"]; ok && date != "" {
		cond += fmt.Sprintf(" AND %s.published_at >= '%s' ", tableName, date)
	}

	if customer, ok := entry["customer"]; ok && customer != "" {
		cond += fmt.Sprintf(" AND %s.customer_inn = '%s' ", tableName, customer)
	}

	if region, ok := entry["region"]; ok && region != "" {
		cond += fmt.Sprintf(" AND %s.bidding_region = '%s' ", tableName, region)
	}

	if priceFrom, ok := entry["price_from"]; ok && priceFrom != "" {
		cond += fmt.Sprintf(" AND %s.bidding_volume >= '%s' ", tableName, priceFrom)
	}

	if priceTo, ok := entry["price_to"]; ok && priceTo != "" {
		cond += fmt.Sprintf(" AND %s.bidding_volume <= '%s' ", tableName, priceTo)
	}

	if grntShare, ok := entry["grnt_share_from"]; ok && grntShare != "" { // гарантия выполнения контракта более чем (>=)
		cond += fmt.Sprintf(" AND %s.execution_security_amount >= '%s' ", tableName, grntShare)
	}

	if grntShare, ok := entry["grnt_share_to"]; ok && grntShare != "" { // гарантия выполнения контракта менее чем (>=)
		cond += fmt.Sprintf(" AND %s.execution_security_amount <= '%s' ", tableName, grntShare)
	}

	if cond == "" {
		cond = " ORDER BY published_at DESC LIMIT 100 "
	}

	return cond
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
