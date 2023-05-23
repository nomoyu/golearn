package conf

import (
	"github.com/nomoyu/golearn/model"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

// InitDb /*配置数据库*/
func InitDb() (*gorm.DB, error) {
	LogMod := logger.Info
	if !viper.GetBool("mod.develop") {
		LogMod = logger.Error
	}
	db, err := gorm.Open(mysql.Open(viper.GetString("db.dsn")), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "sys_",
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(LogMod),
	})
	if err != nil {
		return nil, err
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetConnMaxLifetime(time.Hour)
	// 初始化数据库表
	db.AutoMigrate(&model.User{})
	return db, nil

}
