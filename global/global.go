package global

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Log *zap.SugaredLogger
	DB  *gorm.DB
)
