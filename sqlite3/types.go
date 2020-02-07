package sqlite3

import (
	"sync"
	"time"

	"github.com/codingbeard/cbsession"
)

// Config session sqlite3 configuration
type Config struct {

	// sqlite3 db file path
	DBPath string

	// session table name
	TableName string

	// sqlite3 max free idle
	SetMaxIdleConn int

	// sqlite3 max open idle
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
