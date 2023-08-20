/**
 * @Author: xueyanghan
 * @File: crypto_config.go
 * @Version: 1.0.0
 * @Description: desc.
 * @Date: 2023/8/20 14:19
 */

package types

// CryptoConfig - crypto config
type CryptoConfig struct {
	Soft    bool   `yaml:"soft"`
	LibPath string `yaml:"lib_path"`
}
