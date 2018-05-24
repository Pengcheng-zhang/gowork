package common

import (
	"github.com/Unknwon/goconfig"
)

var g_config *goconfig.ConfigFile

//初始化加载config文件
func init() {
	config, err := goconfig.LoadConfigFile("conf/config.ini")
	if err != nil {
		Debug("load config file fail:", err.Error())
	}
	g_config = config
}

//检查config加载的正确性
func checkConfigValid() bool {
	if g_config == nil {
		return false
	}
	return true
}

//获取区域配置
func GetConfigSection(section string) map[string]string {
	if checkConfigValid() {
		secontion, err := g_config.GetSection(section)
		if err != nil {
			Debug("get section failed:", section, err.Error())
			return nil
		}
		return secontion
	}
	return nil
}

//获取精确配置
func GetConfigValue (section, key string) string{
	if checkConfigValid() {
		value, err := g_config.GetValue(section, key)
		if err != nil {
			Error("get config value fail:", section, key, err.Error())
			return ""
		}
		return value
	}
	return ""
}