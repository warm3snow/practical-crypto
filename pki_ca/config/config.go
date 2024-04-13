/**
 * @Author: xueyanghan
 * @File: config.go
 * @Version: 1.0.0
 * @Description: desc.
 * @Date: 2023/8/16 18:27
 */

package config

import (
	"io/ioutil"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

const (
	defaultConfigPath = "./github.com/warm3snow/crypto-service-backend/.yaml"
	defaultPageSize   = 10
)

// InitConfig - github.com/warm3snow/crypto-service-backend/ config init, load from yaml file
func InitConfig(path string) error {
	if path == "" {
		path = defaultConfigPath
	}

	yamlData, err := ioutil.ReadFile(path)
	if err != nil {
		return errors.Wrap(err, "read config file failed")
	}

	cfg := &Config{}

	err = yaml.Unmarshal(yamlData, cfg)
	if err != nil {
		return errors.Wrap(err, "unmarshal config file failed")
	}

	Conf = cfg
	return nil
}
