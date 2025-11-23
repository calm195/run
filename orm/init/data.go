package init

import (
	"fmt"
	"path/filepath"
	"run/global"
	"run/models"
	"run/models/constant"
	"run/util"
	"runtime"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func initEvents() {
	if global.Db == nil {
		global.Log.Fatal("not connect database")
		return
	}

	global.Log.Info("init events")
	err := global.Db.Transaction(func(tx *gorm.DB) error {
		for _, meta := range constant.EventList {
			event := models.Event{
				Base:     models.Base{Id: uint(meta.ID)},
				Name:     meta.Name,
				Distance: meta.Distance,
			}
			// 如果记录不存在（或已软删除），则创建；否则跳过
			if err := tx.Unscoped().Where("id = ?", meta.ID).Assign(event).FirstOrCreate(&models.Event{}).Error; err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		global.Log.Fatal("failed to init events", zap.Error(err))
		panic(err.Error())
	}
	global.Log.Info("init events successfully")
}

func initStandards() {
	_, filename, _, _ := runtime.Caller(0)
	projectRoot := filepath.Dir(filepath.Dir(filepath.Dir(filename))) // 从 orm/init/data.go 回到根目录
	csvPath := filepath.Join(projectRoot, "env", "standard.csv")
	standards, err := util.LoadStandardsFromCSV(csvPath)
	if err != nil {
		global.Log.Fatal("failed to load standards csv", zap.Error(err))
	}
	if len(standards) == 0 {
		global.Log.Info("standards csv is empty")
	}

	if global.Db == nil {
		global.Log.Fatal("not connect database")
		return
	}

	err = global.Db.Transaction(func(tx *gorm.DB) error {
		for _, std := range standards {
			// 幂等插入：四元组 (event_id, gender, level, standard_system) 唯一
			if err := tx.Unscoped().
				Where("event_id = ? AND gender = ? AND level = ? AND standard_system = ?",
					std.EventID, std.Gender, std.Level, std.StandardSystem).
				Assign(std).
				FirstOrCreate(&models.Standard{}).Error; err != nil {
				return fmt.Errorf("failed to seed standard %+v: %w", std, err)
			}
		}
		return nil
	})
	if err != nil {
		global.Log.Fatal("failed to init standards", zap.Error(err))
		panic(err.Error())
	}
	global.Log.Info("init standards successfully")
}
