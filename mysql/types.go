package mysql

import (
	"sync"
	"time"

	"github.com/codingbeard/cbsession"
)

// Config session mysql configuration
type Config struct {

	// mysql server host
	Host string

	// mysql server port
	Port int

	// mysql username
	Username string

	// mysql password
	Password string

	// mysql conn charset
	Charset string

	// mysql Collate
	Collate string

	// database name
	Database string

	// session table name
	TableName string

	// mysql conn timeout(s)
	Timeout int

	// mysql read timeout(s)
	ReadTimeout int

	// mysql write timeout(s)
	WriteTimeout int

	// mysql max free idle
	SetMaxIdleConn int

	// mysql max open idle
	SetMaxOpenConn int

	// session value serialize func
	SerializeFunc func(src cbsession.Dict) ([]byte, error)

	// session value unSerialize func
	UnSerializeFunc func(dst *cbsession.Dict, src []byte) error
}

// Provider provider struct
type Provider struct {
	config     *Config
	db         *Dao
	expiration time.Duration

	storePool sync.Pool
}

// Store store struct
type Store struct {
	cbsession.Store
}

// Dao database access object
type Dao struct {
	cbsession.Dao

	tableName string

	sqlGetSessionBySessionID string
	sqlCountSessions         string
	sqlUpdateBySessionID     string
	sqlDeleteBySessionID     string
	sqlDeleteExpiredSessions string
	sqlInsert                string
	sqlRegenerate            string
}

// DBRow database row definition
type DBRow struct {
	sessionID  string
	contents   string
	lastActive int64
	expiration time.Duration
}
