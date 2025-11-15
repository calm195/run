package orm

import (
	"run/global"
	"run/models"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func Gorm() *gorm.DB {
	p := global.Config.Pgsql

	if p.Dbname == "" {
		global.Log.Fatal("Database name is empty", zap.String("config", p.Dsn()))
		panic("Database name is empty")
	}
	db, err := gorm.Open(postgres.Open(p.DefaultDsn()))
	if err != nil {
		global.Log.Fatal("Failed to connect to database", zap.String("config", p.DefaultDsn()), zap.Error(err))
		panic(err.Error())
	}

	var count int
	db.Raw(PostgresCheckDatabaseExist, p.Dbname).Scan(&count)
	if count == 0 {
		global.Log.Info("Database does not exist", zap.String("database", p.Dbname))
		createTable := PostgresCreateDatabase + p.Dbname
		global.Log.Info("begin create database", zap.String("config", createTable))
		err = db.Exec(createTable).Error
		if err != nil {
			global.Log.Fatal("Failed to create database", zap.String("database", p.Dbname), zap.Error(err))
			panic(err.Error())
		}
		global.Log.Info("create database successfully", zap.String("database", p.Dbname))
	}

	closeDb, _ := db.DB()
	err = closeDb.Close()
	if err != nil {
		global.Log.Fatal("Failed to close default database", zap.String("config", p.DefaultDsn()), zap.Error(err))
		panic(err.Error())
	}

	pgConfig := postgres.Config{
		DSN:                  p.Dsn(),
		PreferSimpleProtocol: false,
	}
	gormConfig := &gorm.Config{
		Logger: logger.New(NewWriter(p), logger.Config{
			SlowThreshold: 200 * time.Millisecond,
			LogLevel:      p.LogLevel(),
			Colorful:      true,
		}),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	}
	global.Log.Info("connect to database", zap.String("config", p.Dsn()))
	if db, err = gorm.Open(postgres.New(pgConfig), gormConfig); err != nil {
		global.Log.Fatal("Failed to connect to database", zap.String("config", p.Dsn()), zap.Error(err))
		panic(err.Error())
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(p.MaxIdleConnections)
		sqlDB.SetMaxOpenConns(p.MaxOpenConnections)
		global.Log.Info("connect to database successfully", zap.String("config", p.Dsn()))
		return db
	}
}

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
	)
	if err != nil {
		global.Log.Fatal("failed to migrate table", zap.Error(err))
		panic(err.Error())
	}
	global.Log.Info("register table successfully")
}
