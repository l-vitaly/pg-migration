package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/go-pg/migrations"
	"github.com/go-pg/pg"
	"github.com/l-vitaly/pg-migration/pkg/config"
)

const usageText = `This program runs command on the db. Supported commands are:
  - init                   - initalization migrations.
  - up                     - runs all available migrations.
  - up [target]            - runs available migrations up to the target one.
  - down                   - reverts last migration.
  - reset                  - reverts all migrations.
  - version                - prints current db version.
  - set_version [version]  - sets db version without running migrations.
Usage:
  migration <command> [args]
`

func main() {
	flag.Usage = usage
	flag.Parse()

	cfg, err := config.Parse()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	opts, err := pg.ParseURL(cfg.DbConn.String())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	db := pg.Connect(opts)
	defer db.Close()

	oldVersion, newVersion, err := migrations.Run(db, flag.Args()...)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if newVersion != oldVersion {
		fmt.Printf("migrated from version %d to %d\n", oldVersion, newVersion)
	} else {
		fmt.Printf("version is %d\n", oldVersion)
	}
}

func usage() {
	fmt.Print(usageText)
	flag.PrintDefaults()
	os.Exit(2)
}
