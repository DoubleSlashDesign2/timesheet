package models

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // needed becouse of gorm design
	// _ "github.com/jinzhu/gorm/dialects/sqlite"
)

// DB abstraction
type DB struct {
	*gorm.DB
}

// UpdatedValue used to pass updated ReportedRecord value type and relevant ID
type UpdatedValue struct {
	ID    int64
	Type  string
	Value string
}

// DateTime used for csv marshalling and unmarshalling
type DateTime struct {
	time.Time
}

// Date used for csv marshalling and unmarshalling
type Date struct {
	time.Time
}

// MarshalCSV Convert the internal date as CSV string
func (date *DateTime) MarshalCSV() (string, error) {
	return date.Time.Format("2006-01-02 15:04:05"), nil
}

// UnmarshalCSV Convert the CSV string as internal date
func (date *DateTime) UnmarshalCSV(csv string) (err error) {
	date.Time, err = time.Parse("2006-01-02 15:04:05", csv)
	return err
}

// MarshalCSV Convert the internal date as CSV string
func (date *Date) MarshalCSV() (string, error) {
	return date.Time.Format("2006-01-02"), nil
}

// UnmarshalCSV Convert the CSV string as internal date
func (date *Date) UnmarshalCSV(csv string) (err error) {
	date.Time, err = time.Parse("2006-01-02", csv)
	return err
}

// NewPostgresDB - postgres database
func NewPostgresDB(dataSourceName string) *DB {

	db, err := gorm.Open("postgres", dataSourceName)
	if err != nil {
		panic(err)
	}

	if err = db.DB().Ping(); err != nil {
		panic(err)
	}

	// db.LogMode(true)

	return &DB{db}
}

// NewSqliteDB - sqlite database
// func NewSqliteDB(dataSourceName string) *DB {

// 	db, err := gorm.Open("sqlite3", dataSourceName)
// 	if err != nil {
// 		panic(err)
// 	}

// 	if err = db.DB().Ping(); err != nil {
// 		panic(err)
// 	}

// 	// db.LogMode(true)

// 	return &DB{db}
//   }
