package services

import (
	"github.com/Unknwon/goconfig"
)

type ConfigService struct {
	configFile *goconfig.ConfigFile
}

//初始化加载config文件
func (this *ConfigService) initConfigService() error{
	if this.configFile != nil {
		return nil
	}
	var err error
	this.configFile, err = goconfig.LoadConfigFile("conf/config.ini")
	if err != nil {
		Debug("load config file fail:", err.Error())
	}
	return err
}

//检查config加载的正确性
func (this *ConfigService) checkConfigValid() bool {
	if this.configFile == nil {
		return false
	}
	return true
}

//获取config
func (this *ConfigService) getConfigInstance() {
	if this.configFile == nil {
		this.initConfigService()
	}
}

//获取区域配置
func GetConfigSection(section string) map[string]string {
	service := &ConfigService{}
	service.getConfigInstance()
	if service.checkConfigValid() {
		secontion, err := service.configFile.GetSection(section)
		if err != nil {
			Debug("get section failed:", section, err.Error())
			return nil
		}
		return secontion
	}
	return nil
}

//获取精确配置
func GetConfigValue(section, key string) string{
	service := &ConfigService{}
	service.getConfigInstance()
	if service.checkConfigValid() {
		value, err := service.configFile.GetValue(section, key)
		if err != nil {
			Error("get config value fail:", section, key, err.Error())
			return ""
		}
		return value
	}
	return ""
}