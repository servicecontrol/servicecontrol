// Package postgresql provides a wrapper around the sqlx package.
package postgresql

import (
	"crypto/tls"
	"fmt"
	"sync"
	"time"

	"gopkg.in/pg.v4"
)

// *****************************************************************************
// Thread-Safe Configuration
// *****************************************************************************

var (
	info       Info
	infoMutex  sync.RWMutex
	dbInstance pg.DB
)

// Info holds the details for the Postgresql connection.
type Info struct {
	Network string
	// TCP host:port or Unix socket depending on Network.
	Addr     string
	Username string
	Password string
	Database string

	// Whether to use secure TCP/IP connections (TLS).
	// TODO: deprecated in favor of TLSConfig
	SSL bool
	// TLS config for secure connections.
	TLSConfig *tls.Config

	// PostgreSQL run-time configuration parameters to be set on connection.
	Params map[string]interface{}

	// Maximum number of retries before giving up.
	// Default is to not retry failed queries.
	MaxRetries int

	// Dial timeout for establishing new connections.
	// Default is 5 seconds.
	DialTimeout time.Duration
	// Timeout for socket reads. If reached, commands will fail
	// with a timeout instead of blocking.
	ReadTimeout time.Duration
	// Timeout for socket writes. If reached, commands will fail
	// with a timeout instead of blocking.
	WriteTimeout time.Duration

	// Maximum number of socket connections.
	// Default is 20 connections.
	PoolSize int
	// Amount of time client waits for free connection if all
	// connections are busy before returning an error.
	// Default is 5 seconds.
	PoolTimeout time.Duration
	// Amount of time after which client closes idle connections.
	// Default is to not close idle connections.
	IdleTimeout time.Duration
	// Frequency of idle checks.
	// Default is 1 minute.
	IdleCheckFrequency time.Duration
}

// SetConfig stores the config.
func SetConfig(i Info) {
	infoMutex.Lock()
	info = i
	infoMutex.Unlock()
}

// Config returns the config.
func Config() Info {
	infoMutex.RLock()
	defer infoMutex.RUnlock()
	return info
}

// ResetConfig removes the config.
func ResetConfig() {
	infoMutex.Lock()
	info = Info{}
	infoMutex.Unlock()
}

// *****************************************************************************
// Database Handling
// *****************************************************************************

// Connect to the database.
func Connect() (db *pg.DB) {
	dbi := pg.Connect(&pg.Options{
		User:     info.Username,
		Database: info.Database,
	})
	err := createSchema(dbi)
	if err != nil {
		fmt.Println(err)
	}
	dbInstance = *dbi
	return dbi
}

//Instance returns the DB instance
func Instance() (db *pg.DB) {
	return &dbInstance
}

// Disconnect the database connection.
func Disconnect(db *pg.DB) error {
	return db.Close()
}

func createSchema(db *pg.DB) error {
	queries := []string{
		`CREATE TEMP TABLE projects (id serial, name text, description text, created timestamp)`,
	}
	for _, q := range queries {
		_, err := db.Exec(q)

		if err != nil {
			return err
		}
	}
	return nil
}
