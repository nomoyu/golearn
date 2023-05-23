package conf

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path/filepath"
	"time"
)

/*配置日志*/
func InitLog() *zap.SugaredLogger {
	logMod := zapcore.DebugLevel
	if !viper.GetBool("mod.develop") {
		logMod = zapcore.InfoLevel
	}
	core := zapcore.NewCore(getEncoder(), zapcore.NewMultiWriteSyncer(getWriterSync(), zapcore.AddSync(os.Stdout)), logMod)
	return zap.New(core).Sugar()
}

/*格式定制*/
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(t.Local().Format(time.DateTime))
	}
	return zapcore.NewJSONEncoder(encoderConfig)
}

/*输出位置*/
func getWriterSync() zapcore.WriteSyncer {
	separator := string(filepath.Separator)
	rootDir, _ := os.Getwd()
	logFilePath := rootDir + separator + "log" + separator + time.Now().Format(time.DateOnly) + ".txt"

	lumberjackSyncer := &lumberjack.Logger{
		Filename:   logFilePath,
		MaxSize:    viper.GetInt("log.MaxSize"), // megabytes
		MaxBackups: viper.GetInt("log.MacBackups"),
		MaxAge:     viper.GetInt("log.MaxAge"), //days
		Compress:   true,                       // disabled by default
	}

	return zapcore.AddSync(lumberjackSyncer)

}
