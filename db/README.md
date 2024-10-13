# db

This directory stores all migrations, data access constructs and queries.

- `/db/migration` contains the migration files for running with [golang-migrate](https://github.com/golang-migrate/migrate).

  - Note that all migration files must be prefixed with `xxx_` in order for [sqlc](https://docs.sqlc.dev/en/latest/howto/ddl.html#golang-migrate) to parse correctly.

- `/db/sqlc` contains queries which `sqlc` will use to generate access methods and models.

- `/db/queries` contains sqlc generated code
