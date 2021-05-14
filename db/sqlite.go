package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
    "math/rand"
    "time"
)

type FastReport struct {
	Id          int
	ServiceName string
	Ok          int
	Fail        int
}

type DailyLog struct {
    Id          int
    ServiceName string
    LogDate     string
    Ok          int
    Fail        int
}

var qTableFastReport = "CREATE TABLE IF NOT EXISTS fast_report (id INTEGER PRIMARY KEY AUTOINCREMENT, service_name TEXT, ok INTEGER, fail INTEGER);"

var qTableDailyLog = "CREATE TABLE IF NOT EXISTS daily_log (id INTEGER PRIMARY KEY AUTOINCREMENT, service_name TEXT, log_date TEXT, ok INTEGER, fail INTEGER);"

var fReportTbl = "fast_report"
var dLogTbl = "daily_log"

var db, dberr = sql.Open("sqlite3", "xcheck.db")

func execute(qry string) {
	_, err := db.Exec(qry)
	if err != nil {
		log.Panic(err)
	}
}

func InitSqlite() {
	if dberr != nil {
		log.Panic(dberr)
	}

	execute(qTableFastReport)
	execute(qTableDailyLog)
}

func InsertLog(serviceName string, ok bool) {
    currentTime := time.Now()
    today := currentTime.Format("2006-01-02")
	var fReports, _ = db.Query("select * from fast_report;")
    for fReports.Next() {
        var fastReport FastReport
        fReports.Scan(&fastReport.Id,
            &fastReport.ServiceName,
            &fastReport.Ok,
            &fastReport.Fail)
        fmt.Println(fastReport)
    }


    var dLogs, _ = db.Query("select * from daily_log;")
    for dLogs.Next() {
        var dLog DailyLog
        dLogs.Scan(&dLog.Id,
            &dLog.ServiceName,
            &dLog.LogDate,
            &dLog.Ok,
            &dLog.Fail)
        fmt.Println(dLog)
    }

	qry := fmt.Sprintf("INSERT INTO %s (service_name,log_date,ok,fail) VALUES ('%s','%s',%d,%d);",
		dLogTbl,
		serviceName,
		today,
        rand.Intn(100),
        rand.Intn(100))
	execute(qry)

}
