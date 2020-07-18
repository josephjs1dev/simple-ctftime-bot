package config

import (
	"reflect"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/mitchellh/mapstructure"
)

// EnvReader is implementation of ConfigReader which read env
type EnvReader struct {
	EnvFiles []string
	env      map[string]string
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
	var envData map[string]string
	var err error

	envData, err = godotenv.Read(r.EnvFiles...)

	r.env = envData

	return err
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
