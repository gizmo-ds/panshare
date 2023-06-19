package global

import "github.com/jinzhu/configor"

type (
	Logger struct {
		Level         int8   `toml:"level" yaml:"level" env:"LOGGER_LEVEL"`
		ConsoleWriter bool   `toml:"console_writer" yaml:"console_writer" env:"LOGGER_CONSOLE_WRITER"`
		RequestLogger bool   `toml:"request_logger" yaml:"request_logger" env:"LOGGER_REQUEST_LOGGER"`
		Filename      string `toml:"filename" yaml:"filename" env:"LOGGER_FILENAME"`
		MaxAge        int    `toml:"max_age" yaml:"max_age" env:"LOGGER_MAX_AGE"`
		MaxBackups    int    `toml:"max_backups" yaml:"max_backups" env:"LOGGER_MAX_BACKUPS"`
	}
	Server struct {
		Url         string `toml:"url" yaml:"url" env:"SERVER_URL"`
		BindAddress string `toml:"bind_address" yaml:"bind_address" env:"SERVER_BIND_ADDRESS"`
		TlsCertFile string `toml:"tls_cert_file" yaml:"tls_cert_file" env:"SERVER_TLS_CERT_FILE"`
		TlsKeyFile  string `toml:"tls_key_file" yaml:"tls_key_file" env:"SERVER_TLS_KEY_FILE"`
	}

	Config struct {
		Logger Logger `toml:"logger" yaml:"logger"`
		Server Server `toml:"server" yaml:"server"`
	}
)

var conf *Config

func InitConfig(filename string) error {
	return configor.New(&configor.Config{Environment: "production"}).Load(&conf, filename)
}

func C() *Config {
	return conf
}
