package config

import (
	"reflect"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/mitchellh/mapstructure"
)

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

// Service is our config object
type Service struct {
	reader Reader
	config *Config
}

// GetConfig will return service config
func (s Service) GetConfig() *Config {
	return s.config
}

// EnvReader is implementation of ConfigReader which read env
type EnvReader struct {
	env map[string]string
}

func configHook(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
	if f.Kind() == reflect.String && t.Kind() == reflect.Int {
		result, err := strconv.Atoi(data.(string))

		return result, err
	}

	return data, nil
}

// Read utilizes godotenv Read to read the environment
func (r *EnvReader) Read() error {
	env, err := godotenv.Read()

	if err != nil {
		return err
	}

	r.env = env

	return nil
}

// Decode uses mapstructure to change env to ConfigData and return Config
func (r *EnvReader) Decode() (*Config, error) {
	config := Config{}
	dc := &mapstructure.DecoderConfig{Result: &config, DecodeHook: configHook}
	decoder, err := mapstructure.NewDecoder(dc)

	if err != nil {
		return nil, err
	}

	err = decoder.Decode(r.env)

	if err != nil {
		return nil, err
	}

	return &config, nil
}

// ReadConfig read configuration file and return Service
func ReadConfig(reader Reader) (*Service, error) {
	service := &Service{
		reader: reader,
		config: nil,
	}
	err := service.reader.Read()

	if err != nil {
		return service, err
	}

	config, err := service.reader.Decode()

	if err != nil {
		return service, err
	}

	service.config = config

	return service, nil
}
