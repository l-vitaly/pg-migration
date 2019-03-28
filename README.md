# Go postgres migration app.

## Instalation

```bash
go get -u github.com/l-vitaly/pg-migration/cmd/migration
```

## Usage

```bash
$ export MIGRATION_DB_CONN=postgres://root:root@localhost:5432/db?sslmode=disable
$ migration init
$ migration up
```

## Supported commands:

```
init                   - initalization migrations.
up                     - runs all available migrations.
up [target]            - runs available migrations up to the target one.
down                   - reverts last migration.
reset                  - reverts all migrations.
version                - prints current db version.
set_version [version]  - sets db version without running migrations.
```
