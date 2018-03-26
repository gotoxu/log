package rotate

import (
	"os"
	"path/filepath"
)

// Config is Configuration for logging
type Config struct {
	Directory string
	Filename  string

	// MaxSize the max size in MB of the logfile before it's rolled
	MaxSize int

	// MaxBackups the max number of rolled files to keep
	MaxBackups int

	// MaxAge the max age in days to keep a logfile
	MaxAge int

	// Compress determines if the rotated log files should be compressed
	Compress bool
}

// Option 是rotate logger的配置函数
type Option func(*Config)

func defaultConfig() *Config {
	wd, _ := os.Getwd()
	dir := filepath.Join(wd, "logs")

	return &Config{
		Directory:  dir,
		Filename:   "servant.log",
		MaxSize:    500,
		MaxBackups: 30,
		MaxAge:     30,
		Compress:   true,
	}
}
