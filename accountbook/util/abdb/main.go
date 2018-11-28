package main

import (
	"log"

	"gopkg.in/src-d/go-mysql-server.v0"
	"gopkg.in/src-d/go-mysql-server.v0/auth"
	"gopkg.in/src-d/go-mysql-server.v0/mem"
	"gopkg.in/src-d/go-mysql-server.v0/server"
	"gopkg.in/src-d/go-mysql-server.v0/sql"
)

func main() {
	engine := sqle.NewDefault()
	engine.AddDatabase(mem.NewDatabase("kakeibo"))
	engine.AddDatabase(sql.NewInformationSchemaDatabase(engine.Catalog))

	config := server.Config{
		Protocol: "tcp",
		Address:  "localhost:3306",
		Auth:     auth.NewNativeSingle("test-user", "test-pass", auth.AllPermissions),
	}

	s, err := server.NewDefaultServer(config, engine)
	if err != nil {
		log.Fatal(err)
	}

	s.Start()
}
