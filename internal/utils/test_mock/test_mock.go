package interfaces

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/nichojovi/stbit-test/internal/utils/database"
)

type DatabaseMock struct {
	master *DatabaseMockMaster
	slave  *sqlx.DB
	*DatabaseMockTx
}

type DatabaseMockMaster struct {
	*sqlx.DB
}

type DatabaseMockTx struct {
	*sqlx.Tx
	id int64
}

func (db *DatabaseMockTx) Ping() error {
	return nil
}

func (db *DatabaseMockTx) PingContext(ctx context.Context) error {
	return nil
}

func (db *DatabaseMockTx) ID() int64 {
	return db.id
}
func (db *DatabaseMockTx) Commit() error {
	return db.Tx.Commit()
}

func (db *DatabaseMockTx) Rollback() error {
	return db.Tx.Rollback()
}

type TestDBStore struct {
	Store *database.Store
	DBSet *MockDBSet
}

type MockDB struct {
	DB   *sql.DB
	Mock sqlmock.Sqlmock
}

type MockDBSet struct {
	Master *MockDB
	Slave  *MockDB
}

func (m *MockDBSet) Close() {
	m.Master.DB.Close()
	m.Slave.DB.Close()
}

func NewTestDBStore() *TestDBStore {
	dbSet := MockDBSet{
		Master: newDBMock(),
		Slave:  newDBMock(),
	}
	return &TestDBStore{
		Store: &database.Store{
			Master: sqlx.NewDb(dbSet.Master.DB, "sqlmock"),
			Slave:  sqlx.NewDb(dbSet.Slave.DB, "sqlmock"),
		},
		DBSet: &dbSet,
	}
}

func newDBMock() *MockDB {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal("failed to create db mock caused:", err)
	}
	return &MockDB{
		DB:   db,
		Mock: mock,
	}
}

//NormalizeRegexpQuery, nomralize query and escape regexp character
func NormalizeRegexpQuery(q string) string {
	return fmt.Sprintf("^%s$", regexp.QuoteMeta(RemoveSpace(q)))
}

//RemoveSpace, normalize query by removing space and tab character
func RemoveSpace(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

func GetColumnsRowFromQuery(query string) []string {
	re, _ := regexp.Compile("(?i)select (.*?) from")
	columnClause := re.FindStringSubmatch(query)[1]
	re, _ = regexp.Compile("(?i)case.*?end")
	columnClause = re.ReplaceAllString(columnClause, "")
	columns := strings.Split(columnClause, ",")
	for i := range columns {
		col := columns[i]
		if idx := strings.Index(strings.ToLower(col), " as "); idx > 0 {
			col = col[idx+3 : len(col)]
		}
		if strings.Contains(col, ".") {
			col = strings.Split(col, ".")[1]
		}
		col = strings.TrimSpace(col)
		col = strings.Trim(col, "'")
		col = strings.Trim(col, "\"")
		columns[i] = col
	}
	return columns
}

func MockRowsFromQuery(query string) *sqlmock.Rows {
	return sqlmock.NewRows(GetColumnsRowFromQuery(query))
}
