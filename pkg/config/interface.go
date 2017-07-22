package config

// Create mock using:
// mockgen -source=pkg/config/interface.go -destination=pkg/config/mock/config_mock.go
type Interface interface {
	init() error
	ReadConfig(configFilePath string) error
	Set(key string, value interface{})
	SetDefault(key string, value interface{})
	AllSettings() map[string]interface{}
	IsSet(key string) bool
	Get(key string) interface{}
	GetBool(key string) bool
	GetInt(key string) int
	GetString(key string) string
	GetBase64Decoded(key string) (string, error)
}