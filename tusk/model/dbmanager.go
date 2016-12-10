package model

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type DatabaseManager struct {
	User    string
	Pass    string
	Socket  string
	DBType  string
	DBName  string
	db      *sql.DB
	funcMap map[string]func(Tester) []*TestResult
}

func NewDatabaseManager(user, pass, sock, dbtype, dbname string) *DatabaseManager {
	dbm := new(DatabaseManager)
	dbm.User = user
	dbm.Pass = pass
	dbm.Socket = sock
	dbm.DBType = dbtype
	dbm.DBName = dbname
	dbm.funcMap = dbm.initMap()
	return dbm
}

func (dm *DatabaseManager) Init() {
	dbOpen := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		dm.User,
		dm.Pass,
		dm.Socket,
		dm.DBName)
	db, err := sql.Open(dm.DBType, dbOpen)
	HandleError(err)
	dm.db = db
	err = dm.db.Ping()
	HandleError(err)
}

func (dm *DatabaseManager) GetQueryFunction(query string) func(Tester) []*TestResult {
	return dm.funcMap[query]
}

func (dm *DatabaseManager) QueryAll(tr Tester) []*TestResult {
	query := fmt.Sprintf("SELECT * FROM %s",
		tr.GetTest().Table)
	return dm.execQuery(query)
}

func (dm *DatabaseManager) QueryResult(tr Tester) []*TestResult {
	query := fmt.Sprintf("SELECT * FROM %s WHERE result = \"%d\"",
		tr.GetTest().Table,
		tr.GetTest().Result)
	return dm.execQuery(query)
}

func (dm *DatabaseManager) QueryDate(tr Tester) []*TestResult {
	query := fmt.Sprintf("SELECT * FROM %s WHERE date = \"%s\"",
		tr.GetTest().Table,
		tr.GetTest().Date)
	return dm.execQuery(query)
}

func (dm *DatabaseManager) QueryDateSince(tr Tester) []*TestResult {
	query := fmt.Sprintf("SELECT * FROM %s WHERE date >= '%s'",
		tr.GetTest().Table,
		tr.GetTest().Date)
	return dm.execQuery(query)
}

func (dm *DatabaseManager) QueryDateSinceWithResult(tr Tester) []*TestResult {
	query := fmt.Sprintf("SELECT * FROM %s WHERE date >= '%s' AND result = \"%d\"",
		tr.GetTest().Table,
		tr.GetTest().Date,
		tr.GetTest().Result)
	return dm.execQuery(query)
}

func (dm *DatabaseManager) QueryPattern(tr Tester) []*TestResult {
	query := fmt.Sprintf("SELECT * FROM %s WHERE testID LIKE '%%%s%%'",
		tr.GetTest().Table,
		tr.GetTest().Pattern)
	return dm.execQuery(query)
}

func (dm *DatabaseManager) DeleteEntriesSince(tr Tester) []*TestResult {
	stmt := fmt.Sprintf("DELETE FROM %s WHERE date <= '%s'",
		tr.GetTest().Table,
		tr.GetTest().Date)
	dm.execStatement(stmt)
	return nil
}

func (dm *DatabaseManager) Insert(tr Tester) []*TestResult {
	stmt := fmt.Sprintf("INSERT INTO %s (testID, date, result) VALUES ('%s','%s','%d')",
		tr.GetTest().Table,
		tr.GetTest().Pattern,
		tr.GetTest().Date,
		tr.GetTest().Result)
	dm.execStatement(stmt)
	return nil
}

func (dm *DatabaseManager) Delete(tr Tester) []*TestResult {
	stmt := fmt.Sprintf("DELETE FROM %s WHERE testID='%s' and date='%s' and result='%d'",
		tr.GetTest().Table,
		tr.GetTest().Pattern,
		tr.GetTest().Date,
		tr.GetTest().Result)
	dm.execStatement(stmt)
	return nil
}

func (dm *DatabaseManager) initMap() map[string]func(Tester) []*TestResult {
	funcMap := make(map[string]func(Tester) []*TestResult)
	funcMap["QueryAll"] = dm.QueryAll
	funcMap["QueryResult"] = dm.QueryResult
	funcMap["QueryDate"] = dm.QueryDate
	funcMap["QueryDateSince"] = dm.QueryDateSince
	funcMap["QueryDateSinceWithResult"] = dm.QueryDateSinceWithResult
	funcMap["QueryPattern"] = dm.QueryPattern
	funcMap["Delete"] = dm.Delete
	funcMap["DeleteEntriesSince"] = dm.DeleteEntriesSince
	funcMap["Insert"] = dm.Insert
	dm.funcMap = funcMap
	return funcMap
}

func (dm *DatabaseManager) execStatement(stmt string) {
	_, err := dm.db.Exec(stmt)
	HandleError(err)
}

func (dm *DatabaseManager) execQuery(query string) []*TestResult {
	rows, err := dm.db.Query(query)
	HandleError(err)
	defer rows.Close()
	table := TableFromQuery(query)
	return unwrapQuery(rows, table)
}

func unwrapQuery(rows *sql.Rows, table string) []*TestResult {
	testResults := make([]*TestResult, 0)
	for rows.Next() {
		tr := new(TestResult)
		err := rows.Scan(
			&tr.GetTest().Pattern,
			&tr.GetTest().Date,
			&tr.GetTest().Result,
			&tr.RunID)
		tr.Table = table
		HandleError(err)
		testResults = append(testResults, tr)
	}
	return testResults
}
