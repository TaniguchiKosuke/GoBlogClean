package di

import "GoBlogClean/config"

func InjectDB() *config.DBHandler {
	dbHandler := config.NewDBHandler()
	return dbHandler
}
