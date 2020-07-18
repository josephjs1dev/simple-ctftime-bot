package config

// Config is configuration data used by application
type Config struct {
	Host          string `mapstructure:"HOST"`
	Port          int    `mapstructure:"PORT"`
	ChannelSecret string `mapstructure:"CHANNEL_SECRET"`
	ChannelToken  string `mapstructure:"CHANNEL_TOKEN"`
}

// Reader is interface for global application config
type Reader interface {
	Read() error
	Decode() (*Config, error)
}

// ReadConfig read configuration file and return Config value
func ReadConfig(reader Reader) (*Config, error) {
	err := reader.Read()
	if err != nil {
		return nil, err
	}

	config, err := reader.Decode()
	if err != nil {
		return nil, err
	}

	return config, nil
}
