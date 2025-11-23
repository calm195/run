package init

import (
	"run/global"
	"run/models"

	"go.uber.org/zap"
)

func RegisterTables() {
	if global.Db == nil {
		global.Log.Fatal("not connect database")
		return
	}

	db := global.Db
	err := db.AutoMigrate(
		&models.Game{},
		&models.Record{},
		&models.GameRecord{},
		&models.Event{},
		&models.Standard{},
	)
	if err != nil {
		global.Log.Fatal("failed to migrate table", zap.Error(err))
		panic(err.Error())
	}
	global.Log.Info("register table successfully")
}
