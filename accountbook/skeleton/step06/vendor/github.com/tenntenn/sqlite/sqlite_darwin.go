// +build darwin,!windows,!linux

package sqlite

import _ "github.com/mattn/go-sqlite3"

const DriverName = "sqlite3"
