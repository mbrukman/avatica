/*
Package avatica provides an Apache Phoenix Query Server/Avatica driver for Go's database/sql package.

Quickstart

Import the database/sql package along with the avatica driver.

	import "database/sql"
	import _ "github.com/Boostport/avatica"

	db, err := sql.Open("avatica", "http://phoenix-query-server:8765")

See https://github.com/Boostport/avatica#usage for more details
*/
package avatica

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"github.com/Boostport/avatica/message"
	"github.com/satori/go.uuid"
	"golang.org/x/net/context"
)

// Driver is exported to allow it to be used directly.
type Driver struct{}

// Open a Connection to the server.
// See https://github.com/Boostport/avatica#dsn for more information
// on how the DSN is formatted.
func (a *Driver) Open(dsn string) (driver.Conn, error) {

	config, err := ParseDSN(dsn)

	if err != nil {
		return nil, fmt.Errorf("Unable to open connection: %s", err)
	}

	httpClient := NewHTTPClient(config.endpoint)
	connectionId := uuid.NewV4().String()

	// Open a connection to the server
	_, err = httpClient.post(context.Background(), &message.OpenConnectionRequest{
		ConnectionId: connectionId,
		Info: map[string]string{
			"AutoCommit":  "true",
			"Consistency": "8",
		},
	})

	if err != nil {
		return nil, err
	}

	conn := &conn{
		connectionId: connectionId,
		httpClient:   httpClient,
		config:       config,
	}

	return conn, nil
}

func init() {
	sql.Register("avatica", &Driver{})
}
