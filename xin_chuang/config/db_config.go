/**
 * @Author: xueyanghan
 * @File: db_config.go
 * @Version: 1.0.0
 * @Description: desc.
 * @Date: 2024/4/26 18:35
 */

package config

// DBConfig - db config
type DBConfig struct {
	Type string `yaml:"type"`
	URL  string `yaml:"url"`
}

type MinioConfig struct {
	Endpoint        string `yaml:"endpoint"`
	AccessKeyID     string `yaml:"accessKeyID"`
	SecretAccessKey string `yaml:"secretAccessKey"`
	UseSSL          bool   `yaml:"useSSL"`
	BucketName      string `yaml:"bucketName"`
}
