/**
 * @Author: xueyanghan
 * @File: comm_test.go
 * @Version: 1.0.0
 * @Description: desc.
 * @Date: 2023/7/26 11:15
 */

package std

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
	"testing"
)

func TestNilAssert(t *testing.T) {
	var x interface{} = nil

	v, ok := x.(int)
	if !ok {
		log.Println("x is not int type")
		return
	}
	log.Println(v)
}

func TestYaml(t *testing.T) {

	type TLSConfig struct {
		Enable  bool   `yaml:"enable"`
		CrtPath string `yaml:"crt_path"`
		KeyPath string `yaml:"key_path"`
		CaPath  string `yaml:"ca_path"`
		Mutual  bool   `yaml:"mutual"`
	}

	type LogConfig struct {
		Level    string `yaml:"level"`
		FilePath string `yaml:"file_path"`
	}

	type CryptoConfig struct {
		Soft    bool   `yaml:"soft"`
		LibPath string `yaml:"lib_path"`
	}

	type DBConfig struct {
		Type string `yaml:"type"`
		URL  string `yaml:"url"`
	}

	type ChainConfig struct {
		ContractName string `yaml:"contract_name"`
		ChainID      string `yaml:"chain_id"`
		ChainAddr    string `yaml:"chain_addr"`
		OrgID        string `yaml:"org_id"`
		NodeID       string `yaml:"node_id"`
		SignCrtPath  string `yaml:"sign_crt_path"`
		SignKeyPath  string `yaml:"sign_key_path"`
		TLSCrtPath   string `yaml:"tls_crt_path"`
		TLSKeyPath   string `yaml:"tls_key_path"`
		CACrtPath    string `yaml:"ca_crt_path"`
	}

	type TaskConfig struct {
		MaxSpeed int `yaml:"max_speed"`
		Timeout  int `yaml:"timeout"`
	}

	type Config struct {
		Common struct {
			ProxyNo string    `yaml:"proxyNo"`
			Port    int       `yaml:"port"`
			TLS     TLSConfig `yaml:"tls"`
		} `yaml:"common"`
		Log         LogConfig    `yaml:"log"`
		Crypto      CryptoConfig `yaml:"crypto"`
		DB          DBConfig     `yaml:"db"`
		ChainConfig ChainConfig  `yaml:"chain_config"`
		Task        TaskConfig   `yaml:"task"`
	}

	// Replace this with loading your actual YAML data
	yamlData := `
common:
  proxyNo: P202211011719562616055
  port: 18090
  tls:
    enable: false
    crt_path: ./certs/tls.crt
    key_path: ./certs/tls.key
    ca_path: ./certs/ca.crt
    mutual: false
log:
  level: debug
  file_path: ./proxy.log
crypto:
  soft: true
  lib_path: ""
db:
  type: "sqlite3"
  url: ./proxy.db
chain_config:
  contract_name: dev101
  chain_id: test
  chain_addr: 127.0.0.1:12301
  org_id: org1
  node_id: node1
  sign_crt_path: ./certs/sign.crt
  sign_key_path: ./certs/sign.key
  tls_crt_path: ./certs/tls.crt
  tls_key_path: ./certs/tls.key
  ca_crt_path: ./certs/ca.crt
task:
  max_speed: 100
  timeout: 600
`
	config := Config{}
	// Replace this with your actual YAML unmarshalling code
	err := yaml.Unmarshal([]byte(yamlData), &config)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the parsed configuration for testing
	fmt.Printf("%+v\n", config)
}
