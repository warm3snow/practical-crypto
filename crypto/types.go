/**
 * @Author: xueyanghan
 * @File: types.go
 * @Version: 1.0.0
 * @Description: desc.
 * @Date: 2023/8/16 18:27
 */

package crypto

// Conf -
var Conf *Config

// Config - config
type Config struct {
	Common CommonConfig `yaml:"common"`
	Log    LogConfig    `yaml:"log"`
	Crypto CryptoConfig `yaml:"crypto"`
	DB     DBConfig     `yaml:"db"`
	Task   TaskConfig   `yaml:"task"`

	SignCrtBytes []byte
	SignKeyBytes []byte
	TlsCrtBytes  []byte
	TlsKeyBytes  []byte
	CaCrtBytes   []byte
}

// CommonConfig - common config
type CommonConfig struct {
	ServerNo string    `yaml:"server_no"`
	Port     int       `yaml:"port"`
	TLS      TLSConfig `yaml:"tls"`
}

// TLSConfig - tls config for grpc server
type TLSConfig struct {
	Enable  bool   `yaml:"enable"`
	CrtPath string `yaml:"crt_path"`
	KeyPath string `yaml:"key_path"`
	CaPath  string `yaml:"ca_path"`
	Mutual  bool   `yaml:"mutual"`
}

// LogConfig - log config
type LogConfig struct {
	Level    string `yaml:"level"`
	FilePath string `yaml:"file_path"`
}

// CryptoConfig - crypto config
type CryptoConfig struct {
	Soft    bool   `yaml:"soft"`
	LibPath string `yaml:"lib_path"`
}

// DBConfig - db config
type DBConfig struct {
	Type string `yaml:"type"`
	URL  string `yaml:"url"`
}

// TaskConfig - task config
type TaskConfig struct {
	MaxSpeed int `yaml:"max_speed"`
	Timeout  int `yaml:"timeout"`
}
