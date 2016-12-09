package pgsql

import (
	"database/sql"
	"database/sql/driver"
	"time"
)

// String type for work with string in DB
type nullString struct {
	sql.NullString
}

// Function return new nullString, if pointer is not null
func newNullStringByPointer(s *string) nullString {
	if s != nil {
		return newNullString(*s)
	}
	return newNullString("")
}
func newNullString(s string) nullString {
	return nullString{
		NullString: sql.NullString{
			String: s,
			Valid:  s != "",
		},
	}
}

type nullInt64 struct {
	sql.NullInt64
}

func newNullInt64ByUint(i uint) nullInt64 {
	return nullInt64{
		NullInt64: sql.NullInt64{
			Int64: int64(i),
			Valid: i >= 0,
		},
	}
}

func (n nullInt64) uint() uint {
	return uint(n.Int64)
}

type nullTime struct {
	Time  time.Time
	Valid bool
}

func newNullTime(t time.Time) nullTime {
	return nullTime{
		Time:  t,
		Valid: true,
	}
}

func (nt *nullTime) Scan(value interface{}) error {
	nt.Time, nt.Valid = value.(time.Time)
	return nil
}

func (nt nullTime) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}
	return nt.Time, nil
}
