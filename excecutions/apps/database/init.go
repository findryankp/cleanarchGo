package database

func InitDatabase() {
	mysqlConCreate()
	migrationCreate()
}
