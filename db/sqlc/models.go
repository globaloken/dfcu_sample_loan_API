// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"
)

type Currency string

const (
	CurrencyUGX Currency = "UGX"
)

func (e *Currency) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Currency(s)
	case string:
		*e = Currency(s)
	default:
		return fmt.Errorf("unsupported scan type for Currency: %T", src)
	}
	return nil
}

type NullCurrency struct {
	Currency Currency
	Valid    bool // Valid is true if Currency is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullCurrency) Scan(value interface{}) error {
	if value == nil {
		ns.Currency, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Currency.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullCurrency) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return ns.Currency, nil
}

type LogType string

const (
	LogTypeREQUEST          LogType = "REQUEST"
	LogTypeFAILEDVALIDATION LogType = "FAILED_VALIDATION"
	LogTypePOSITIVEREQUEST  LogType = "POSITIVE_REQUEST"
	LogTypeNEGATIVEREQUEST  LogType = "NEGATIVE_REQUEST"
)

func (e *LogType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = LogType(s)
	case string:
		*e = LogType(s)
	default:
		return fmt.Errorf("unsupported scan type for LogType: %T", src)
	}
	return nil
}

type NullLogType struct {
	LogType LogType
	Valid   bool // Valid is true if LogType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullLogType) Scan(value interface{}) error {
	if value == nil {
		ns.LogType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.LogType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullLogType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return ns.LogType, nil
}

type UserType string

const (
	UserTypeADMIN  UserType = "ADMIN"
	UserTypeCLIENT UserType = "CLIENT"
)

func (e *UserType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = UserType(s)
	case string:
		*e = UserType(s)
	default:
		return fmt.Errorf("unsupported scan type for UserType: %T", src)
	}
	return nil
}

type NullUserType struct {
	UserType UserType
	Valid    bool // Valid is true if UserType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullUserType) Scan(value interface{}) error {
	if value == nil {
		ns.UserType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.UserType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullUserType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return ns.UserType, nil
}

type Loan struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	// must be positive
	Amount    int64     `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

type Log struct {
	ID        int64          `json:"id"`
	Username  sql.NullString `json:"username"`
	Type      LogType        `json:"type"`
	CreatedAt time.Time      `json:"created_at"`
}

type User struct {
	Username          string       `json:"username"`
	HashedPassword    string       `json:"hashed_password"`
	FullName          string       `json:"full_name"`
	Type              UserType     `json:"type"`
	Email             string       `json:"email"`
	AccountNo         string       `json:"account_no"`
	Balance           int64        `json:"balance"`
	Currency          Currency     `json:"currency"`
	PasswordChangedAt sql.NullTime `json:"password_changed_at"`
	CreatedAt         time.Time    `json:"created_at"`
}
