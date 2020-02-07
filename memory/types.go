package memory

import (
	"sync"
	"time"

	"github.com/codingbeard/cbsession"
)

// Config session memory configuration
type Config struct{}

// Provider provider struct
type Provider struct {
	config     *Config
	memoryDB   *cbsession.Dict
	expiration time.Duration

	storePool sync.Pool

	lock sync.RWMutex
}

// Store memory store
type Store struct {
	cbsession.Store

	lastActiveTime int64
	lock           sync.RWMutex
}
