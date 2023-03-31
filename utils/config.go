package utils

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Name string
}

//初始化配置
func InitConfig(fileName string) error {

	c := Config{
		Name: fileName,
	}
	// 初始化配置文件
	if err := c.setConfig(); err != nil {
		return err
	}

	// 检测必须的配置项
	if err := c.checkConfig(); err != nil {
		return err
	}

	c.watchConfig()
	return nil
}

// 监听配置文件是否改变,用于热更新
func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("配置文件修改更新: %s\n", e.Name)
	})
}

func (c *Config) setConfig() error {
	viper.AddConfigPath("./conf/")
	if c.Name != "" {
		// 如果指定了配置文件，则解析指定的配置文件
		viper.SetConfigName(c.Name)
	} else {
		// 如果没有指定配置文件，则解析默认的配置文件
		viper.SetConfigName("config")
	}
	// 设置配置文件格式为YAML
	viper.SetConfigType("json")
	// viper解析配置文件
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}

// 检测必须的配置项
func (c *Config) checkConfig() error {
	//if !viper.IsSet("enroll_rule.online_first_course_id") {
	//	return errors.New("配置缺失错误：请配置enroll_rule.online_first_course_id参数")
	//}
	//if !viper.IsSet("enroll_rule.online_second_course_id") {
	//	return errors.New("配置缺失错误：请配置enroll_rule.online_second_course_id 参数")
	//}

	return nil
}
