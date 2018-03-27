package logstash

// DataType specify either list or channel
type DataType int8

const (
	// List use rpush
	List DataType = iota
	// Channel use publish
	Channel
)

// Option 是redis syncer的配置函数
type Option func(*Config)

// Config is Configuration for redis input
type Config struct {
	DataType DataType
	DB       int
	Host     string
	Password string
	Port     int
	Key      string
}

func newConfig(key, host string) *Config {
	return &Config{
		Key:      key,
		Host:     host,
		DataType: List,
		DB:       0,
		Port:     6379,
		Password: "",
	}
}
