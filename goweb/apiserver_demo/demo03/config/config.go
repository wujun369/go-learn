package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
	"strings"
)

type Config struct {
	Name string
}

func Init(cfg string) error {

	c := Config{
		Name: cfg,
	}

	if err := c.initConfig(); err != nil{//初始化配置文件
		return err
	}

	c.initLog()	//初始化日志配置

	c.watchConfig()	//监控配置文件变化，并热加载程序


	return nil
}

func (c *Config) initConfig() error{ //通过 viper 加载配置文件
	if c.Name != "" {
		viper.SetConfigFile(c.Name)
	}else{
		viper.AddConfigPath("apiserver_demo/demo03/conf")
		viper.SetConfigName("config")
	}

	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("APISERVER")

	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	if err := viper.ReadInConfig(); err != nil{
		return err
	}
	return nil
}

func (c *Config) initLog() { //从配置文件加载日志配置到 lexkong/log 中
	passLagerCfg := log.PassLagerCfg{
		Writers:        viper.GetString("log.writers"),
		LoggerLevel:    viper.GetString("logger_level"),
		LoggerFile:     viper.GetString("logger_file"),
		LogFormatText:  viper.GetBool("log_format_text"),
		RollingPolicy:  viper.GetString("rollingPolicy"),
		LogRotateDate:  viper.GetInt("log_rotate_date"),
		LogRotateSize:  viper.GetInt("log_rorate_size"),
		LogBackupCount: viper.GetInt("log_backup_count"),
	}

	log.InitWithConfig(&passLagerCfg)

}

func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		log.Infof("Config file changed: %s",in.Name)
	})
}


