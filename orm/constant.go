package orm

var (
	PostgresCheckDatabaseExist = "SELECT 1 FROM pg_database WHERE datname = ?"
	PostgresCreateDatabase     = "CREATE DATABASE "
	PostgresDropDatabase       = "DROP DATABASE IF EXISTS "
)
