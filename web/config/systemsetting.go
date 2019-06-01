package config

import (
	"github.com/go-logging/utilities"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const(
	SERVER_ID="server.serverID"
	SERVER_PORT="server.port"
	LOG_FILE_NAME="log.filename"
	LOG_LEVEL="log.level"
	LOG_KEYS="log.keys"
)

func InitializeConfig() {
	viper.SetDefault(LOG_LEVEL, "debug")
	viper.SetDefault(SERVER_ID, "M01")
	viper.SetDefault(LOG_FILE_NAME,"/Users/300067308/app.log")
	viper.SetDefault(LOG_KEYS,"ReqID Username")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("toml")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatalf("Error reading config file, %s", err)
	}
}

func GetServerID() string{
	id,ok:=utilities.HasValue(viper.GetString(SERVER_ID))
	if ok{
		return id
	}
	return ""
}

func GetPort() string{
	v,ok:=utilities.HasValue(viper.GetString(SERVER_PORT))
	if ok{
		return v
	}
	return ""
}

func GetLogFileName() string{
	v,ok:=utilities.HasValue(viper.GetString(LOG_FILE_NAME))
	if ok{
		return v
	}
	return ""
}

func GetLogLevel() string{
	v,ok:=utilities.HasValue(viper.GetString(LOG_LEVEL))
	if ok{
		return v
	}
	return ""
}

func GetLogKeys() []string{
	return viper.GetStringSlice(LOG_KEYS)
}