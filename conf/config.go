package conf

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"os"
)

func init() {
	var err error
	Conf, err = loadConfig()
	if err != nil {
		log.Error("Abnormal loading of yml。")
	}

}
func loadConfig() (Config, error) {
	// 读取配置文件
	configData, err := os.ReadFile("conf/config.yml")
	if err != nil {
		return Config{}, err
	}

	// 解析配置文件
	var config Config
	err = yaml.Unmarshal(configData, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
