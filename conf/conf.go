package conf

import (
	"github.com/spf13/viper"
)

/*配置中心*/
func InitConfig() {
	viper.SetConfigName("settings")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./conf/")
	err := viper.ReadInConfig()
	if err != nil {
		panic("Load config Error" + err.Error())
	}

}
